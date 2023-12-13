package randtime

import (
	"math/rand"
	"time"
)

const randFactor = 30

// GetRandTime 返回0-30分钟之间的随机时间
func GetRandTime() time.Duration {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return time.Duration(r.Intn(randFactor)) * time.Minute
}
