package validation

import (
	"errors"
	models "microServiceBoilerplate/services/hellgate/graph/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func CreateUser(data *models.CreateUserInput) error {
	return validation.ValidateStruct(data,
		validation.Field(&data.Name, validation.Length(1, 32)),
		validation.Field(&data.Gender, validation.By(func(value interface{}) error {
			if value != "male" && value != "female" {
				return errors.New("gender must be 'male' or 'female'")
			}
			return nil
		})),
		validation.Field(&data.Email, is.Email),
		validation.Field(&data.Password, validation.Length(8, 32)),
	)
}

func CreatePost(data *models.CreatePostInput) error {
	return validation.ValidateStruct(data,
		validation.Field(&data.Title, validation.Length(5, 50)),
		validation.Field(&data.Body, validation.Length(5, 400)),
	)
}
