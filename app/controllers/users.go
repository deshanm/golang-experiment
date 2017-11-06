package controllers

import (
	"github.com/revel/revel"
)

type Stuff struct {
	Foo string ` json:"foo" xml:"foo" `
	Bar int    ` json:"bar" xml:"bar" `
}

type Users struct {
	*revel.Controller
}

func (c Users) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}

func (c Users) List(myName string) revel.Result {
	data := make(map[string]interface{})
	data["error"] = nil
	stuff := Stuff{Foo: "xyz", Bar: 999}
	data["stuff"] = stuff
	return c.RenderJSON(data)
}
