package trex

import "testing"

type BasicSuite struct {
	t *testing.T
	*queue[Test]
}

type Test func()

func Basic(t *testing.T) BasicSuite {
	return BasicSuite{
		t:     t,
		queue: newQueue[Test](),
	}
}

func (s *BasicSuite) T() *testing.T   { return s.t }
func (s *BasicSuite) Setup() error    { return nil }
func (s *BasicSuite) Teardown() error { return nil }
