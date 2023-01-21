package trex

import "testing"

type Suite interface {
	T() *testing.T
	Setup() error
	Teardown() error
	Next() Test
}

func Run(suite Suite) {
	runnable(suite).run()
}
