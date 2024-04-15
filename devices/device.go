package devices

type Cpu struct {
	Prod string
	Core int
}

type Pc struct {
	Cpu
}

func (p Pc) On() {
	println("[ON] starting computer ...")
}

func (p Pc) Off() {
	println("[OFF] shutting down ...")
}