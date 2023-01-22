package trex

import (
	"reflect"
	"strings"
	"testing"
)

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
	for i, suiteT := 0, reflect.TypeOf(r.suite); i < suiteT.NumMethod(); i++ {
		if method := suiteT.Method(i); strings.HasPrefix(method.Name, "Test") {
			r.test(method.Func)
		}
	}
}

func (r *runner) test(test Test) {
	defer r.teardown()
	r.setup()
	test.Call([]reflect.Value{
		reflect.ValueOf(r.suite),
		reflect.ValueOf(r.t),
	})
}

func (r *runner) setup() {
	if err := r.suite.Setup(); err != nil {
		r.t.Fatalf("test case setup failed: %s", err)
	}
}

func (r *runner) teardown() {
	if err := r.suite.Teardown(); err != nil {
		r.t.Fatalf("test case teardown failed: %s", err)
	}
}
