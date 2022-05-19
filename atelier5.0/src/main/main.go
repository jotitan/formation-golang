package main

import (
	"formation-go/logger"
	"formation-go/tools"
)

func main() {
	logger.Log.Println("I am a new Gopher")
	from := "~/img_test.jpg"
	to := "~/img_resize.jpg"
	tools.LocalResize{}.Resize(from, to, 300, 300)
}
