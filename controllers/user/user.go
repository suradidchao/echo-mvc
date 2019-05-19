package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/suradidchao/echo-mvc/models/user"
)

// Controller ...
type Controller struct {
	UserModel user.Model
}

// NewController ...
func NewController(userModel user.Model) Controller {
	return Controller{
		UserModel: userModel,
	}
}

//RegisterRoutes ...
func (uc *Controller) RegisterRoutes(route string, e *echo.Echo) {
	user := e.Group(route)
	user.GET("", uc.GetUsers)
	user.GET("/:id", uc.GetUser)
	user.POST("", uc.CreateUser)
	user.PUT("/:id", uc.UpdateUser)
	user.DELETE("/:id", uc.DeleteUser)
}

// GetUsers ...
func (uc *Controller) GetUsers(c echo.Context) error {
	users, err := uc.UserModel.GetUsers()
	if err != nil {
		log.Fatal("Cannot get users")
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser ...
func (uc *Controller) GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := uc.UserModel.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}

// CreateUser ...
func (uc *Controller) CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Age field is malformed")
	}
	nationality := c.FormValue("nationality")
	userData := map[string]interface{}{
		"name":        name,
		"age":         age,
		"nationality": nationality,
	}
	createdUser, err := uc.UserModel.CreateUser(userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Cannot create user")
	}
	return c.JSON(http.StatusCreated, createdUser)
}

// UpdateUser ...
func (uc *Controller) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Age field is malformed")
	}
	nationality := c.FormValue("nationality")
	userData := map[string]interface{}{
		"name":        name,
		"age":         age,
		"nationality": nationality,
	}
	updatedUser, err := uc.UserModel.UpdateUser(id, userData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Cannot update user")
	}
	return c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser ...
func (uc *Controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	deletedUser, err := uc.UserModel.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "User to be deleted not found in the database")
	}
	return c.JSON(http.StatusOK, deletedUser)
}
