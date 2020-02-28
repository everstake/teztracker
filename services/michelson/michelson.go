package michelson

import (
	script "blockwatch.cc/tzindex/micheline"
	"encoding/json"
	"fmt"
)

func NewBigMapContainer() BigMapContainer {
	return BigMapContainer{
		pathMap:  map[string]contractElement{},
		finalMap: map[string]interface{}{},
	}
}

type BigMapContainer struct {
	pathMap    map[string]contractElement
	finalMap   map[string]interface{}
	searchFunc searchFunc
}

type searchFunc func(prim *script.Prim, path string)

type contractElement struct {
	Name string
	Type script.PrimType
}

type vertex struct {
	visited    bool
	value      *script.Prim
	neighbours []*script.Prim
}

func newVertex(prim *script.Prim) *vertex {
	if prim == nil {
		return nil
	}

	return &vertex{
		visited:    false,
		value:      prim,
		neighbours: prim.Args,
	}
}

func (v *vertex) Value() interface{} {
	return v.value.Value(0)
}

func (g *BigMapContainer) pathInit(prim *script.Prim, path string) {
	//For now process only values with annotation
	if prim == nil || prim.GetAnno() == "" {
		return
	}

	g.pathMap[path] = contractElement{
		Name: prim.GetAnno(),
		Type: prim.Type,
	}
}

func (g *BigMapContainer) searchByPath(prim *script.Prim, path string) {
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

		g.finalMap[elem.Name] = data
	}
}

func (g *BigMapContainer) InitPath(prim *script.Prim) {
	g.searchFunc = g.pathInit
	g.dfs(newVertex(prim), "")
}

func (g *BigMapContainer) ParseValues(prim *script.Prim) {
	g.searchFunc = g.searchByPath
	g.dfs(newVertex(prim), "")
}

func (g *BigMapContainer) dfs(vertex *vertex, path string) {
	if vertex == nil || vertex.visited {
		return
	}

	vertex.visited = true

	g.searchFunc(vertex.value, path)

	for i, v := range vertex.neighbours {
		vert := newVertex(v)
		g.dfs(vert, path+fmt.Sprint(i))
	}
}

func (g *BigMapContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.finalMap)
}
