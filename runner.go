package trex

import "testing"

type runner struct {
	suite Suite
}

func runnable(suite Suite) *runner {
	return &runner{suite}
}

func (r *runner) run() {
	for r.test(r.suite.Next()) {
	}
}

func (r *runner) t() *testing.T { return r.suite.T() }

func (r *runner) setup() {
	if err := r.suite.Setup(); err != nil {
		r.t().Fatalf("test case setup failed: %s", err)
	}
}

func (r *runner) teardown() {
	if err := r.suite.Teardown(); err != nil {
		r.t().Fatalf("test case teardown failed: %s", err)
	}
}

func (r *runner) test(test Test) (ok bool) {
	if ok = test != nil; ok {
		r.testWithSetupAndTeardown(test)
	}
	return
}

func (r *runner) testWithSetupAndTeardown(test Test) {
	defer r.teardown()
	r.setup()
	test()
}
