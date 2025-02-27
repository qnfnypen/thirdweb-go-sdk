
## ERC1155 Signature Minting

You can access this interface from the edition contract under the signature interface

```go
type ERC1155SignatureMinting struct {
    Helper *contractHelper
}
```

### func \(\*ERC1155SignatureMinting\) [Generate](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L207>)

```go
func (signature *ERC1155SignatureMinting) Generate(ctx context.Context, payloadToSign *Signature1155PayloadInput) (*SignedPayload1155, error)
```

#### Generate a payload to mint a new token ID

payloadToSign: the payload containing the data for the signature mint

returns: the payload signed by the minter's private key

#### Example

```
payload := &thirdweb.Signature721PayloadInput{
	To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6", // address to mint to
	Price:                0,                                            // cost of minting
	CurrencyAddress:      "0x0000000000000000000000000000000000000000", // currency to pay in order to mint
	MintStartTime:        0,                                            // time where minting is allowed to start (epoch seconds)
	MintEndTime:          100000000000000,                              // time when this signature expires (epoch seconds)
	PrimarySaleRecipient: "0x0000000000000000000000000000000000000000", // address to receive the primary sales of this mint
	Metadata: &thirdweb.NFTMetadataInput{																// metadata of the NFT to mint
 		Name:  "ERC721 Sigmint!",
	},
	RoyaltyRecipient: "0x0000000000000000000000000000000000000000",     // address to receive royalties of this mint
	RoyaltyBps:       0,                                                // royalty cut of this mint in basis points
	Quantity:         1,   																					    // number of tokens to mint
}

signedPayload, err := contract.Signature.Generate(context.Background(), payload)
```

### func \(\*ERC1155SignatureMinting\) [GenerateBatch](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L302>)

```go
func (signature *ERC1155SignatureMinting) GenerateBatch(ctx context.Context, payloadsToSign []*Signature1155PayloadInput) ([]*SignedPayload1155, error)
```

#### Generate a batch of payloads to mint multiple new token IDs

payloadToSign: the payloads containing the data for the signature mint

returns: the payloads signed by the minter's private key

#### Example

```
payload := []*thirdweb.Signature1155PayloadInput{
	&thirdweb.Signature1155PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
 			Name:  "ERC1155 Sigmint 1",
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
		Quantity:         1,
	},
	&thirdweb.Signature1155PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
 			Name:  "ERC1155 Sigmint 2",
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
		Quantity:         1,
	},
}

signedPayload, err := contract.Signature.GenerateBatch(context.Background(), payload)
```

### func \(\*ERC1155SignatureMinting\) [GenerateBatchFromTokenIds](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L369>)

```go
func (signature *ERC1155SignatureMinting) GenerateBatchFromTokenIds(ctx context.Context, payloadsToSign []*Signature1155PayloadInputWithTokenId) ([]*SignedPayload1155, error)
```

#### Generate a batch of payloads to mint multiple new token IDs

payloadToSign: the payloads containing the data for the signature mint

returns: the payloads signed by the minter's private key

#### Example

```
payload := []*thirdweb.Signature1155PayloadInputWithTokenId{
	&thirdweb.Signature1155PayloadInputWithTokenId{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
 			Name:  "ERC1155 Sigmint 1",
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
		Quantity:         1,
		TokenId:          0,
	},
	&thirdweb.Signature1155PayloadInputWithTokenId{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: nil
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
		Quantity:         1,
		TokenId:          1,
	},
}

signedPayload, err := contract.Signature.GenerateBatchFromTokenIds(context.Background(), payload)
```

### func \(\*ERC1155SignatureMinting\) [GenerateFromTokenId](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L253>)

```go
func (signature *ERC1155SignatureMinting) GenerateFromTokenId(ctx context.Context, payloadToSign *Signature1155PayloadInputWithTokenId) (*SignedPayload1155, error)
```

#### Generate a new payload to mint additionaly supply to an existing token ID

payloadToSign: the payload containing the data for the signature mint

returns: the payload signed by the minter's private key

#### Example

```
payload := &thirdweb.Signature1155PayloadInputWithTokenId{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata:             nil,                                          // we don't need to pass NFT metadata since we are minting an existing token
		RoyaltyRecipient:     "0x0000000000000000000000000000000000000000",
		RoyaltyBps:           0,
		Quantity:             1,
 	TokenId:              0,                                            // now we need to specify the token ID to mint supply to
	}

	signedPayload, err := contract.Signature.GenerateFromTokenId(context.Background(), payload)
```

### func \(\*ERC1155SignatureMinting\) [Mint](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L55>)

```go
func (signature *ERC1155SignatureMinting) Mint(ctx context.Context, signedPayload *SignedPayload1155) (*types.Transaction, error)
```

Mint a token with the data in given payload.

signedPayload: the payload signed by the minters private key being used to mint

returns: the transaction receipt of the mint

#### Example

```
// Learn more about how to craft a payload in the Generate() function
signedPayload, err := contract.Signature.Generate(payload)
tx, err := contract.Signature.Mint(signedPayload)
```

### func \(\*ERC1155SignatureMinting\) [MintBatch](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L104>)

```go
func (signature *ERC1155SignatureMinting) MintBatch(ctx context.Context, signedPayloads []*SignedPayload1155) (*types.Transaction, error)
```

Mint a batch of token with the data in given payload.

signedPayload: the list of payloads signed by the minters private key being used to mint

returns: the transaction receipt of the batch mint

#### Example

```
// Learn more about how to craft multiple payloads in the GenerateBatch() function
signedPayloads, err := contract.Signature.GenerateBatch(payloads)
tx, err := contract.Signature.MintBatch(signedPayloads)
```

### func \(\*ERC1155SignatureMinting\) [MintWithOpts](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L69>)

```go
func (signature *ERC1155SignatureMinting) MintWithOpts(ctx context.Context, signedPayload *SignedPayload1155, txOpts *bind.TransactOpts) (*types.Transaction, error)
```

### func \(\*ERC1155SignatureMinting\) [Verify](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_signature_minting.go#L167>)

```go
func (signature *ERC1155SignatureMinting) Verify(ctx context.Context, signedPayload *SignedPayload1155) (bool, error)
```

#### Verify that a signed payload is valid

signedPayload: the payload to verify

returns: true if the payload is valid, otherwise false.

#### Example

```
// Learn more about how to craft a payload in the Generate() function
signedPayload, err := contract.Signature.Generate(payload)
isValid, err := contract.Signature.Verify(context.Background(), signedPayload)
```
