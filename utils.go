package moxy

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func pick(array []string) (string, error) {
	size := len(array)
	switch size {
	case 0:
		return "", errors.New(fmt.Sprintf("empty"))
	case 1:
		return array[0], nil
	default:
		rand.Seed(time.Now().Unix())
		return array[rand.Intn(size)], nil
	}
}
