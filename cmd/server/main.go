package main

import (
	"github.com/gone-io/gone"
	"template_module/internal"
)

func main() {
	gone.
		Default.
		Loads(internal.Load).
		Serve()
}
