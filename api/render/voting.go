package render

import (
	genModels "github.com/everstake/teztracker/gen/models"
	"github.com/everstake/teztracker/models"
)

// Snapshot renders an app level model to a gennerated OpenAPI model.
func Period(b models.PeriodInfo) *genModels.PeriodInfo {
	return &genModels.PeriodInfo{}
}
