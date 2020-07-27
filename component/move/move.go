package move

import (
	"taylz.io/game/entity"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type T struct {
	I
}

type I interface {
	Next(entity.T) space.Vector
	Data() types.Dict
}

//go:generate go-gengen -p=move -k=entity.T -v=*T -i=taylz.io/game/entity
