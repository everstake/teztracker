package michelson

import (
	script "blockwatch.cc/tzindex/micheline"
	"encoding/json"
	"fmt"
)

func NewBigMapContainer() *BigMapContainer {
	return &BigMapContainer{
		pathMap:  map[string]ContractElement{},
		FinalMap: map[string]interface{}{},
	}
}

type BigMapContainer struct {
	pathMap    map[string]ContractElement
	FinalMap   map[string]interface{}
	SearchFunc searchFunc
}

type searchFunc func(prim *script.Prim, path string)

type ContractElement struct {
	Name string
	Type script.PrimType
}

type Vertex struct {
	visited    bool
	value      *script.Prim
	neighbours []*script.Prim
}

func NewVertex(prim *script.Prim) *Vertex {
	if prim == nil {
		return nil
	}

	return &Vertex{
		visited:    false,
		value:      prim,
		neighbours: prim.Args,
	}
}

func (v *Vertex) Value() interface{} {
	return v.value.Value(0)
}

func (g *BigMapContainer) PathInit(prim *script.Prim, path string) {
	//For now process only values with annotation
	if prim == nil || prim.GetAnno() == "" {
		return
	}

	g.pathMap[path] = ContractElement{
		Name: prim.GetAnno(),
		Type: prim.Type,
	}
}

func (g *BigMapContainer) SearchByPath(prim *script.Prim, path string) {
	if elem, ok := g.pathMap[path]; ok {
		var data interface{}
		switch elem.Type {
		//Array type
		case script.PrimUnaryAnno:
			elems := make([]interface{}, len(prim.Args))
			for key, value := range prim.Args {
				elems[key] = value.Value(0)
			}
			data = elems
		//Simple data
		case script.PrimNullaryAnno:
			data = prim.Value(0)
		}

		g.FinalMap[elem.Name] = data
	}
}

func (g *BigMapContainer) Dfs(vertex *Vertex, path string) {
	if vertex == nil || vertex.visited {
		return
	}

	vertex.visited = true

	g.SearchFunc(vertex.value, path)

	for i, v := range vertex.neighbours {
		vert := NewVertex(v)
		g.Dfs(vert, path+fmt.Sprint(i))
	}
}

func (g *BigMapContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.FinalMap)
}
