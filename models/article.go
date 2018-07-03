package models

import (
	"github.com/ismayilmalik/golang-blog/config"
	"github.com/ismayilmalik/golang-blog/infrastructure/database"
	"gopkg.in/mgo.v2/bson"
)

type (
	Article struct {
		ID       bson.ObjectId `bson:"_id" json:"id"`
		AuthorID bson.ObjectId `bson:"authorId" json:"authorId"`
		Title    string        `bson:"title" json:"title"`
		Content  string        `bson:"content" json:"content"`
	}

	Articles []Article
)

func (a *Article) GetAll() (Articles, error) {
	var articles Articles
	db := database.Instance()
	defer db.Close()

	err := db.DB(config.DATABASE).C(config.ARTICLES).Find(nil).All(&articles)

	return articles, err
}
