package main

import (
	"testing"

	"github.com/linkpoolio/bridges/bridge"
	"github.com/stretchr/testify/assert"
)

func TestBea_Run(t *testing.T) {
	bea := Bea{}
	data := map[string]interface{}{}
	query, _ := bridge.ParseInterface(data)
	obj, err := bea.Run(bridge.NewHelper(query))
	assert.Nil(t, err)

	resp, ok := obj.(map[string]interface{})
	assert.True(t, ok)

	val, ok := resp["result"]
	assert.True(t, ok)
	assert.Greater(t, val, 0.0)
}
