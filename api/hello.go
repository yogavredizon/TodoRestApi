package api

import (
	"fmt"

	m "github.com/yogavredizon/go-modules/models"
)

func SayHello() {
	person := m.Person{Name: "Yoga"}
	fmt.Println(person.GetName())
}
