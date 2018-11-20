package jsonrpc


type EngineRPC struct {
	counter int
}

func (t *EngineRPC) Add(arg int) int {
	t.counter += arg
	return t.counter
}

func (t *EngineRPC) Sub(arg int) int {
	t.counter += arg
	return t.counter
}