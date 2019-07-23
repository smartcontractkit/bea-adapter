package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDpcergAvg(t *testing.T) {
	avg, err := GetDpcergAvg(3)
	assert.Nil(t, err)
	assert.NotEqual(t, 0.0, avg)
}
