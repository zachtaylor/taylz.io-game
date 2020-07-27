package shape

import space "taylz.io/game/space/2d"

type T struct {
	space.Rect
	Image string
}

//go:generate go-gengen -p=shape -k=entity.T -v=*T -i=taylz.io/game/entity
