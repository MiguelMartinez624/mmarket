package tests

import (
	"context"
	"testing"

	"github.com/miguelmartinez624/mmarket/modules/stores/core"
	"github.com/miguelmartinez624/mmarket/modules/stores/core/domains/stores"
)

var store MuckStoreRepository = MuckStoreRepository{}
var pRepo MuckProductRepository = MuckProductRepository{}
var ctx = context.TODO()
var profileMuck MuckProfileModule = MuckProfileModule{}

type ModuleTestFunction func(s *core.Module, t *testing.T)

type TestCase struct {
	Name     string
	Callback ModuleTestFunction
}

func ModuleSuite(s *core.Module, t *testing.T) {
	tc := []TestCase{
		{Name: "Create store success", Callback: CreateStoreSuccess},
		{Name: "Attemp Create without profileID", Callback: CreateStoreFailMissinProfile},
	}

	for _, tCase := range tc {
		t.Run(tCase.Name, func(t *testing.T) {
			tCase.Callback(s, t)
		})
	}
}

func CreateStoreSuccess(m *core.Module, t *testing.T) {
	s := stores.Store{
		Name:      "My MArket",
		ProfileID: "ID",
	}

	_, err := m.CreateStore(ctx, &s)
	if err != nil {
		t.Errorf("Error wasnet expected %v", err)
	}
}

func CreateStoreFailMissinProfile(m *core.Module, t *testing.T) {
	s := stores.Store{
		Name: "My MArket",
	}

	_, err := m.CreateStore(ctx, &s)
	if err == nil {
		t.Error("It should have returned a error")
	}

	switch te := err.(type) {
	case stores.MissinField:
		if te.Field != "ProfileID" {
			t.Errorf("expected missing  field [%v]  to be [ProfileID]", te.Field)
		}
		return
	default:
		t.Errorf("expected [%v] to be [MissinField]", te)
	}
}

func TestModule(t *testing.T) {

	module := core.NewModule(&store, &pRepo)
	module.ConnectToProfiles(profileMuck)
	ModuleSuite(module, t)
}
