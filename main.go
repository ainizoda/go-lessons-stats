package main

import (
	dv "github.com/ainizoda/goLessons/v3/devices"
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
