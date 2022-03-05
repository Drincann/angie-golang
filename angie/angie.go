package angie

import (
	"github.com/Drincann/angie-golang/types"
	"github.com/Drincann/angie-golang/webContext"
)

// export
func New() types.IApplication {
	return newApplication()
}

type Context = webContext.WebContext
