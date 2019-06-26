package user

import (
	"errors"

	"github.com/asdine/storm"
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
var (
	ErrRecordInvalid = errors.New("Invalid record")
)

// All retrieves all users from database
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

// One retrieves a single user record from database
func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	user := new(User)
	err = db.One("ID", id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete removes a given record from databse
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	user := new(User)
	err = db.One("ID", id, user)
	if err != nil {
		return err
	}
	return db.DeleteStruct(user)
}

// Save updates or creates a given record in the database
func (user *User) Save() error {
	if err := user.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(user)
}

// validate makes sure that the record contains valid data
func (user *User) validate() error {
	if user.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
