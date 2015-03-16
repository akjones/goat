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
		mockRequest := &http.Request{}                                                     // A blank request is sufficient for this test
		handleMessage(mockWriter, mockRequest)                                             // The actual function under test
		So(mockWriter.Body.String(), ShouldEqual, `"message": {"id": 1, "body": "stuff"}`) // Because the function doesn't return anything, we're actually testing side-effects
	})

	Convey("returns a json array of messages", t, nil)

}
