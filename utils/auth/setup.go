package auth

import (
	util "whm-api/utils"

	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func Setup() {
	apiBasePath := "/v1/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://whm-auth:3567",
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

						(*originalImplementation.SignUpPOST) = func(formFields []epmodels.TypeFormField, options epmodels.APIOptions) (epmodels.SignUpResponse, error) {
							return SignUp(formFields, options, originalSignUpPOST)
						}

						return originalImplementation
					},
				},
			}),
			session.Init(nil),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}
