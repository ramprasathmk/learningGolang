package structs

import "fmt"

type MyBoxItem struct {
	Name string
}

type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) []MyBoxItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func RunStructure() {
	item1 := MyBoxItem{Name: "Test Item 1"}
	item2 := MyBoxItem{Name: "Test Item 2"}

	box := MyBox{}

	box.AddItem(item1)
	box.AddItem(item2)

	fmt.Printf("length of box: %v, capacity: %v\n", len(box.Items), cap(box.Items))
	fmt.Println("Box items: ", box.Items)
}
