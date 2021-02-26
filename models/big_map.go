package models

//DB model
type BigMapContent struct {
	BigMapId           uint64 //big_map_id
	Key                string //key
	KeyHash            string //key_hash
	OperationGroupHash string //operation_group_id
	Value              string //value
}
