package log

import "math/rand"

type Log struct {
	Id    string
	Value string
}

func NewRandom(id string) Log {
	const allowedLetters = "abcdefg"

	var v string

	for i := 0; i < 10; i++ {
		v += string(allowedLetters[rand.Intn(len(allowedLetters))])
	}

	return Log{
		Id:    id,
		Value: v,
	}
}
