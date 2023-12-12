package base62

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestEncode(t *testing.T) {
	convey.Convey("BasicEncode", t, func() {
		num := uint64(201314)
		got := Encode(num)
		convey.So(got, convey.ShouldEqual, "NIn")
	})

	convey.Convey("EncodeZero", t, func() {
		num := uint64(0)
		got := Encode(num)
		convey.So(got, convey.ShouldEqual, "n")
	})

	convey.Convey("EncodeMax", t, func() {
		num := uint64(18446744073709551615)
		got := Encode(num)
		convey.So(got, convey.ShouldEqual, "htvw01rgwGe")
	})
}

func TestDecode(t *testing.T) {
	convey.Convey("BasicDecode", t, func() {
		str := "NIn"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(201314))
	})

	convey.Convey("DecodeZero", t, func() {
		str := "n"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(0))
	})

	convey.Convey("DecodeMultiZero", t, func() {
		str := "nnnn"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(0))
	})

	convey.Convey("DecodeMax", t, func() {
		str := "htvw01rgwGe"
		got, err := Decode(str)
		convey.So(err, convey.ShouldBeNil)
		convey.So(got, convey.ShouldEqual, uint64(18446744073709551615))
	})

	convey.Convey("DecodeFail", t, func() {
		str := "Qn-"
		_, err := Decode(str)
		convey.So(err, convey.ShouldEqual, errInvalidChar)
	})
}
