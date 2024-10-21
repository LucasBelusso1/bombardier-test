package main

import (
	"github.com/LucasBelusso1/bombardier-test/internal/api/chiserver"
	"github.com/LucasBelusso1/bombardier-test/internal/api/ginserver"
	"github.com/LucasBelusso1/bombardier-test/internal/api/gorillaserver"
	"github.com/LucasBelusso1/bombardier-test/internal/api/standardserver"
)

func main() {
	go standardserver.Start()
	go gorillaserver.Start()
	go ginserver.Start()
	chiserver.Start()
}
