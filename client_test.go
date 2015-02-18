package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoatClient(t *testing.T) {

	Convey("Chat client", t, func() {

		Convey("should return all unseen messages", func() {
			So(GetMessage(), ShouldEqual, "something")
		})

		Convey("should allow me to send a message", func() {
			So(SendMessage("message blah"), ShouldBeNil)
		})

		Convey("should return the last sent message", func() {
			SendMessage("hello")
			So(GetMessage(), ShouldEqual, "hello")
		})

	})

}
