package main

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/log"
	"github.com/theHoracle/smart-contract.git/app/bank/single/contract/go/bank"
)

const (
	keyStoreFile     = "zarf/ethereum/keystore/UTC--2025-01-27T20-59-34.898093000Z--b4c5b8ea361584dd0f8d28a094b9e7e1e9f336a4"
	passPhrase       = "1234"
	coinMarketCapKey = "a8cd12fb-d056-423f-877b-659046af0aa5"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() (err error) {
	ctx := context.Background()

	stdOut := log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stdout, log.LevelInfo, true))
	discard := log.NewLogger(log.NewTerminalHandlerWithLevel(io.Discard, log.LevelInfo, true))

	backend, err := ethereum.CreateDialedBackend(ctx, ethereum.NetworkLocalhost)
	if err != nil {
		return err
	}
	defer backend.Close()

	privateKey, err := ethereum.PrivateKeyByKeyFile(keyStoreFile, passPhrase)
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

	const gasLimit = 1700000
	const gasPriceGwei = 39.576
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, currency.GWei2Wei(big.NewFloat(gasPriceGwei)), big.NewFloat(valueGwei))
	if err != nil {
		return err
	}
	address, tx, _, err := bank.DeployBank(tranOpts, clt.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())

	if err := os.WriteFile("zarf/ethereum/bank_single.cid", []byte(address.Hex()), 0o644); err != nil {
		return fmt.Errorf("exporting bank_single.cid file: %w", err)
	}

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.SetDefault(stdOut)

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))
	log.SetDefault(discard)

	// =========================================================================
	bankContract, err := bank.NewBank(address, clt.Backend)
	if err != nil {
		return fmt.Errorf("new proxy connection: %w", err)
	}

	callOpts, err := clt.NewCallOpts(ctx)
	if err != nil {
		return err
	}

	version, err := bankContract.Version(callOpts)
	if err != nil {
		return err
	}

	fmt.Println("\nValidate Version")
	fmt.Println("----------------------------------------------------")
	fmt.Println("version         :", version)

	return nil
}
