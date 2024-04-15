package main

import (
	dv "github.com/ainizoda/go-lessons/v3/devices"
)

func main() {
	pc := dv.Pc{
		Cpu: dv.Cpu{
			Core: 1,
			Prod: "intel",
		},
	}

	pc.On()
}
