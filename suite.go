package trex

import "testing"

type Suite interface {
	Setup() error
	Teardown() error
	Next() Test
}

type Test func(*testing.T)

func Run(t *testing.T, suite Suite) {
	runnable(suite).run(t)
}

type BasicSuite struct {
	queue[Test]
}

func (s *BasicSuite) Setup() error    { return nil }
func (s *BasicSuite) Teardown() error { return nil }
