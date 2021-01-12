package tezosnodes

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"io/ioutil"
	"net/http"
	"time"
)

const apiURL = "https://api.tezos-nodes.com/v1"
const routeBakers = "/bakers"

type (
	API struct {
		client *http.Client
	}
	Baker struct {
		Rank       int     `json:"rank"`
		Name       string  `json:"name"`
		Address    string  `json:"address"`
		Fee        float64 `json:"fee"`
		Yield      float64 `json:"yield"`
		Efficiency float64 `json:"efficiency"`
		Freespace  int64   `json:"freespace"`
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
	thirdPartyBakers = make([]models.ThirdPartyBaker, len(bakers))
	for i, b := range bakers {
		thirdPartyBakers[i] = models.ThirdPartyBaker{
			Number:            b.Rank,
			Name:              b.Name,
			Address:           b.Address,
			Yield:             b.Yield,
			Fee:               b.Fee,
			AvailableCapacity: b.Freespace * 1e6,
			Efficiency:        b.Efficiency,
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
