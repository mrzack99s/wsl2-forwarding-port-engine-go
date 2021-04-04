package main

import (
	"github.com/mrzack99s/wsl2-forwarding-port-engine-go/pkgs/runtimes"
)

func main() {
	runtimes.Parse()
	runtimes.MGMTServe()
}
