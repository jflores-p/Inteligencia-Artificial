package astar

import "fmt"

var (
	mapa1 = map[string]map[string]int{
		"alabama": map[string]int{
			"1": 0xf}}
)

func a() {

	fmt.Println(mapa1["alabama"])
}
