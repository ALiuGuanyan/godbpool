package postgre

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Postgre SQL
type Postgre struct {
	args interface{}
}

// New will init a struct
func New(args interface{}) Postgre {
	return Postgre{args}
}

// Open will build a connection with database
func (p Postgre) Open() (*gorm.DB, error) {
	return gorm.Open("postgres", p.args)
}

// Args will return the connection args
func (p Postgre) Args() interface{} {
	return p.args
}
