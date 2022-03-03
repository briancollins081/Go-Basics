package user

import (
	"errors"

	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
)

// User holds data for a single user
type User struct {
	ID   bson.ObjectId `json:"id" storm:"id"`
	Name string        `json:"name"`
	Role string        `json:"role"`
}

const (
	dbPath = "users.db"
)

// errors
var ErrorRecordInvalid error = errors.New("record is invalid")

// All retrieves all users from the database
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	users := []User{}
	err = db.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// One retrieves a single user from the database that matches the id passed
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Delete removes a given user record that matches the id passed from the db
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	u := new(User)
	err = db.One("ID", id, u)
	if err != nil {
		return err
	}

	return db.DeleteStruct(u)
}

// Save saves an update or adds a new user record
func (u *User) Save() error {
	if err := u.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// validate makes sure that therecord contains valid data
func (u *User) validate() error {
	if u.Name == "" {
		return ErrorRecordInvalid
	}
	return nil
}
