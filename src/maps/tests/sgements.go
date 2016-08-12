package tests

import "github.com/revel/revel/testing"

type SegmentsTest struct {
	testing.TestSuite
}

func (t *SegmentsTest) Before() {
	println("Set up")
}

func (t *SegmentsTest) TestThatIndexPageWorks() {
	t.Get("/api/v1/segments/explore")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *SegmentsTest) After() {
	println("Tear down")
}
