package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncriptedPassword string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncriptedPassword == "")), validation.Length(6, 100)),
	)
}

//BeforeCreate ...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncriptedPassword = enc
	}

	return nil
}

func encryptString(s string) (string, error) {
	//Передаваемая на входе строка приводится к массиву байтов
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	//Обратное преобразование зашифрованного массива байтов в строку
	return string(b), nil
}
