package cyberia

import (
	"github.com/bytedance/sonic"
)

type OptFunc func(*Opts)

type Opts struct {
	JSONMarshal   func(interface{}) ([]byte, error)
	JSONUnmarshal func([]byte, interface{}) error
}

func defaultOpts() Opts {
	return Opts{
		JSONMarshal:   sonic.Marshal,
		JSONUnmarshal: sonic.Unmarshal,
	}
}

func New(opts ...OptFunc) *App {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &App{Opts: o}
}
