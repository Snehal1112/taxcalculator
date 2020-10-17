package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type User struct {
	FirstName string `json:"first_name,omitempty" bson:"first_name, omitempty"`
	LastName string `json:"last_name,omitempty" bson:"last_name, omitempty"`
	UserID string `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

func NewUser() *User {
	return &User{}
}

// ToJson convert a User to a json string
func (u *User) ToJson() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// FromJson will decode the input and return a User
func (u *User)FromJson(data io.Reader) *User {
	decoder := json.NewDecoder(data)
	err := decoder.Decode(&u)
	if err == nil {
		return u
	}
	return nil
}

// ComparePassword compares the hash.
func (u *User)ComparePassword(hash string, password string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword encrypt the password.
func (u *User)HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

// ToBSON will decode the input and return a User
func (u *User)ToBSON(data bson.M) *User {
	uProps, _ := bson.Marshal(data)
	bson.Unmarshal(uProps, &u)
	return u
}

func (u *User) Sensitized() *User {
	u.Password = ""
	return u
}