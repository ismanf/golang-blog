package models

import (
	"github.com/ismayilmalik/golang-blog/config"
	"github.com/ismayilmalik/golang-blog/infrastructure/database"
	"gopkg.in/mgo.v2/bson"
)

type (
	Credentials struct {
		Login    string `bson:"login" json:"login"`
		Password string `bson:"string" json:"string"`
	}

	Author struct {
		Credentials
		ID        bson.ObjectId `bson:"_id" json:"id"`
		FirstName string        `bson:"firstName" json:"firstName"`
		LastName  string        `bson:"lastName" json:"lastName"`
		Email     string        `bson:"email" json:"email"`
	}

	Authors []Author
)

func (a *Author) GetByID() error {
	db := database.Instance()
	defer db.Close()

	return db.DB(config.DATABASE).C(config.AUTHORS).Find(bson.M{"_id": a.ID}).One(a)
}

func (a *Author) GetByLogin() error {
	db := database.Instance()
	defer db.Close()

	return db.DB(config.DATABASE).C(config.AUTHORS).Find(bson.M{"login": a.Login}).One(a)
}

func (a *Author) Create() error {
	db := database.Instance()
	defer db.Close()

	a.ID = bson.NewObjectId()

	return db.DB(config.DATABASE).C(config.AUTHORS).Insert(a)
}
