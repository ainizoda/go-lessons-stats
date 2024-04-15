package main

import (
	dv "alif/devices"
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
