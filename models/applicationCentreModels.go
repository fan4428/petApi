package models

//User fan
type User struct {
	ID        string `bson:"_id"`
	Name      string `bson:"name"`
	Password  string `bson:"password"`
	TermFlag  bool   `bson:"termflag"`
	AdminRole string `bson:"adminrole"`
}
