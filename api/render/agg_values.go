package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

func AggTimeInt(items []models.AggTimeInt) (result []*genModels.AggTimeInt) {
	result = make([]*genModels.AggTimeInt, len(items))
	for i, item := range items {
		result[i] = &genModels.AggTimeInt{
			Date:  item.Date.Unix(),
			Value: item.Value,
		}
	}
	return result
}
