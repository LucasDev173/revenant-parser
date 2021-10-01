package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // parser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {
		r, err := parse((testData.Input))
		if testData.Success != true {
			assert.Equal(t, err != nil, testData.Success)
		} else {
			assert.Equal(t, err == nil, testData.Success)
		}
		if testData.Type != "" {
			assert.Equal(t, r.Type, testData.Type)
		}
		if testData.Length != 0 {
			assert.Equal(t, r.Length, testData.Length)
		}
		if testData.Value != "" {
			assert.Equal(t, r.Value, testData.Value)
		}
	}
}
