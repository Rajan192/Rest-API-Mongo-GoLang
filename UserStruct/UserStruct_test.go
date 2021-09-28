package Userstruct

import (
	"log"
	"testing"
)

func TestCheckVAlidation(t *testing.T) {
	u := User{FirstName: "Rajan", LastName: "Prajapati",Mobile: "123456789"}
	err := u.Validate()
	if err != nil {
		log.Fatal(err)
	}
}
