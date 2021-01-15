package tzstats

import (
	"encoding/json"
	"fmt"
	"github.com/everstake/teztracker/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"
)

const apiURL = "https://api.tzstats.com"
const routeIncome = "/tables/income"
const routeTip = "/explorer/tip"
const routeAccount = "/explorer/account/%s"

type (
	API struct {
		client *http.Client
	}
	Tip struct {
		Cycle int64 `json:"cycle"`
	}
	Account struct {
		Address          string  `json:"address"`
		DelegatedBalance float64 `json:"delegated_balance"`
		StakingBalance   float64 `json:"staking_balance"`
		BlocksBaked      int64   `json:"blocks_baked"`
		BlocksMissed     int64   `json:"blocks_missed"`
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
	var tip Tip
	err = api.get(routeTip, nil, &tip)
	if err != nil {
		return nil, fmt.Errorf("can`t get tip: %s", err.Error())
	}
	params := url.Values{}
	params.Add("columns", "address,rolls")
	params.Add("limit", "1000")
	params.Add("cycle", fmt.Sprintf("%d", tip.Cycle))
	params.Add("order", "desc")
	var incomes [][]interface{}
	err = api.get(routeIncome, params, &incomes)
	if err != nil {
		return nil, fmt.Errorf("can`t get income: %s", err.Error())
	}
	sort.Slice(incomes, func(i, j int) bool {
		if len(incomes[i]) < 2 || len(incomes[j]) < 2 {
			return false
		}
		return incomes[i][1].(float64) > incomes[j][1].(float64)
	})
	if len(incomes) > 100 {
		incomes = incomes[:100]
	}
	thirdPartyBakers = make([]models.ThirdPartyBaker, len(incomes))
	for i, income := range incomes {
		if len(income) == 0 {
			continue
		}
		address := income[0].(string)
		if address == "" {
			continue
		}
		var acc Account
		err = api.get(fmt.Sprintf(routeAccount, address), nil, &acc)
		if err != nil {
			return nil, fmt.Errorf("can`t get account: %s", err.Error())
		}
		var efficiency float64
		if acc.BlocksBaked > 0 {
			efficiency = float64(acc.BlocksBaked) / float64(acc.BlocksBaked+acc.BlocksMissed)
		}
		thirdPartyBakers[i] = models.ThirdPartyBaker{
			Number:            i + 1,
			Address:           address,
			StakingBalance:    int64(acc.StakingBalance * 1e6),
			AvailableCapacity: 0,
			Efficiency:        efficiency,
		}
	}
	return thirdPartyBakers, nil
}

func (api *API) get(endpoint string, param url.Values, dst interface{}) error {
	u := fmt.Sprintf("%s%s?%s", apiURL, endpoint, param.Encode())
	resp, err := api.client.Get(u)
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
