package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.Nick == "" {
		return errors.New("Nick is required")
	}

	if user.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid email")
	}

	if (step == "create") && user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

func trimString(text string) string {
	return strings.TrimSpace(text)
}

func (user *User) format() {
	user.Name = trimString(user.Name)
	user.Nick = trimString(user.Nick)
	user.Email = trimString(user.Email)
}
