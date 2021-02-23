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
const routeConfig = "/explorer/config/head"
const routeAccount = "/explorer/account/%s"
const bakersMaxCount = 200

type (
	API struct {
		client *http.Client
	}
	Tip struct {
		Cycle int64   `json:"cycle"`
		Rolls float64 `json:"rolls"`
	}
	Config struct {
		BlockSecurityDeposit       float64 `json:"block_security_deposit"`
		BlocksPerCycle             float64 `json:"blocks_per_cycle"`
		EndorsersPerBlock          float64 `json:"endorsers_per_block"`
		EndorsementSecurityDeposit float64 `json:"endorsement_security_deposit"`
		PreservedCycles            float64 `json:"preserved_cycles"`
		TokensPerRoll              float64 `json:"tokens_per_roll"`
	}
	Account struct {
		Address          string  `json:"address"`
		DelegatedBalance float64 `json:"delegated_balance"`
		StakingBalance   float64 `json:"staking_balance"`
		BlocksBaked      int64   `json:"blocks_baked"`
		BlocksMissed     int64   `json:"blocks_missed"`
		SpendableBalance float64 `json:"spendable_balance"`
		FrozenDeposits   float64 `json:"frozen_deposits"`
		FrozenFees       float64 `json:"frozen_fees"`
		IsActiveDelegate bool    `json:"is_active_delegate"`
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
	var config Config
	err = api.get(routeConfig, nil, &config)
	if err != nil {
		return nil, fmt.Errorf("can`t get config: %s", err.Error())
	}
	blockDeposits := config.BlockSecurityDeposit + config.EndorsementSecurityDeposit*config.EndorsersPerBlock
	networkBond := blockDeposits * config.BlocksPerCycle * (config.PreservedCycles + 1)
	if networkBond == 0 {
		return nil, fmt.Errorf("wrong network bond")
	}
	networkStake := tip.Rolls * config.TokensPerRoll
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
	if len(incomes) > bakersMaxCount {
		incomes = incomes[:bakersMaxCount]
	}
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
		if !acc.IsActiveDelegate {
			continue
		}
		var efficiency float64
		if acc.BlocksBaked > 0 {
			efficiency = (float64(acc.BlocksBaked) / float64(acc.BlocksBaked+acc.BlocksMissed)) * 100
		}
		totalBalance := acc.SpendableBalance + acc.FrozenDeposits + acc.FrozenFees
		capacity := (totalBalance / networkBond) * networkStake
		thirdPartyBakers = append(thirdPartyBakers, models.ThirdPartyBaker{
			Number:            i + 1,
			Address:           address,
			StakingBalance:    int64(acc.StakingBalance * 1e6),
			AvailableCapacity: int64((capacity - acc.StakingBalance) * 1e6),
			Efficiency:        efficiency,
		})
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
