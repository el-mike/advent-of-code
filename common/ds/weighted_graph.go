package ds

type Vertex[T any] struct {
	ID     string
	Data   T
	Weight int
}

func (v *Vertex[T]) Clone() *Vertex[T] {
	return &Vertex[T]{
		ID:     v.ID,
		Data:   v.Data,
		Weight: v.Weight,
	}
}

type WeightedGraph[T any] struct {
	AdjacencyList map[string]map[string]*Vertex[T]
}

func NewWeightedGraph[T any]() *WeightedGraph[T] {
	return &WeightedGraph[T]{
		AdjacencyList: map[string]map[string]*Vertex[T]{},
	}
}

func (wg *WeightedGraph[T]) AddVertex(v *Vertex[T]) {
	if _, ok := wg.AdjacencyList[v.ID]; !ok {
		wg.AdjacencyList[v.ID] = map[string]*Vertex[T]{}
	}
}

func (wg *WeightedGraph[T]) AddEdge(vertex, neighbor *Vertex[T], weight int) {
	if _, ok := wg.AdjacencyList[vertex.ID][neighbor.ID]; !ok {
		// We clone the vertices to avoid mutating Weight value on already
		// existing vertices.
		neighbor = neighbor.Clone()
		neighbor.Weight = weight

		wg.AdjacencyList[vertex.ID][neighbor.ID] = neighbor
	}

	if _, ok := wg.AdjacencyList[neighbor.ID][vertex.ID]; !ok {
		vertex = vertex.Clone()
		vertex.Weight = weight
		wg.AdjacencyList[neighbor.ID][vertex.ID] = vertex
	}
}
