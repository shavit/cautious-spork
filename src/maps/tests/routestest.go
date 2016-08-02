package tests

import "github.com/revel/revel/testing"

type RoutesTest struct {
	testing.TestSuite
}

func (t *RoutesTest) Before() {
	println("Set up")
}

func (t *RoutesTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *RoutesTest) After() {
	println("Tear down")
}
