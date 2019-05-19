package adapters

import (
	"github.com/suradidchao/echo-mvc/entities"
)

// Adapter ...
type Adapter interface {
	FindOne(id string) (entities.User, error)
	FindMany() ([]entities.User, error)
	CreateOne(payload map[string]interface{}) (entities.User, error)
	UpdateOne(id string, payload map[string]interface{}) (entities.User, error)
	DeleteOne(id string) (entities.User, error)
}
