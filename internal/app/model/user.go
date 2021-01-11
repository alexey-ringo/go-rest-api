package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncriptedPassword string
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
