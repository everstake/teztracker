package mytezosbaker

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const apiURL = "https://api.mytezosbaker.com/v1"
const routeBakers = "/bakers/"

type (
	API struct {
		client *http.Client
	}
	BakerResponse struct {
		Bakers []Baker `json:"bakers"`
	}
	Baker struct {
		Rank                int     `json:"rank"`
		BakerName           string  `json:"baker_name"`
		DelegationCode      string  `json:"delegation_code"`
		Fee                 float64 `json:"fee,string"`
		BakerEfficiency     float64 `json:"baker_efficiency"`
		Logo                string  `json:"logo"`
		AvailableCapacity   float64 `json:"available_capacity,string"`
		NominalStakingYield string  `json:"nominal_staking_yield"`
		RealStakingYield    string  `json:"real_staking_yield"`
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
	var resp BakerResponse
	err = api.get(routeBakers, &resp)
	if err != nil {
		return nil, fmt.Errorf("get: %s", err.Error())
	}
	for _, b := range resp.Bakers {
		strs := strings.Split(b.NominalStakingYield, " ")
		yield, _ := strconv.ParseFloat(strs[0], 64)
		thirdPartyBakers = append(thirdPartyBakers, models.ThirdPartyBaker{
			Number:            b.Rank,
			Name:              b.BakerName,
			Address:           b.DelegationCode,
			Yield:             yield,
			Fee:               b.Fee / 100,
			Efficiency:        b.BakerEfficiency,
			AvailableCapacity: int64(b.AvailableCapacity * 1e6),
		})
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
