
## ERC1155 Standard

This interface is currently support by the Edition and Edition Drop contracts. You can access all of its functions through an Edition or Edition Drop contract instance.

```go
type ERC1155Standard struct {}
```

### func \(\*ERC1155Standard\) [Balance](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L93>)

```go
func (erc1155 *ERC1155Standard) Balance(ctx context.Context, tokenId int) (int, error)
```

Get the NFT balance of the connected wallet for a specific token ID.

tokenId: the token ID of a specific token to check the balance of

returns: the number of NFTs of the specified token ID owned by the connected wallet

### func \(\*ERC1155Standard\) [BalanceOf](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L108>)

```go
func (erc1155 *ERC1155Standard) BalanceOf(ctx context.Context, address string, tokenId int) (int, error)
```

Get the NFT balance of a specific wallet.

address: the address of the wallet to get the NFT balance of

returns: the number of NFTs of the specified token ID owned by the specified wallet

#### Example

```
address := "{{wallet_address}}"
tokenId := 0
balance, err := contract.BalanceOf(context.Background(), address, tokenId)
```

### func \(\*ERC1155Standard\) [Burn](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L157>)

```go
func (erc1155 *ERC1155Standard) Burn(ctx context.Context, tokenId int, amount int) (*types.Transaction, error)
```

Burn an amount of a specified NFT from the connected wallet.

tokenId: tokenID of the token to burn

amount: number of NFTs of the token ID to burn

returns: the transaction receipt of the burn

#### Example

```
tokenId := 0
amount := 1
tx, err := contract.Burn(context.Background(), tokenId, amount)
```

### func \(\*ERC1155Standard\) [Get](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L40>)

```go
func (erc1155 *ERC1155Standard) Get(ctx context.Context, tokenId int) (*EditionMetadata, error)
```

Get metadata for a token.

tokenId: token ID of the token to get the metadata for

returns: the metadata for the NFT and its supply

#### Example

```
nft, err := contract.Get(context.Background(), 0)
 supply := nft.Supply
	name := nft.Metadata.Name
```

### func \(\*ERC1155Standard\) [GetAll](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L53>)

```go
func (erc1155 *ERC1155Standard) GetAll(ctx context.Context) ([]*EditionMetadata, error)
```

Get the metadata of all the NFTs on this contract.

returns: the metadatas and supplies of all the NFTs on this contract

#### Example

```
nfts, err := contract.GetAll(context.Background())
supplyOne := nfts[0].Supply
nameOne := nfts[0].Metadata.Name
```

### func \(\*ERC1155Standard\) [GetOwned](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L75>)

```go
func (erc1155 *ERC1155Standard) GetOwned(ctx context.Context, address string) ([]*EditionMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address.

address: the address of the owner of the NFTs

returns: the metadatas and supplies of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(context.Background(), owner)
name := nfts[0].Metadata.Name
```

### func \(\*ERC1155Standard\) [GetTotalCount](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L60>)

```go
func (erc1155 *ERC1155Standard) GetTotalCount(ctx context.Context) (int, error)
```

Get the total number of NFTs on this contract.

returns: the total number of NFTs on this contract

### func \(\*ERC1155Standard\) [IsApproved](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L119>)

```go
func (erc1155 *ERC1155Standard) IsApproved(ctx context.Context, address string, operator string) (bool, error)
```

Check whether an operator address is approved for all operations of a specifc addresses assets.

address: the address whose assets are to be checked

operator: the address of the operator to check

returns: true if the operator is approved for all operations of the assets, otherwise false

### func \(\*ERC1155Standard\) [SetApprovalForAll](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L170>)

```go
func (erc1155 *ERC1155Standard) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error)
```

Set the approval for all operations of a specific address's assets.

address: the address whose assets are to be approved

operator: the address of the operator to set the approval for

approved: true if the operator is approved for all operations of the assets, otherwise false

returns: the transaction receipt of the approval

### func \(\*ERC1155Standard\) [TotalSupply](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L84>)

```go
func (erc1155 *ERC1155Standard) TotalSupply(ctx context.Context, tokenId int) (int, error)
```

Get the total number of NFTs of a specific token ID.

tokenId: the token ID to check the total supply of

returns: the supply of NFTs on the specified token ID

### func \(\*ERC1155Standard\) [Transfer](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/erc1155_standard.go#L140>)

```go
func (erc1155 *ERC1155Standard) Transfer(ctx context.Context, to string, tokenId int, amount int) (*types.Transaction, error)
```

Transfer a specific quantity of a token ID from the connected wallet to a specified address.

to: wallet address to transfer the tokens to

tokenId: the token ID of the NFT to transfer

amount: number of NFTs of the token ID to transfer

returns: the transaction of the NFT transfer

#### Example

```
to := "0x..."
tokenId := 0
amount := 1

tx, err := contract.Transfer(context.Background(), to, tokenId, amount)
```
