package main

import (
	"github.com/gone-io/gone/v2"
	"template_module/internal"
)

func main() {
	gone.
		Loads(internal.Load).
		Serve()
}
