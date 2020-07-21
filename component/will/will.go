package will

import "taylz.io/game/component/update"

type T struct {
	name string
	move Mover
}

func (t *T) Name() string { return t.name }
func (t *T) Move() Mover  { return t.move }

func (t *T) UpdateName(u *update.T, name string) {
	t.name = name
	u.Set("a", name)
}

func (t *T) UpdateMove(u *update.T, move Mover) {
	t.move = move
	u.Set("v", move.Data())
}
