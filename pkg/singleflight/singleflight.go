package singleflight

import "golang.org/x/sync/singleflight"

var (
	G = &singleflight.Group{}
)
