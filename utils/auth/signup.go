package auth

import (
	"errors"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/supertokens"
	"whm-api/utils/db/users"
	userRoles "whm-api/utils/db/users/roles"

	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
)

func SignUp(formFields []epmodels.TypeFormField, options epmodels.APIOptions, orginalSignUp func(formFields []epmodels.TypeFormField, options epmodels.APIOptions) (epmodels.SignUpResponse, error)) (epmodels.SignUpResponse, error) {

	username := FromFormFields("name", formFields)

	if username == "" {
		return epmodels.SignUpResponse{}, errors.New("please add the name of the user")
	}

	response, err := orginalSignUp(formFields, options)
	if err != nil {
		return epmodels.SignUpResponse{}, err
	}

	if response.OK != nil {
		user := response.OK.User

		dbUser := users.User{
			ID:    user.ID,
			Name:  username,
			Email: user.Email,
			Role:  userRoles.Admin,
		}

		if err := dbUser.Create(); err != nil {
			userFromMail, err := emailpassword.GetUserByEmail(user.Email)
			if err != nil {
				return epmodels.SignUpResponse{}, errors.New("Couldn't insert user in the Database and then also couldn't find an user with mail: " + user.Email)
			}

			if err := supertokens.DeleteUser(userFromMail.ID); err != nil {
				return epmodels.SignUpResponse{}, errors.New("Couldn't insert user in the Database and then also couldn't delete created user with mail: " + user.Email)
			}

			return epmodels.SignUpResponse{}, errors.New("Couldn't insert user in the Database with mail: " + user.Email)
		}
	}
	return response, nil
}

func FromFormFields(id string, formFields []epmodels.TypeFormField) string {
	for _, field := range formFields {
		if field.ID == id {
			return field.Value
		}
	}

	return ""
}
