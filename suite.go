package trex

import "testing"

type Suite interface {
	Setup()
	Teardown()
}

func Run(t *testing.T, suite Suite) {
	runnable(suite).run(t)
}

type BasicSuite struct{}

func (s *BasicSuite) Setup()    {}
func (s *BasicSuite) Teardown() {}
