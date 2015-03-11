package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoatServer(t *testing.T) {

	Convey("Given a get request to /messages", t, func() {
		var msgs []string
		resp, _ := http.Get("http://localhost:1234/messages")
		decoder := json.NewDecoder(resp.Body)
		_ = decoder.Decode(&msgs)

		Convey("it returns a list of all messages", func() {
			So(len(msgs), ShouldBeGreaterThan, 0)
		})

		Convey("Given a list of messages", func() {
			resp, _ := http.Get("http://localhost:1234/messages/" + msgs[0])
			defer resp.Body.Close()
			b, _ := ioutil.ReadAll(resp.Body)

			Convey("it returns the message content", func() {
				So(string(b), ShouldEqual, `"hello"`)
			})
		})

	})

}
