package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type data struct {
	Msisdn string `json:"msisdn"`
}
type data1 struct {
	login  string
	passwd string
}

func TestMain(t *testing.T) {
	router := setupRouter()

	Convey("Users endpoints should respond correctly", t, func() {
		Convey("Test Index", func() {
			// it's safe to ignore error here, because we're manually entering URL
			req, _ := http.NewRequest("GET", "http://localhost:5000/billing/api/v1/set-msisdn_test", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
			body := strings.TrimSpace(w.Body.String())
			So(body, ShouldEqual, "ok")
		})

		/*Convey("test home", func() {
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
		})*/

		Convey("Test API", func() {
			//send message
			beginTime_ := time.Now().UnixNano()
			msisdn := &data{Msisdn: "9872305570"}
			dat, err := json.Marshal(msisdn)
			So(err, ShouldBeNil)
			buf := bytes.NewBuffer(dat)
			req, err := http.NewRequest("POST", "http://localhost:5000/billing/api/v1/set-msisdn_test", buf)
			So(err, ShouldBeNil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			timeCount := time.Now().UnixNano() - beginTime_
			So(w.Code, ShouldEqual, http.StatusOK)
			//resp
			body := w.Body.Bytes()
			var obj *data1

			if err := json.Unmarshal(body, &obj); err != nil {
				panic(err)
			}
			var l int = 0
			if timeCount > 3 {
				l = 1
			}
			So(l, ShouldEqual, 0)
		})
	})
}
