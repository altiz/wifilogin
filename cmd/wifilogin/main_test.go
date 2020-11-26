package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wifilogin/cmd/wifilogin/models"
	"wifilogin/version"

	logs "github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

type data struct {
	User string `json:"user"`
}
type data1 struct {
	Status int
}

func TestMain(t *testing.T) {
	router := setupRouter()

	Convey("Users endpoints should respond correctly", t, func() {
		Convey("Test Index", func() {
			// it's safe to ignore error here, because we're manually entering URL
			req, _ := http.NewRequest("GET", "http://localhost:5000/billing/api/v1/", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
			body := strings.TrimSpace(w.Body.String())
			So(body, ShouldEqual, "ok")
		})

		Convey("test home", func() {
			//send message
			req, err := http.NewRequest("POST", "http://localhost:5000/billing/api/v1/home", nil)
			So(err, ShouldBeNil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
			//resp
			body := w.Body.Bytes()
			var obj *models.TVersion
			if err := json.Unmarshal(body, &obj); err != nil {
				panic(err)
			}
			logs.WithFields(logs.Fields{
				"commit":     version.Commit,
				"build time": version.BuildTime,
				"release":    version.Release,
			}).Info(obj.BuildTime)
			So(obj.BuildTime, ShouldEqual, version.BuildTime)
		})

		Convey("Test API", func() {
			//send message
			user := &data{User: "test"}
			dat, err := json.Marshal(user)
			So(err, ShouldBeNil)
			buf := bytes.NewBuffer(dat)
			req, err := http.NewRequest("POST", "http://localhost:5000/billing/api/v1/test", buf)
			So(err, ShouldBeNil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
			//resp
			body := w.Body.Bytes()
			var obj *data1

			if err := json.Unmarshal(body, &obj); err != nil {
				panic(err)
			}
			So(obj.Status, ShouldEqual, 0)
		})
	})
}
