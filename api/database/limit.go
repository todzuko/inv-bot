package database

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserLimit struct {
	ID     primitive.ObjectID `bson:"_id"`
	ChatId int64              `bson:"chat_id"`
	Limit  int                `bson:"limit"`
}

func GetLimit(chat int64) (int, bool) {
	filterDoc := bson.M{"chat_id": chat}
	resultDoc := &UserLimit{}

	res := db.Collection("user_limit").FindOne(ctx, filterDoc)
	if err := res.Decode(resultDoc); err != nil {
		fmt.Println("Decode error:", err)
		return 0, false
	}
	if resultDoc.Limit > 0 {
		return resultDoc.Limit, true
	}
	return 0, false
}

func SetLimit(chat int64, limit int) bool {
	filterDoc := bson.M{"chat_id": chat}
	updateDoc := bson.M{"$set": bson.M{"limit": limit}}
	upd := db.Collection("user_limit").FindOneAndUpdate(ctx, filterDoc, updateDoc)
	if upd.Err() == nil {
		return true
	}
	fmt.Println(upd.Err().Error())

	_, err := db.Collection("user_limit").InsertOne(ctx, UserLimit{ChatId: chat, Limit: limit})
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
