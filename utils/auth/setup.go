package auth

import (
	"fmt"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"log"
	util "whm-api/utils"
	"whm-api/utils/db/users"

	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Setup() {
	apiBasePath := "/v1/auth"
	websiteBasePath := "/auth"
	cookiesSecure := util.GodotEnvBool("SECURE_COOKIES", true)

	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: fmt.Sprintf("http://%s:3567", util.GodotEnv("AUTH_HOST")),
			APIKey:        util.GodotEnv("AUTH_KEY"),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "WHM",
			APIDomain:       "https://api.antonwy.me",
			WebsiteDomain:   "https://antonwy.me",
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(&epmodels.TypeInput{
				SignUpFeature: &epmodels.TypeInputSignUp{
					FormFields: []epmodels.TypeInputFormField{
						{
							ID: "name",
						},
					},
				},
				Override: &epmodels.OverrideStruct{
					APIs: func(originalImplementation epmodels.APIInterface) epmodels.APIInterface {
						originalSignUpPOST := *originalImplementation.SignUpPOST

						*originalImplementation.SignUpPOST = func(formFields []epmodels.TypeFormField, options epmodels.APIOptions) (epmodels.SignUpResponse, error) {
							return SignUp(formFields, options, originalSignUpPOST)
						}

						return originalImplementation
					},
				},
			}),
			session.Init(&sessmodels.TypeInput{
				CookieSecure: &cookiesSecure,
			}),
		},
	})

	if err != nil {
		panic(err.Error())
	}

	user := users.User{
		Email: util.GodotEnv("DEFAULT_EMAIL"),
		Name:  util.GodotEnv("DEFAULT_USERNAME"),
		Role:  "admin",
	}

	response, err := emailpassword.SignUp(user.Email, util.GodotEnv("DEFAULT_PASSWORD"))
	if err != nil {
		log.Println("Couldn't signup default user!")
		panic(err.Error())
	}

	if response.OK != nil {
		user.ID = response.OK.User.ID
		if err := user.Create(); err != nil {
			log.Println("Couldn't create default user for db!")
			panic(err.Error())
		}

		log.Println(response)
	}

}
