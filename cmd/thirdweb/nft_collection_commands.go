package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/qnfnypen/thirdweb-go-sdk/v2/thirdweb"
	"github.com/spf13/cobra"
)

var (
	nftContractAddress string
)

var nftCmd = &cobra.Command{
	Use:   "nft [command]",
	Short: "Interact with an nft contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var nftGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftCollection.GetAll(context.Background())
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftGetOwnedCmd = &cobra.Command{
	Use:   "getOwned",
	Short: "Get owned nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftCollection.GetOwned(context.Background(), "")
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftMintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		if tx, err := nftCollection.Mint(context.Background(), &thirdweb.NFTMetadataInput{
			Name:  "Special NFT",
			Image: imageFile,
		}); err != nil {
			panic(err)
		} else {
			// TODO return the minted token ID
			log.Printf("Minted nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var nftMintLinkCmd = &cobra.Command{
	Use:   "mintLink",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		if tx, err := nftCollection.Mint(context.Background(), &thirdweb.NFTMetadataInput{
			Name:  "NFT Test",
			Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
		}); err != nil {
			panic(err)
		} else {
			// TODO return the minted token ID
			log.Printf("Minted nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var nftSigmintCmd = &cobra.Command{
	Use:   "sigmint",
	Short: "Sign and mint an nft",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/1.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		payload, err := nftCollection.Signature.Generate(
			context.Background(),
			&thirdweb.Signature721PayloadInput{
				To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
				Price:                0,
				CurrencyAddress:      "0x0000000000000000000000000000000000000000",
				MintStartTime:        0,
				MintEndTime:          100000000000000,
				PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
				Metadata: &thirdweb.NFTMetadataInput{
					Name:        "Go #1",
					Description: "Minted with the Go SDK",
					Image:       imageFile,
				},
				RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
				RoyaltyBps:       0,
			},
		)
		if err != nil {
			panic(err)
		}

		valid, err := nftCollection.Signature.Verify(context.Background(), payload)
		if err != nil {
			panic(err)
		} else if !valid {
			panic("Invalid signature")
		}

		tx, err := nftCollection.Signature.MintAndAwait(context.Background(), payload)
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	nftCmd.PersistentFlags().StringVarP(&nftContractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
	nftCmd.AddCommand(nftGetOwnedCmd)
	nftCmd.AddCommand(nftMintCmd)
	nftCmd.AddCommand(nftMintLinkCmd)
	nftCmd.AddCommand(nftSigmintCmd)
}
