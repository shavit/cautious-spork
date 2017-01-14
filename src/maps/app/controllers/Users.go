package controllers

import (
  "github.com/revel/revel"
  "encoding/json"
)

type Users struct {
  *revel.Controller
}

const (
  email = "user@example.com"
  password = "g00dpassword"
)

func (c *Users) Login() (revel.Result) {
  var err error
  var requestData map[string]string

  // Encode the request
  err = json.NewDecoder(c.Request.Body).Decode(&requestData)
  if err != nil {
    panic(err)
  }

  if (requestData["username"] == email) && (requestData["password"] == password) {
    return c.RenderJson(`{"message": "Success"}`)
  }

  c.Response.Status = 400
  return c.RenderJson(`{"message": "Invalid credentials"}`)
}
