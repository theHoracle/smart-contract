package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/theHoracle/smart-contract.git/app/bank/proxy/contract/go/bank"
)

const (
	ownerStoreFile    = "zarf/ethereum/keystore/UTC--2025-01-27T20-59-34.898093000Z--b4c5b8ea361584dd0f8d28a094b9e7e1e9f336a4"
	account1StoreFile = "zarf/ethereum/keystore/UTC--2025-01-27T21-42-18.117661000Z--95ad8d78b3cd669c215736cf9e5d5070b123e1f5"

	account2StoreFile = "zarf/ethereum/keystore/UTC--2025-01-31T23-20-21.617023000Z--96790d0661a0907bae5570931d570921efa5c626"
	account3StoreFile = "zarf/ethereum/keystore/UTC--2025-01-31T23-20-31.165948000Z--0fbca64d7c2de24e7d3083d15c8241c00cf7fc54"
	account4StoreFile = "zarf/ethereum/keystore/UTC--2025-01-31T23-27-34.469442000Z--72402fa4294e1457d6bc81c0b79911ed66e633fb"
	passPhrase        = "1234" // All accounts use the same passphrase
	coinMarketCapKey  = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx := context.Background()

	balanceTarget := os.Getenv("BALANCE_TARGET")
	var ethAccount string

	// Validate the deposit target is valid.
	switch balanceTarget {
	case "owner":
		ethAccount = ownerStoreFile
	case "account1":
		ethAccount = account1StoreFile
	case "account2":
		ethAccount = account2StoreFile
	case "account3":
		ethAccount = account3StoreFile
	case "account4":
		ethAccount = account4StoreFile
	default:
		ethAccount = account1StoreFile
	}

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(ethAccount, passPhrase)
	if err != nil {
		return err
	}

	clt, err := ethereum.NewClient(backend, privateKey)
	if err != nil {
		return err
	}

	fmt.Println("\nInput Values")
	fmt.Println("----------------------------------------------------")
	fmt.Println("fromAddress:", clt.Address())

	// =========================================================================

	converter, err := currency.NewConverter(bank.BankMetaData.ABI, coinMarketCapKey)
	if err != nil {
		converter = currency.NewDefaultConverter(bank.BankMetaData.ABI)
	}
	oneETHToUSD, oneUSDToETH := converter.Values()

	fmt.Println("oneETHToUSD:", oneETHToUSD)
	fmt.Println("oneUSDToETH:", oneUSDToETH)

	// =========================================================================

	startingBalance, err := clt.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := clt.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	// =========================================================================

	contractIDBytes, err := os.ReadFile("zarf/ethereum/bank.cid")
	if err != nil {
		return fmt.Errorf("importing bank.cid file: %w", err)
	}

	contractID := string(contractIDBytes)
	if contractID == "" {
		return errors.New("need to export the bank.cid file")
	}
	fmt.Println("contractID:", contractID)

	proxyContract, err := bank.NewBank(common.HexToAddress(contractID), clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	balance, err := proxyContract.Balance(callOpts)
	if err != nil {
		return err
	}
	fmt.Printf("account balance: %v", balance)

	return nil
}
