package union

import (
	"fmt"
)

type Union struct {
	id []int
}

func (union Union) Connected(id1, id2 int) bool {
	if union.id[id1] == union.id[id2] {
		fmt.Println("ids are connected", id1, id2)
		return true
	}
	fmt.Println("ids not connected", id1, id2)
	return false
}

func (union *Union) Unite(id1, id2 int) {
	for index, value := range union.id {
		if value == union.id[id2] {
			union.id[index] = union.id[id1]
		}
	}
}

func main() {
	union := Union{
		id: []int{1, 2, 3, 4, 5, 6, 7, 8},
	}
	union.Unite(1, 2)
	union.Connected(0, 1)
	union.Connected(1, 2)
	fmt.Println(union)
}
