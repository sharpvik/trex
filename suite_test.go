package trex_test

import (
	"testing"

	"github.com/sharpvik/trex"
)

type Suite struct {
	trex.BasicSuite
	variable int
}

func TestSuite(t *testing.T) {
	trex.Run(t, &Suite{})
}

func (s *Suite) Setup() {
	s.variable = 42
}

func (s *Suite) Teardown() {
	s.variable = 0
}

func (s *Suite) TestFortyTwo(t *testing.T) {
	if s.variable != 42 {
		t.Fatalf("variable value mismatch: want 42; got %d", s.variable)
	}
}

func (s *Suite) TestIsEven(t *testing.T) {
	if s.variable%2 != 0 {
		t.Fatalf("variable value %d is odd; want even", s.variable)
	}
}
