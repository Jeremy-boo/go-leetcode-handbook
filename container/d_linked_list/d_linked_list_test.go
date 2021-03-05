package d_linked_list

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	list := Constructor()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.InsertByIndex(4, 1)
	d, err := list.FindByIndex(1)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(d.data)
}
