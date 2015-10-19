package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResultMap map[string]int

type TestObj struct {
	TestStrings []string
	Result      ResultMap
}

func TestProcessLines(t *testing.T) {
	assert := assert.New(t)
	cases := []TestObj{
		TestObj{TestStrings: []string{"yo!", "yo.", "yo", "Yo?"}, Result: ResultMap{"yo": 4}},
	}
	for _, testCase := range cases {
		res := ProcessLines(testCase.TestStrings)
		assert.Equal(ResultMap(res), testCase.Result)
	}
}
