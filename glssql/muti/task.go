package muti

import (
	"flag"
	"fmt"
	"strconv"
	"sync"
)

type Mutiwork interface {
	InitArgs()
	GetArgs() *sync.Map
	SetArgs(k string, v interface{})
	GetArg(k string) interface{}
	Prepare() error
	PreThread(t, c, p, r uint)
	Work(t, c, p, r uint) uint
	PartThread(t, c, p, r uint)
	EndThread(t, c, p, r uint)
}

type Mutask struct {
	Lck  sync.RWMutex
	args *sync.Map
}

func NewMutask() *Mutask {
	var arg = sync.Map{}
	return &Mutask{
		sync.RWMutex{},
		&arg,
	}
}

func (m *Mutask) GetArgs() *sync.Map {
	return m.args
}

func (m *Mutask) GetArg(k string) interface{} {
	if v, ok := m.args.Load(k); ok {
		return v
	} else {
		return nil
	}
}

func UintVar(args *sync.Map, p *uint, name string, value uint, usage string) {
	flag.UintVar(p, name, value, usage)
	if args == nil {
		panic("UintVar Args is Null")
	}
	args.Store(name, p)
}

func StringVar(args *sync.Map, p *string, name string, value string, usage string) {
	flag.StringVar(p, name, value, usage)
	if args == nil {
		panic("UintVar Args is Null")
	}
	args.Store(name, p)
}

func Atoi(a interface{}) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%v", a))
	return i
}

