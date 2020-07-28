package marsh

import (
	"fmt"
	"math/rand"
)

type FieldObject struct {
	Index      uint32
	Random     uint64
	Attributes []string
}

func New(index uint32) *FieldObject {
	fielder := &FieldObject{
		Index:      index,
		Random:     rand.Uint64(),
		Attributes: make([]string, 10),
	}
	return fielder
}

func (fielder *FieldObject) AddAttribute(value string) {
	fielder.Attributes = append(fielder.Attributes, value)
}

func (fielder *FieldObject) Compare(other *FieldObject) bool {
	if fielder.Index != other.Index {
		fmt.Sprintf("left: %d != right: %d", fielder.Index, other.Index)
		return false
	}
	if fielder.Random != other.Random {
		fmt.Sprintf("left: %d != right: %d", fielder.Random, other.Random)
		return false
	}
	if len(fielder.Attributes) != len(other.Attributes) {
		fmt.Sprintf("left: %d != right: %d", len(fielder.Attributes), len(other.Attributes))
		return false
	}
	for index := 0; index < len(fielder.Attributes); index++ {
		if fielder.Attributes[index] != other.Attributes[index] {
			fmt.Sprintf("left: %d != right: %d", fielder.Attributes[index], other.Attributes[index])
		}
		return false
	}
	return true
}
