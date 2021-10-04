package nft

import (
	"context"
	"fmt"
	"time"

	"blockwatch.cc/tzgo/micheline"
	"github.com/anchorageoss/tezosprotocol/v2"
	"github.com/everstake/teztracker/models"
	"github.com/everstake/teztracker/repos/nft"
	"github.com/everstake/teztracker/repos/operation"
	"github.com/everstake/teztracker/services/ipfs"
	"github.com/sirupsen/logrus"
)

type IPFSClient interface {
	GetIPFSMetadata(ipfsID string) (desc ipfs.NuncNFTDescription, err error)
}

type UnitOfWork interface {
	GetOperation() operation.Repo
	GetNFT() nft.Repo
}

func ProcessNFTMintOperations(ctx context.Context, unit UnitOfWork, ipfsClient IPFSClient) (err error) {
	nftRepo := unit.GetNFT()

	//Get contracts list
	list, _, err := nftRepo.NTFContractsList("", 100, 0)
	if err != nil {
		return err
	}

	repo := unit.GetOperation()

	for i := range list {

		var nftTokens []models.NFTToken
		//Limit 20 to avoid restarts by ipfs timeout
		operations, _, err := repo.ContractOperationsList(list[i].AccountId, []string{"transaction"}, []string{"mint"}, list[i].LastHeight, 0, 20, models.SortAsc)
		if err != nil {
			return err
		}

		var contractTokens []models.NFTToken
		p := &micheline.Prim{}

		for j := range operations {
			err = p.UnmarshalJSON([]byte(operations[j].ParametersMicheline))
			if err != nil {
				return err
			}

			switch list[i].ContractType {
			//TODO move to consts
			case "nun":
				contractTokens, err = processNunMintOperation(list[i], p, ipfsClient)
				if err != nil {
					return err
				}
			case "kalamint":
				contractTokens, err = processKalamintMintOperation(list[i], p, ipfsClient)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("Unknown contract")
			}
			nftTokens = append(nftTokens, contractTokens...)

			list[i].LastHeight = operations[j].BlockLevel.Int64
		}

		//Safe new operations
		err = nftRepo.CreateBulk(nftTokens)
		if err != nil {
			return err
		}

		//TODO uncomment after tests
		//Count contract operations
		//_, operationsNum, err := repo.ContractOperationsList(list[i].AccountId, []string{"transaction"}, []string{}, 0, 0, 1, models.SortAsc)
		//if err != nil {
		//	return err
		//}
		//
		//list[i].OperationsNum = operationsNum

		err = nftRepo.UpdateNTFContractLastHeight(list[i])
		if err != nil {
			logrus.Errorf("failed to ProcessNFTOperations UpdateNTFContractLastHeight: %s", err.Error())
			return err
		}
	}

	return nil
}

func ProcessNFTOperations(ctx context.Context, unit UnitOfWork) (err error) {

	nftRepo := unit.GetNFT()

	//Get contracts list
	list, _, err := nftRepo.NTFContractsList("", 100, 0)
	if err != nil {
		return err
	}

	repo := unit.GetOperation()

	for i := range list {

		//transfer
		operations, _, err := repo.ContractOperationsList(list[i].SwapContract, []string{"transaction"}, []string{"buy", "list_token", "delist_token", "register_auction", "end_auction", "swap"}, list[i].LastUpdateHeight, 0, 20, models.SortAsc)
		if err != nil {
			return err
		}

		p := &micheline.Prim{}

		var isForSale bool
		var tokenID uint64

		for j := range operations {

			var lastPrice *int64

			err = p.UnmarshalJSON([]byte(operations[j].ParametersMicheline))
			if err != nil {
				return err
			}

			switch operations[j].ParametersEntrypoints {

			case "list_token":
				isForSale = true
				tokenID = p.Args[1].Int.Uint64()

				lastPrice = new(int64)
				*lastPrice = p.Args[0].Int.Int64()

			case "register_auction":
				isForSale = true
				tokenID = p.Args[1].Args[1].Int.Uint64()

				lastPrice = new(int64)
				*lastPrice = p.Args[1].Args[0].Int.Int64()

			case "buy", "delist_token":
				isForSale = false
				tokenID = p.Int.Uint64()

			case "end_auction":
				isForSale = false
				if len(p.Args) == 2 {
					tokenID = p.Args[1].Int.Uint64()
				} else {
					tokenID = p.Args[0].Int.Uint64()
				}

				lastPrice = new(int64)
				*lastPrice = operations[j].Amount

			case "swap": //Another Contract!
				isForSale = false
				tokenID = p.Args[1].Args[0].Int.Uint64()

				lastPrice = new(int64)
				*lastPrice = p.Args[1].Args[1].Int.Int64()

			}

			err = nftRepo.UpdateNFTToken(list[i].ID, tokenID, isForSale, lastPrice, operations[j].Timestamp)
			if err != nil {
				return err
			}

			list[i].LastUpdateHeight = operations[j].BlockLevel.Int64
		}

		err = nftRepo.UpdateNTFContractLastOPHeight(list[i])
		if err != nil {
			logrus.Errorf("failed to ProcessNFTOperations UpdateNTFContract: %s", err.Error())
			return err
		}
	}

	return nil
}

func processNunMintOperation(contract models.NFTContract, p *micheline.Prim, ipfs IPFSClient) (tokens []models.NFTToken, err error) {

	d := tezosprotocol.ContractID("")
	err = d.UnmarshalBinary(p.Args[0].Args[0].Bytes)
	if err != nil {
		return nil, err
	}

	token := models.NFTToken{
		ContractID: contract.ID,
		ID:         p.Args[1].Args[0].Int.Uint64(),
		LastPrice:  0,
		Amount:     p.Args[0].Args[1].Int.Int64(),
		IssuedBy:   string(d),

		IsForSale:    false,
		IpfsSource:   string(p.Args[1].Args[1].Args[0].Args[1].Bytes),
		CreatedAt:    time.Now(),
		LastActiveAt: time.Now(),
	}

	data, err := ipfs.GetIPFSMetadata(string(p.Args[1].Args[1].Args[0].Args[1].Bytes)[7:])
	if err != nil {
		return nil, err
	}

	token.Name = data.Name
	token.Description = data.Description
	token.Decimals = data.Decimals
	tokens = append(tokens, token)

	return tokens, nil
}

func processKalamintMintOperation(contract models.NFTContract, p *micheline.Prim, ipfs IPFSClient) (tokens []models.NFTToken, err error) {
	//Kalamint create n single NFTs
	// n = editions
	editions := p.Args[0].Args[1].Args[1].Args[0].Int.Int64()
	tokenID := p.Args[1].Args[1].Args[1].Args[0].Int.Int64()
	ipfsID := string(p.Args[1].Args[1].Args[1].Args[1].Bytes)

	tokens = make([]models.NFTToken, editions)

	data, err := ipfs.GetIPFSMetadata(ipfsID[7:])
	if err != nil {
		return nil, err
	}

	for i := int64(0); i < editions; i++ {

		var creator string
		if len(data.Creators) > 0 {
			creator = data.Creators[0]
		}
		tokens[i] = models.NFTToken{
			ContractID:   contract.ID,
			ID:           uint64(tokenID + i),
			Name:         data.Name,
			Category:     data.Category,
			Decimals:     data.Decimals,
			Description:  data.Description,
			Amount:       1,
			LastPrice:    p.Args[1].Args[1].Args[0].Args[0].Int.Int64(),
			IssuedBy:     creator,
			IsForSale:    p.Args[1].Args[0].Args[1].Args[1].OpCode == micheline.D_TRUE,
			IpfsSource:   ipfsID,
			CreatedAt:    time.Now(),
			LastActiveAt: time.Now(),
		}

	}

	return tokens, nil
}
