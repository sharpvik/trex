package trex

import (
	"reflect"
	"strings"
	"testing"
)

type testFunc = reflect.Value // func(*Suite, *testing.T)

type runner struct {
	t     *testing.T
	suite Suite
}

func runnable(suite Suite) *runner {
	return &runner{
		suite: suite,
	}
}

func (r *runner) run(t *testing.T) {
	r.t = t
	r.tests()
}

func (r *runner) tests() {
	for i, suite := 0, reflect.TypeOf(r.suite); i < suite.NumMethod(); i++ {
		if method := suite.Method(i); strings.HasPrefix(method.Name, "Test") {
			r.t.Run(method.Name, r.test(method.Func))
		}
	}
}

func (r *runner) test(test testFunc) func(*testing.T) {
	return func(t *testing.T) {
		t.Cleanup(r.suite.Teardown)
		r.suite.Setup()
		test.Call([]reflect.Value{
			reflect.ValueOf(r.suite),
			reflect.ValueOf(t),
		})
	}
}
