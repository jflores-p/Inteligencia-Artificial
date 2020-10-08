package astar

var (
	mapa = [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
)

// Posicion en un tabero
type Posicion struct {
	X, Y int
}

// Nodo de estados del camino
type Nodo struct {
	VieneDe  *Nodo
	f, g, h  float32
	visitado bool
	indice   int

	Posicion
}

//Igual regresa verdadero si n y otro estan en la misma posicion
func (n *Nodo) Igual(otro *Nodo) bool {
	return n.Posicion.X == otro.Posicion.X &&
		n.Posicion.Y == otro.Posicion.Y
}

// Astar algoritmo de busqueda
func Astar(inicio, fin Nodo, mapa [][]int) []Nodo {
	return nil
}
