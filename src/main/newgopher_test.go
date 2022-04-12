package main

import (
	"formation-go/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGopherMessage(t *testing.T) {
	main()
	assert.True(t, logger.Log.CheckMessage("I am a new Gopher"))
}
