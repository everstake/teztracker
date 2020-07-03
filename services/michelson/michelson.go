package michelson

import (
	script "blockwatch.cc/tzindex/micheline"
	"encoding/json"
	"fmt"
)

func NewBigMapContainer() BigMapContainer {
	return BigMapContainer{
		pathMap:     map[string]contractElement{},
		namePathMap: map[string]string{},
		finalMap:    map[string]interface{}{},
	}
}

type BigMapContainer struct {
	pathMap     map[string]contractElement
	namePathMap map[string]string
	finalMap    map[string]interface{}
	searchFunc  searchFunc
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

	g.namePathMap[prim.GetAnno()] = path
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
			if len(prim.Args) == 0 {
				data = prim.Value(0)
			} else {
				data = prim.Args[0].Value(0)
			}

		case script.PrimBinaryAnno:

			//To convert args map to array
			switch prim.OpCode {
			case script.D_LEFT, script.D_RIGHT:
				if prim.Args[0].OpCode == script.D_PAIR {
					data = []interface{}{prim.Args[0].Args[0].Value(0), prim.Args[0].Args[1].Value(0)}
				}
			default:
				data = prim.Value(0)
			}

		default:

		}

		g.finalMap[elem.Name] = data
	}
}

func (g *BigMapContainer) InitPath(prim *script.Prim) {
	g.searchFunc = g.pathInit
	g.dfs(newVertex(prim), "")
}

func (g *BigMapContainer) FlushValues() {
	g.finalMap = map[string]interface{}{}
}

func (g *BigMapContainer) ParseValues(entrypoint string, prim *script.Prim) {
	if prim == nil {
		return
	}
	g.searchFunc = g.searchByPath
	path := ""
	if prim.OpCode == script.D_RIGHT {
		path = "1"
	}

	if value, ok := g.namePathMap[entrypoint]; ok {
		path = value
	}

	g.dfs(newVertex(prim), path)
}

func (g *BigMapContainer) dfs(vertex *vertex, path string) {
	if vertex == nil || vertex.visited {
		return
	}

	vertex.visited = true

	g.searchFunc(vertex.value, path)

	for i, v := range vertex.neighbours {
		pathIndex := i
		vert := newVertex(v)

		if i == 0 {
			if vertex.neighbours[0].OpCode == script.D_RIGHT {
				pathIndex = 1
			}

		}

		g.dfs(vert, path+fmt.Sprint(pathIndex))
	}
}

func (g *BigMapContainer) MarshalPath() ([]byte, error) {
	return json.Marshal(g.pathMap)
}

func (g *BigMapContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.finalMap)
}
