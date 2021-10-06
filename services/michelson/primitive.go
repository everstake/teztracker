package michelson

import script "blockwatch.cc/tzgo/micheline"

type BigMap struct {
	Code    script.Prim `json:"code"`
	Storage script.Prim `json:"storage"`
}
