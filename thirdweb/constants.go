package thirdweb

import "errors"

// SERVER URLS

const defaultIpfsGatewayUrl = "https://ipfs.io/ipfs/"
const twIpfsServerUrl = "https://upload.nftlabs.co"
const twStorageUploadUrl = "https://storage.thirdweb.com/ipfs/upload"
const pinataIpfsUrl = "https://api.pinata.cloud/pinning/pinFileToIPFS"

// CONSTANT VALUES

const zeroAddress = "0x0000000000000000000000000000000000000000"
const nativeTokenAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
const defaultMerkleRoot = "0x0000000000000000000000000000000000000000000000000000000000000000"

// NATIVE TOKEN BY CHAIN

type ChainID int
type ChainName string

const (
	MAINNET           ChainID = 1
	RINKEBY                   = 4
	GOERLI                    = 5
	SEPOLIA                   = 11155111
	POLYGON                   = 137
	FANTOM                    = 250
	FANTOM_TESTNET            = 4002
	AVALANCHE_TESTNET         = 43113
	AVALANCHE                 = 43114
	MUMBAI                    = 80001
	OPTIMISM                  = 10
	OPTIMISM_TESTNET          = 69
	ARBITRUM                  = 42161
	ARBITRUM_TESTNET          = 421611
)

func getNativeTokenByChainId(chainId ChainID) (*NativeToken, error) {
	switch chainId {
	case MAINNET:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case RINKEBY:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0xc778417E063141139Fce010982780140Aa0cD5Ab",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case GOERLI:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0x0bb7509324ce409f7bbc4b701f932eaca9736ab7",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case SEPOLIA:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0xfFf9976782d46CC05630D1f6eBAb18b2324d6B14",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case POLYGON:
		return &NativeToken{
			"Matic",
			"MATIC",
			18,
			&WrappedToken{
				"0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
				"Wrapped Matic",
				"WMATIC",
			},
		}, nil
	case MUMBAI:
		return &NativeToken{
			"Matic",
			"MATIC",
			18,
			&WrappedToken{
				"0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889",
				"Wrapped Matic",
				"WMATIC",
			},
		}, nil
	case AVALANCHE:
		return &NativeToken{
			"Avalanche",
			"AVAX",
			18,
			&WrappedToken{
				"0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7",
				"Wrapped AVAX",
				"WAVAX",
			},
		}, nil
	case AVALANCHE_TESTNET:
		return &NativeToken{
			"Avalanche",
			"AVAX",
			18,
			&WrappedToken{
				"0xd00ae08403B9bbb9124bB305C09058E32C39A48c",
				"Wrapped AVAX",
				"WAVAX",
			},
		}, nil
	case FANTOM:
		return &NativeToken{
			"Fantom",
			"FTM",
			18,
			&WrappedToken{
				"0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83",
				"Wrapped Fantom",
				"WFTM",
			},
		}, nil
	case FANTOM_TESTNET:
		return &NativeToken{
			"Fantom",
			"FTM",
			18,
			&WrappedToken{
				"0xf1277d1Ed8AD466beddF92ef448A132661956621",
				"Wrapped Fantom",
				"WFTM",
			},
		}, nil
	case ARBITRUM:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case ARBITRUM_TESTNET:
		return &NativeToken{
			"Arbitrum Rinkeby Ether",
			"ARETH",
			18,
			&WrappedToken{
				"0xEBbc3452Cc911591e4F18f3b36727Df45d6bd1f9",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case OPTIMISM:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0x4200000000000000000000000000000000000006",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case OPTIMISM_TESTNET:
		return &NativeToken{
			"Kovan Ether",
			"KOR",
			18,
			&WrappedToken{
				"0xbC6F6b680bc61e30dB47721c6D1c5cde19C1300d",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	default:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				zeroAddress,
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	}
}

// CONTRACT ADDRESSES BY CHAIN ID

const twRegistryAddress = "0x7c487845f98938Bb955B1D5AD069d9a30e4131fd"
const twFactoryAddress = "0x5DBC7B840baa9daBcBe9D2492E45D7244B54A2A0"

// const twFactoryAddress = "0x76F948E5F13B9A84A81E5681df8682BBf524805E"
const ozDefenderForwarderAddress = "0xc82BbE41f2cF04e3a8efA18F7032BDD7f6d98a81"

func getContractAddressByChainId(chainId ChainID, contractName string) (string, error) {
	var addresses map[string]string

	switch chainId {
	case MAINNET:
		addresses = map[string]string{
			"BiconomyForwarder": "0x84a0856b038eaAd1cC7E297cF34A7e72685A8693",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case RINKEBY:
		addresses = map[string]string{
			"BiconomyForwarder": "0xFD4973FeB2031D4409fB57afEE5dF2051b171104",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case GOERLI:
		addresses = map[string]string{
			"BiconomyForwarder": "0xE041608922d06a4F26C0d4c27d8bCD01daf1f792",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case SEPOLIA:
		addresses = map[string]string{
			"BiconomyForwarder":  "0xB79B22f66ED21f226720087E82fbE2b544869DEB",
			"TWFactory":          "0x7D58560196Dc564d63a07abb659D19b25c7594A5",
			"TWStatelessFactory": "0xd4c837e80fDA233F0a5d14D2b116C6A26cb7Bc80",
			"TWRegistry":         "0x8eeC22AEd54921f31b29c0d2cfdA2064974A6019",
		}
	case POLYGON:
		addresses = map[string]string{
			"BiconomyForwarder": "0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case MUMBAI:
		addresses = map[string]string{
			"BiconomyForwarder": "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case AVALANCHE:
		addresses = map[string]string{
			"BiconomyForwarder": "0x64CD353384109423a966dCd3Aa30D884C9b2E057",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case AVALANCHE_TESTNET:
		addresses = map[string]string{
			"BiconomyForwarder": "0x6271Ca63D30507f2Dcbf99B52787032506D75BBF",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case FANTOM:
		addresses = map[string]string{
			"BiconomyForwarder": "0x64CD353384109423a966dCd3Aa30D884C9b2E057",
			"TWFactory":         "0x97EA0Fcc552D5A8Fb5e9101316AAd0D62Ea0876B",
			"TWRegistry":        twRegistryAddress,
		}
	case FANTOM_TESTNET:
		addresses = map[string]string{
			"BiconomyForwarder": "0x69FB8Dca8067A5D38703b9e8b39cf2D51473E4b4",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case ARBITRUM:
		addresses = map[string]string{
			"BiconomyForwarder": "0xfe0fa3C06d03bDC7fb49c892BbB39113B534fB57",
			"TWFactory":         "0xd24b3de085CFd8c54b94feAD08a7962D343E6DE0",
			"TWRegistry":        twRegistryAddress,
		}
	case ARBITRUM_TESTNET:
		addresses = map[string]string{
			"BiconomyForwarder": zeroAddress,
			"TWFactory":         "0xb0435b47ad26115A39c59735b814f3769F07C2c1",
			"TWRegistry":        "0xcF4c511551aE4dab1F997866FC3900cd2aaeC40D",
		}
	case OPTIMISM:
		addresses = map[string]string{
			"BiconomyForwarder": "0xefba8a2a82ec1fb1273806174f5e28fbb917cf95",
			"TWFactory":         "0xd24b3de085CFd8c54b94feAD08a7962D343E6DE0",
			"TWRegistry":        twRegistryAddress,
		}
	case OPTIMISM_TESTNET:
		addresses = map[string]string{
			"BiconomyForwarder": zeroAddress,
			"TWFactory":         "0xd24b3de085CFd8c54b94feAD08a7962D343E6DE0",
			"TWRegistry":        twRegistryAddress,
		}
	default:
		return "", errors.New("Unsupported chain id")
	}

	if address, ok := addresses[contractName]; ok {
		return address, nil
	} else {
		return "", errors.New("Unsupported contract name")
	}
}

// CLAIM CONDITIONS

type ClaimEligibility string

const (
	NotEnoughSupply     ClaimEligibility = "There is not enough supply to claim."
	AddressNotAllowed                    = "This address is not on the allowlist."
	InsufficientBalance                  = "There isn't enough of the required currency in the wallet to pay for the claim."
	NoActiveClaimPhase                   = "There is no active claim phase at the moment. Please check back in later."
	NoClaimConditionSet                  = "There is no claim condition set."
	ExceedsMaxClaimable                  = "The quantity of tokens to claim is above the remaining limit for this wallet."
	NoWallet                             = "No wallet connected."
	Unknown                              = "No claim conditions found."
)
