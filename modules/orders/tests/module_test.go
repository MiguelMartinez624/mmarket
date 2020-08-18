package tests

import (
	"context"
	"testing"

	"github.com/miguelmartinez624/mmarket/modules/orders/core"
)

var ctx = context.Background()

type ModuleTestFunction func(s *core.Module, t *testing.T)

type TestCase struct {
	Name     string
	Callback ModuleTestFunction
}

func ModuleSuite(s *core.Module, t *testing.T) {
	tc := []TestCase{}

	for _, tCase := range tc {
		t.Run(tCase.Name, func(t *testing.T) {
			tCase.Callback(s, t)
		})
	}
}

func TestModule(t *testing.T) {

	module := core.NewModule()
	ModuleSuite(module, t)
}
