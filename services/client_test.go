package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	c1, err := NewClient()
	assert.Nil(t, err)
	assert.Equal(t, "https://apps.bea.gov/api/data/", c1.uri.String())

	c2, err := NewClient("https://chain.link")
	assert.Nil(t, err)
	assert.Equal(t, "https://chain.link", c2.uri.String())

	_, err = NewClient("https ://extra.whitespace")
	assert.Error(t, err)
}
