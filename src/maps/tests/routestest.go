package tests

import "github.com/revel/revel/testing"

type RoutesTest struct {
	testing.TestSuite
}

func (t *RoutesTest) Before() {
	println("Set up")
}

func (t *RoutesTest) TestThatIndexPageWorks() {
	t.Get("/api/v1/routes")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *RoutesTest) After() {
	println("Tear down")
}
