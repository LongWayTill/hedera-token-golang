package main

import (
	"fmt"
	"os"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func main() {

	//LOADS THE .ENV FILE AND THROWS AN EROOR IF IT CANNOT LOAD THE VARIABLES
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("Unable to load enviroment variables from .env file. Error:\n%v\n", err))
	}

	//GRAB YOUR TESTNET ACCOUNT ID AND KEY FROMZ THE .ENV FILE
	myAccountId, err := hedera.AccountIDFromString(os.Getenv("MY_ACCOUNT_ID"))
	if err != nil {
		panic(err)
	}

	myPrivateKey, err := hedera.PrivateKeyFromString(os.Getenv("MY_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	//PRINT ACCOUNT ID AND KEY TO MAKE SURE THERE WASN'T AN ERROR READING FROM THE .ENV FILE
	fmt.Printf("The account ID is = %v\n", myAccountId)
	fmt.Printf("The private key is = %v\n", myPrivateKey)

	//CREATE TESTNET CLIENT
	client := hedera.ClientForTestnet()
	client.SetOperator(myAccountId, myPrivateKey)

}
