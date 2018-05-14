package entities

import "time"

//Credentials structure for Authentication
type Credentials struct {
	Username string    `datastore:"description"`
	Password string    `datastore:"description"`
	Created  time.Time `datastore:"created"`
}
