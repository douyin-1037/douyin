package util

import "math/rand"

type ExpireTimeUtil struct {
	ExpireTime     int
	MaxRandAddTime int
}

func (etu ExpireTimeUtil) GetRandTime() int {
	return rand.Intn(etu.MaxRandAddTime) + etu.ExpireTime
}
