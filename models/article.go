package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Article struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	AuthorID bson.ObjectId `bson:"authorId" json:"authorId"`
	Title string `bson:"title" json:"title"`
	Content string `bson:"content" json:"content"`
}

type Articles []Article