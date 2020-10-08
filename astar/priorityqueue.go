package astar

type priorityQueue []*Nodo

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].f < (*pq)[j].f
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].indice = i
	(*pq)[j].indice = j
}

// Push agrega un *Nodo al final de pq
func (pq *priorityQueue) Push(nodo interface{}) {
	n := len(*pq)
	tempNodo := nodo.(*Nodo)
	tempNodo.indice = n
	*pq = append(*pq, tempNodo)
}

// Pop regresa el ultimop valor del priorityQueue
func (pq *priorityQueue) Pop() interface{} {
	pqVieja := *pq
	n := len(pqVieja)

	tempNodo := pqVieja[n-1] // sacando el ultimo valor
	tempNodo.indice = -1     // como ya no esta en la pq, evitamos problemas
	// asignamos nil a la ultima posicion del pq para evitar problemas de asignacion de memoria
	// (pq es un arrelgo de punteros)
	pqVieja[n-1] = nil

	*pq = (*pq)[:n-1] // sacamos el ultimos valor del pq

	return tempNodo
}
