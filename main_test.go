package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Clock(t *testing.T) {
	c, err := NewClock()
	if err != nil {
		t.Error(err)
	}
	if len(c.etag()) == 0 {
		t.Error("failed to generate etag")
	}
}

func Test_ClockHandlerEtag(t *testing.T) {
	var etag string
	testfunc := func(expectedStatusCode int) {
		req, _ := http.NewRequest("GET", "/api/clock", nil)
		req.Header.Add("If-None-Match", etag)
		res := httptest.NewRecorder()
		ClockHandler(res, req)
		if res.Code != expectedStatusCode {
			t.Errorf("expected %v, actual %v", expectedStatusCode, res.Code)
		}
		// save ETag string to etag variable in functional scope
		etag = res.Header().Get("ETag")
	}

	etag = ""
	testfunc(http.StatusOK)
	testfunc(http.StatusNotModified)
	etag = "A"
	testfunc(http.StatusOK)
	testfunc(http.StatusNotModified)
}
