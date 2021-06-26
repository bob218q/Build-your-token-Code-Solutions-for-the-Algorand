# Build-your-token-Code-Solutions-for-the-Algorand
Hackathon: Code Solutions And Tutorials For The Algorand Developer Portal

Title: BUILD YOUR TOKEN

CUSTOMIZE,TRANSFORM AND SHARE IT YOUR TOKEN

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
 //developer.algorand.org/tutorials/creating-go-transaction-purestake-api/

const algodAddress = "https://testnet-algorand.api.purestake.io/ps1"
const psToken = "..."


    //Initialize the account and check funds before running the program
const mn = "..."
const ownerAddress = "..."   //I only encoded but it could be derived from mnemonic

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
assetUnitNamep := "PoPcoin"
popassetName := "POPcoin"
assetMetadataHash := "" //There is no commitment hash is a simple example. . . .

tx, err := transaction.MakeAssetCreateTxn(ownerAddress, txParams.Fee, txParams.LastRound, txParams.LastRound+10, nil, txParams.GenesisID, base64.StdEncoding.EncodeToString(txParams.GenesisHash),
    coinTotalIssuancepop, coinDecimalsForDisplay, accountsAreDefaultFrozen, managerAddress, assetReserveAddress, addressWithFreezingPrivileges,  addressWithClawbackPrivileges, assetUnitNamep, popassetName, assetMetadataHash)

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

In line 23 we give credits to purestake for their work.
In row 25-26 we import arguments for creating and communicating on the network.
![Schermata 2021-06-26 alle 11 43 43](https://user-images.githubusercontent.com/73669069/123509096-e1a1ea00-d673-11eb-9851-8382447f9980.png)
If it is not clear I attach documents (https://github.com/PureStake/api-examples).
For this demonstration we will connect to the PureStake node instead of algodClient.
It would be more appropriate to use environment variables to manage keys and next you will see a mnemonic address and the corresponding public address.
From line 33-58, at the beginning there are some settings like executing queries on recent network information, and also pay close attention to the step where you convert the mnemonic backup phrase to a private key.
The private key will be used to sign and authorize the transaction.
Once we understand all this, we can build our transaction and create the asset.
![Schermata 2021-06-26 alle 11 45 42](https://user-images.githubusercontent.com/73669069/123509134-1f9f0e00-d674-11eb-90e1-dacf416e2dd4.png)
The total coin issuance is 1 000 000, so I declare that there are one million POPcoin units.
Then I declare that the number of decimals to be used for accounting purposes is 0, i. e. I do not use as sats or wei or any other subdivision.
It’s worth remembering that this is just a fun demo currency, and POPcoin accounts are not frozen, they don’t need to do any whitelist before someone can make a transaction, they just need to participate (https://developer.algorand.org/docs/features/asa/#receiving-an-asset).
Then I name the asset and its units and instead of a link to a whitepaper I load the work on my github account. This transaction defines a new asset on Algorand’s blockchain.
![Schermata 2021-06-26 alle 12 01 29](https://user-images.githubusercontent.com/73669069/123509522-49593480-d676-11eb-8e5f-ab908861db0b.png)











