package shared

type Message interface {
	Call([]byte) ([]byte, error)
}

type SystemFunction func([]byte) ([]byte, error)

func (sf SystemFunction) Call(payload []byte) ([]byte, error) {
	return sf(payload)
}

type MSI struct {
	Fn  SystemFunction
	In  interface{}
	Out interface{}
}
