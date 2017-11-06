package controllers

import (
	"github.com/revel/revel"
)

type Users struct {
	*revel.Controller
}

func (c Users) Index() revel.Result {
	greeting := "Users Page"
	return c.Render(greeting)
}

func (c Users) UserById() revel.Result {
	greeting := "UserById Page"
	return c.Render(greeting)
}

func (c Users) CreateUser() revel.Result {
	email := c.Params.Form.Get("email")
	name := c.Params.Form.Get("name")
	password := c.Params.Form.Get("password")

	greeting := "CreateUser Page"
	return c.Render(greeting)
}

func (c Users) DeleteUser(id int) revel.Result {
	greeting := "DeleteUser Page"
	return c.Render(greeting)
}

func (c Users) UpdateUser() revel.Result {
	greeting := "UpdateUser Page"
	return c.Render(greeting)
}
