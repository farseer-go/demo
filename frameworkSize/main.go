package main

import (
	"github.com/farseer-go/fs"
)

// @test
func main() {
	fs.Initialize[StartupModule]("demo")
}
