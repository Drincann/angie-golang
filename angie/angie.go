package angie

import (
	"github.com/Drincann/angie-golang/types"
)

func New() types.IApplication {
	return newApplication()
}

type Context = types.Context
