package main

type Room struct {
	id         int64
	createTime int32
	players    map[int64]bool
}

func (r Room) Len() int {
	panic("not implemented")
}

func (r Room) Less(i int, j int) bool {
	panic("not implemented")
}

func (r Room) Swap(i int, j int) {
	panic("not implemented")
}
