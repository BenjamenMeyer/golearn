package marsh

import (
	"math/rand"
)

type FieldObject struct {
	Index   uint32
	Random  uint64
}

func New(index uint32) *FieldObject {
	fielder := &FieldObject{
		Index: index,
		Random: rand.Uint64(),
	}
	return fielder
}
