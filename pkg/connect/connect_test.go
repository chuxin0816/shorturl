package connect

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	convey.Convey("GetPass", t, func() {
		url := "https://www.baidu.com"
		got := Get(url)
		convey.So(got, convey.ShouldBeTrue)
	})

	convey.Convey("GetFail", t, func() {
		url := "/1111/"
		got := Get(url)
		convey.So(got, convey.ShouldBeFalse)
	})
}
