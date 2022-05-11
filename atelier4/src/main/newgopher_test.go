package main

import (
	"formation-go/logger"
	"testing"
)

func TestGopherMessage(t *testing.T) {
	main()
	if !logger.Log.CheckMessage("I am a new Gopher") {
		t.Error("Must show 'I am a new gopher'")
	}
}
