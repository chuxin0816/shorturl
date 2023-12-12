package md5

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSum(t *testing.T) {
	convey.Convey("md5.Sum", t, func() {
		data := []byte("123456")
		got := Sum(data)
		convey.So(got, convey.ShouldEqual, "e10adc3949ba59abbe56e057f20f883e")
	})
}
