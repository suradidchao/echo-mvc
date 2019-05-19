package user

import (
	"github.com/suradidchao/echo-mvc/entities"
	"github.com/suradidchao/echo-mvc/models/user/adapters"
)

// Model is ...
type Model struct {
	Adapter adapters.Adapter
}

// NewModel ...
func NewModel(adapter adapters.Adapter) Model {
	return Model{
		Adapter: adapter,
	}
}

// GetUsers is ...
func (m *Model) GetUsers() ([]entities.User, error) {
	return m.Adapter.FindMany()
}

// GetUser ...
func (m *Model) GetUser(userID string) (entities.User, error) {
	return m.Adapter.FindOne(userID)
}

// CreateUser ...
func (m *Model) CreateUser(userData map[string]interface{}) (entities.User, error) {
	return m.Adapter.CreateOne(userData)
}

// UpdateUser ...
func (m *Model) UpdateUser(userID string, userData map[string]interface{}) (entities.User, error) {
	return m.Adapter.UpdateOne(userID, userData)
}

// DeleteUser ...
func (m *Model) DeleteUser(userID string) (entities.User, error) {
	return m.Adapter.DeleteOne(userID)
}
