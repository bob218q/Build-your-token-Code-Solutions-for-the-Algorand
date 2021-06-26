package main

import (

	"encoding/base64"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/algod" //"github.com/algorand/go-algorand-sdk/client/algod"
	"github.com/algorand/go-algorand-sdk/crypto"       //"https://github.com/algorand/go-algorand-sdk/tree/develop/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"    // "https://github.com/algorand/go-algorand-sdk/tree/develop/mnemonic"
	"github.com/algorand/go-algorand-sdk/transaction" // "https://github.com/algorand/go-algorand-sdk/tree/develop/transaction"

)

//POPcoin example

const algodAddress = "https://testnet-algorand.api.purestake.io/ps1"
const psToken = "..."

//Initialize the account and check funds before running the program
const mn = "..."
const ownerAddress = "..." //I only encoded but it could be derived from mnemonic

func main() {
//Create a client for Algorand
	
	var headers []*algod.Header
	headers = append(headers, &algod.Header{"X-API-Key", psToken})
	algodClient, err := algod.MakeClientWithHeaders(algodAddress, "", headers)
	if err != nil {
		fmt.Printf("failed to make algod client: %s\n", err)
		return
	}


	 //Recover private mnemonic key
	
	fromAddrPvtKey, err := mnemonic.ToPrivateKey(mn)
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}
     
	//The transaction parameters
	txParams, err := algodClient.SuggestedParams()
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}


	// Realize transaction
	coinTotalIssuancepop := uint64(1000000)
	coinDecimalsForDisplay := uint32(0) // 1 accounting unit in a transfer == 1 coin;
	accountsAreDefaultFrozen := false // if you have this coin, you can transact, the freeze manager doesn't need to unfreeze you first
	managerAddress := ownerAddress //The account you issue is also the account you manage
	assetReserveAddress := "" //there is no reserve of assets
	addressWithFreezingPrivileges := ownerAddress // can select accounts to be frozen to receive and send
	addressWithClawbackPrivileges := ownerAddress // this account is allowed to clawback coins from others
	assetUnitNamep := "POPCOIN"
	popassetName := "POPCOIN"
	assetMetadataHash := "" //There is no commitment hash is a simple example. . . .
	tx, err := transaction.MakeAssetCreateTxn(ownerAddress, txParams.Fee, txParams.LastRound, txParams.LastRound+10, nil, txParams.GenesisID, base64.StdEncoding.EncodeToString(txParams.GenesisHash),
		coinTotalIssuancepop, coinDecimalsForDisplay, accountsAreDefaultFrozen, managerAddress, assetReserveAddress, addressWithFreezingPrivileges, addressWithClawbackPrivileges,
		assetUnitNamep, popassetName, assetUrl, assetMetadataHash)

	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the Transaction
	_, bytes, err := crypto.SignTransaction(fromAddrPvtKey, tx)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}

	// Broadcast the transaction to the network
	txHeaders := append([]*algod.Header{}, &algod.Header{"Content-Type", "application/x-binary"})
	sendResponse, err := algodClient.SendRawTransaction(bytes, txHeaders...)
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}

	fmt.Printf("Transaction successful with ID: %s\n", sendResponse.TxID)
}
