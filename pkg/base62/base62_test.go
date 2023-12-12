package base62

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	convey.Convey("BasicEncode", t, func() {
		num := uint64(201314)
		got := Encode(num)
		convey.So(got, convey.ShouldEqual, "Qn0")
	})

	convey.Convey("EncodeZero", t, func() {
		num := uint64(0)
		got := Encode(num)
		convey.So(got, convey.ShouldEqual, "0")
	})
}

func TestDecode(t *testing.T) {
	convey.Convey("BasicDecode", t, func() {
		str := "Qn0"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(201314))
	})

	convey.Convey("DecodeZero", t, func() {
		str := "0"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(0))
	})

	convey.Convey("DecodeMultiZero", t, func() {
		str := "0000"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(0))
	})

	convey.Convey("DecodeFail", t, func() {
		str := "Qn-"
		_, err := Decode(str)
		convey.So(err, convey.ShouldEqual, errInvalidChar)
	})
}
