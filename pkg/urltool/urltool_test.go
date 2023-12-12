package urltool

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetBaseURL(t *testing.T) {
	convey.Convey("WithPaths", t, func() {
		url := "https://www.baidu.com/query"
		got, err := GetBaseURL(url)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, "query")
	})

	convey.Convey("WithQueryArgs", t, func() {
		url := "https://www.baidu.com/query?name=123&age=3"
		got, err := GetBaseURL(url)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, "query")
	})

	convey.Convey("WithoutPaths", t, func() {
		url := "www.baidu.com"
		got, err := GetBaseURL(url)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, "www.baidu.com")
	})

	convey.Convey("WithoutPaths", t, func() {
		url := "https://www.baidu.com"
		got, err := GetBaseURL(url)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, "www.baidu.com")
	})
}
