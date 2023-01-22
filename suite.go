package trex

import (
	"reflect"
	"testing"
)

type Suite interface {
	Setup() error
	Teardown() error
}

type Test = reflect.Value // func(*Suite, *testing.T)

func Run(t *testing.T, suite Suite) {
	runnable(suite).run(t)
}

type BasicSuite struct{}

func (s *BasicSuite) Setup() error    { return nil }
func (s *BasicSuite) Teardown() error { return nil }
