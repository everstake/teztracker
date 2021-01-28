package bakingbad

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

const (
	apiURL      = "https://api.baking-bad.org/v2"
	noDataField = "no_data"
	routeBakers = "/bakers?health=active"
)

type (
	API struct {
		client *http.Client
	}
	Baker struct {
		Address           string  `json:"address"`
		Name              string  `json:"name"`
		Logo              string  `json:"logo"`
		Balance           float64 `json:"balance"`
		StakingBalance    float64 `json:"stakingBalance"`
		StakingCapacity   float64 `json:"stakingCapacity"`
		MaxStakingBalance float64 `json:"maxStakingBalance"`
		FreeSpace         float64 `json:"freeSpace"`
		Fee               float64 `json:"fee"`
		MinDelegation     float64 `json:"minDelegation"`
		EstimatedRoi      float64 `json:"estimatedRoi"`
		PayoutAccuracy    string  `json:"payoutAccuracy"`
		InsuranceCoverage float64 `json:"insuranceCoverage"`
	}
)

func New() *API {
	return &API{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (api *API) GetBakers() (thirdPartyBakers []models.ThirdPartyBaker, err error) {
	var bakers []Baker
	err = api.get(routeBakers, &bakers)
	if err != nil {
		return nil, fmt.Errorf("get: %s", err.Error())
	}
	sort.Slice(bakers, func(i, j int) bool {
		if bakers[i].PayoutAccuracy == noDataField {
			return false
		}
		if bakers[j].PayoutAccuracy == noDataField {
			return true
		}
		if bakers[i].InsuranceCoverage != bakers[j].InsuranceCoverage {
			return bakers[i].InsuranceCoverage > bakers[j].InsuranceCoverage
		}
		return bakers[i].EstimatedRoi > bakers[j].EstimatedRoi
	})
	thirdPartyBakers = make([]models.ThirdPartyBaker, len(bakers))
	for i, b := range bakers {
		stakingBalance := int64(b.StakingBalance * 1e6)
		thirdPartyBakers[i] = models.ThirdPartyBaker{
			Number:            i + 1,
			Name:              b.Name,
			Address:           b.Address,
			Yield:             b.EstimatedRoi,
			StakingBalance:    stakingBalance,
			Fee:               b.Fee,
			AvailableCapacity: int64(b.FreeSpace * 1e6),
			PayoutAccuracy:    b.PayoutAccuracy,
		}
	}
	return thirdPartyBakers, nil
}

func (api *API) get(endpoint string, dst interface{}) error {
	url := fmt.Sprintf("%s%s", apiURL, endpoint)
	resp, err := api.client.Get(url)
	if err != nil {
		return fmt.Errorf("client.Get: %s", err.Error())
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll: %s", err.Error())
	}
	err = json.Unmarshal(data, dst)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %s", err.Error())
	}
	return nil
}
