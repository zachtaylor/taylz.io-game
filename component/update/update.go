package update

import (
	"taylz.io/types"
)

type T struct {
	Data types.Dict
}

func (t *T) Set(k string, v types.I) {
	if t == nil {
	} else if t.Data == nil {
		t.Data = types.Dict{
			k: v,
		}
		return
	}
	t.Data[k] = v
}

func (t *T) Merge(k, kk string, v types.I) {
	if t == nil {
	} else if t.Data == nil {
		t.Data = types.Dict{
			k: types.Dict{
				kk: v,
			},
		}
	} else if t.Data[k] == nil {
		t.Data[k] = types.Dict{
			kk: v,
		}
	} else if d, ok := t.Data[k].(types.Dict); ok {
		d[kk] = v
	} else {
		panic("update merge failed")
	}
}
