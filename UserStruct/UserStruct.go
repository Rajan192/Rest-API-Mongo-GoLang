package Userstruct

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Active    bool               `json:"active,omitempty" bson:"active,omitempty"`
	Age       *Age               `json:"age,omitempty" bson:"age,omitempty"`
}

type Age struct {
	Value    int    `json:"age,omitempty" bson:"age,omitempty"`
	Interval string `json:"interval,omitempty" bson:"interval,omitempty"`
}
