package redis

import (
	"math/rand"
)

type ExpireTimeUtil struct {
	expireTime     int
	maxRandAddTime int
}

func (etu ExpireTimeUtil) GetRandTime() int {
	return rand.Intn(etu.maxRandAddTime) + etu.expireTime
}
