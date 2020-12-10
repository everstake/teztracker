package models

import (
	"blockwatch.cc/tzindex/chain"
	script "blockwatch.cc/tzindex/micheline"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type AssetInfo struct {
	ID           uint64
	Name         string
	Balance      int64
	Ticker       string
	Source       string
	ContractType string
	AccountId    string
	Timestamp    time.Time
	Scale        int64
}

type HolderAddress string
type HolderBalance uint64

func (v *HolderAddress) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}
	data, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type")
	}

	bt, err := hex.DecodeString(data[2:])
	if err != nil {
		return err
	}

	adr := make([]byte, 22)

	switch len(bt) {
	//Bytes base58 address
	case 22:
		adr = bt
	default:
		//Remove prefix of michelson message
		if bt[0] == 5 {
			bt = bt[1:]
		}

		p := script.Prim{}

		err = p.UnmarshalBinary(bt)
		if err != nil {
			return err
		}

		//Contract calls
		if p.OpCode != script.D_PAIR {
			return nil
		}

		adr = p.Args[1].Bytes
	}

	address := chain.Address{}
	err = address.UnmarshalBinary(adr)
	if err != nil {
		return nil
	}

	*v = HolderAddress(address.String())
	return nil
}

func (v *HolderBalance) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	data, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid type")
	}

	if len(data) == 0 {
		return nil
	}

	var bal string
	switch data[0] {
	// bytes 0x
	case '0':
		bt, err := hex.DecodeString(data[2:])
		if err != nil {
			return err
		}

		//Remove prefix of michelson message
		if bt[0] == 5 {
			bt = bt[1:]
		}

		p := script.Prim{}

		err = p.UnmarshalBinary(bt)
		if err != nil {
			return err
		}

		if p.OpCode != script.D_PAIR {
			return nil
		}

		bal = p.Args[0].Int.String()
	case 'P':
		arr := strings.Split(data, " ")
		if len(arr) < 2 {
			return fmt.Errorf("Wrong pair")
		}

		//Int value on Michelson pair
		bal = arr[1]
	default:
		bal = data
	}

	balance, err := strconv.ParseInt(bal, 10, 64)
	if err != nil {
		return err
	}

	*v = HolderBalance(balance)
	return nil
}

type AssetHolder struct {
	Address HolderAddress
	Balance HolderBalance
}

type AssetOperation struct {
	BlockLevel         int64
	TokenId            uint64
	OperationId        int64
	OperationGroupHash string
	Sender             string
	Receiver           string
	Amount             int64
	Type               string
	Timestamp          time.Time
}

type AssetOperationReport struct {
	AssetOperation
	Fee          int64
	GasLimit     int64
	StorageLimit int64
}
