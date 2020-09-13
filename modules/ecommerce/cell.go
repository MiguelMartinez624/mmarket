package ecommerce

import (
	ecommerce "github.com/miguelmartinez624/mmarket/modules/ecommerce/core"
	"github.com/miguelmartinez624/mmarket/nodos"
)

type EcommerceCell struct {
	module *ecommerce.Module
}

func (e EcommerceCell) Join(net *nodos.NeuralRed) {

}

