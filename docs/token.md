
## Token

You can access the Token interface from the SDK as follows:

```
import (
	"github.com/qnfnypen/thirdweb-go-sdk/v2/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetToken("{{contract_address}}")
```

```go
type Token struct {
    *ERC20Standard

    Helper  *contractHelper
    Encoder *ContractEncoder
    Events  *ContractEvents
}
```

### func \(\*Token\) [DelegateTo](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L159>)

```go
func (token *Token) DelegateTo(ctx context.Context, delegatreeAddress string) (*types.Transaction, error)
```

Delegate the connected wallets tokens to a specified wallet.

delegateeAddress: wallet address to delegate tokens to

returns: transaction receipt of the delegation

### func \(\*Token\) [GetDelegation](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L90>)

```go
func (token *Token) GetDelegation(ctx context.Context) (string, error)
```

Get the connected wallets delegatee address for this token.

returns: delegation address of the connected wallet

### func \(\*Token\) [GetDelegationOf](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L97>)

```go
func (token *Token) GetDelegationOf(ctx context.Context, address string) (string, error)
```

Get a specified wallets delegatee for this token.

returns: delegation address of the connected wallet

### func \(\*Token\) [GetVoteBalance](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L69>)

```go
func (token *Token) GetVoteBalance(ctx context.Context) (*CurrencyValue, error)
```

Get the connected wallets voting power in this token.

returns: vote balance of the connected wallet

### func \(\*Token\) [GetVoteBalanceOf](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L78>)

```go
func (token *Token) GetVoteBalanceOf(ctx context.Context, address string) (*CurrencyValue, error)
```

Get the voting power of the specified wallet in this token.

address: wallet address to check the vote balance of

returns: vote balance of the specified wallet

### func \(\*Token\) [Mint](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L111>)

```go
func (token *Token) Mint(ctx context.Context, amount float64) (*types.Transaction, error)
```

Mint tokens to the connected wallet.

amount: amount of tokens to mint

returns: transaction receipt of the mint

### func \(\*Token\) [MintBatchTo](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L150>)

```go
func (token *Token) MintBatchTo(ctx context.Context, args []*TokenAmount) (*types.Transaction, error)
```

Mint tokens to a list of wallets.

args: list of wallet addresses and amounts to mint

returns: transaction receipt of the mint

#### Example

```
args = []*thirdweb.TokenAmount{
	&thirdweb.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    1
	}
	&thirdweb.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    2
	}
}

tx, err := contract.MintBatchTo(context.Background(), args)
```

### func \(\*Token\) [MintTo](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/token.go#L126>)

```go
func (token *Token) MintTo(ctx context.Context, to string, amount float64) (*types.Transaction, error)
```

Mint tokens to a specified wallet.

to: wallet address to mint tokens to

amount: amount of tokens to mint

returns: transaction receipt of the mint

#### Example

```
tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", 1)
```

## type [TokenAmount](<https://github.com/qnfnypen/thirdweb-go-sdk/blob/main/thirdweb/types.go#L105-L108>)

```go
type TokenAmount struct {
    ToAddress string
    Amount    float64
}
```
