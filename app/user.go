package app

import (
	"github.com/snehal1112/taxcal/model"
	"go.mongodb.org/mongo-driver/bson"
)
var collection = "user"

func (a *App) CreateUser(user *model.User) error {
	user.Password = user.HashPassword(user.Password)
	return getSession(a).CreateDocument(collection, user)
}

func (a *App) AuthUser(user *model.User) (bool, *model.User, error) {
	result, err := getSession(a).FindOne(collection, bson.D{{"user_id", user.UserID}})
	if err != nil{
		return false, nil, err
	}

	userFind := model.NewUser().ToBSON(result)
	return user.ComparePassword(userFind.Password, user.Password), userFind.Sensitized(), nil
}
