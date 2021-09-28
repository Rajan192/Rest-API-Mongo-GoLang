package Userstruct

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/go-playground/validator/v10"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty" validate:"required"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty" validate:"required" `
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty" validate:"required"`
	Active    bool               `json:"active,omitempty" bson:"active,omitempty"`
	Age       *Age               `json:"age,omitempty" bson:"age,omitempty"`
}

type Age struct {
	Value    int    `json:"age,omitempty" bson:"age,omitempty" `
	Interval string `json:"interval,omitempty" bson:"interval,omitempty"`
}

func (u *User) Validate() error {
	validate := validator.New()
	//validate.RegisterValidation()
	return validate.Struct(u)
}
