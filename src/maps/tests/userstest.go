package tests

import (
	"github.com/revel/revel/testing"
	"io"
	"fmt"
	"bytes"
)

type UsersTest struct {
	testing.TestSuite
}

const (
  email = "user@example.com"
  password = "g00dpassword"
)

func (t *UsersTest) Before() {
	println("Set up")
}

func (t *UsersTest) TestThatUserCanLogin() {
	var jsonData string
	var ioReader io.Reader

	jsonData = fmt.Sprintf(`{
			"username": "%v",
			"password": "%v"
		}`, email, password)
	ioReader = bytes.NewReader([]byte(jsonData))
	t.Post("/api/v1/users/login", "application/json", ioReader)
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *UsersTest) After() {
	println("Tear down")
}
