package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// START1 OMIT
func TestIntegerExample(t *testing.T) {
	Convey("Subject: Integer incrementation and decrementation", t, func() {
		var x int

		Convey("Given a starting integer value", func() {
			x = 42

			Convey("When incremented", func() {
				x++

				Convey("The value should be greater by one", func() {
					So(x, ShouldEqual, 43)
				})
				Convey("The value should NOT be what it used to be", func() {
					So(x, ShouldNotEqual, 42)
				})
			})
		})
	})
}

// END1 OMIT

func TestBoring(t *tesing.T) {
	// START2 OMIT
	if foo < bar || baz != "abc" {
		t.Error("Test failed...")
	}
	// END2 OMIT
}

func TestSmart(t *testing.T) {
	// START3 OMIT
	So(foo, ShouldBeGreaterThanOrEqualTo, bar)
	So(baz, ShouldEqual, "abc")
	// END3 OMIT
}
