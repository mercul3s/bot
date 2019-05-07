package users

import (
	"errors"

	"github.com/boltdb/bolt"
)

// UserCfg holds types needed to perform user operations
type UserCfg struct {
	store *bolt.DB
}

// New returns an initialied UserCfg.
func New(s *bolt.DB) *UserCfg {
	return &UserCfg{
		store: s,
	}
}

// Add parses slack command text into a user struct, marshals it to JSON,
// and then stores the resulting json in the store.
func (u UserCfg) Add(u string) error {
	// parse string to user struct
	// marshal to json
	// store in db
	return errors.New("error adding new user: ", u)
}

// List all users in the store.
func List() error {

	return errors.New("error listing users: ")
}

// Remove a user from the store.
func (u UserCfg) Remove(u string) error {

}

// AddCommand takes a new command and adds it to the commands store.
func AddCommand(c string) error {
	return errors.New("error adding new command: ", c)
}

// Help returns the list of commands available
func Help() error {
	return errors.New("error getting commands")
}
