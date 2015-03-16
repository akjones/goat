package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMessageHandler(t *testing.T) {

	Convey("Returns single message json", t, func() {
		mockWriter := httptest.NewRecorder()
		mockRequest := &http.Request{}
		handleMessage(mockWriter, mockRequest)
		So(mockWriter.Body.String(), ShouldEqual, `"message": {"id": 1, "body": "stuff"}`)
	})

	Convey("returns a json array of messages", t, nil)

}
