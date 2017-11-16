package data

import (
	"encoding/json"
	"os"
)

//F fan
type F struct {
	A string
	B string
}

//Person fan
type Person struct {
	Name   string
	Mobile string
	Lis    F
}

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

//Insert fan
func Insrert() {
	var db = GetDB("ApplicationCentre")
	defer db.Session.Close()

	p := Person{}
	p.Lis = F{}
	p.Mobile = "321"
	p.Name = "a"
	err := db.C("fan").Insert(p)

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	group.ID = 3
	if err != nil {

	}
	b, err1 := json.Marshal(group)
	if err1 != nil {

	}
	os.Stdout.Write(b)
}
