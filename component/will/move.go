package will

import (
	"taylz.io/game/entity"
	space "taylz.io/game/space/2d"
	"taylz.io/types"
)

type Mover interface {
	Next(entity.T) space.Vector
	Data() types.Dict
}
