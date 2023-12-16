// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ITokenERC1155MintRequest is an auto generated low-level Go binding around an user-defined struct.
type ITokenERC1155MintRequest struct {
	To                     common.Address
	RoyaltyRecipient       common.Address
	RoyaltyBps             *big.Int
	PrimarySaleRecipient   common.Address
	TokenId                *big.Int
	Uri                    string
	Quantity               *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
	ValidityStartTimestamp *big.Int
	ValidityEndTimestamp   *big.Int
	Uid                    [32]byte
}

// TokenERC1155MetaData contains all meta data concerning the TokenERC1155 contract.
var TokenERC1155MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"indexed\":false,\"internalType\":\"bool\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_fromTokenId\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_toTokenId\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultRoyalty\",\"inputs\":[{\"type\":\"address\",\"name\":\"newRoyaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"newRoyaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlatPlatformFeeUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"flatFee\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"indexed\":false,\"internalType\":\"uint8\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataFrozen\",\"inputs\":[],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnerUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"prevOwner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeInfoUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"platformFeeBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeTypeUpdated\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"feeType\",\"indexed\":false,\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrimarySaleRecipientUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"previousAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"newAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoyaltyForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensMinted\",\"inputs\":[{\"type\":\"address\",\"name\":\"mintedTo\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenIdMinted\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"indexed\":false,\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"quantityMinted\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensMintedWithSignature\",\"inputs\":[{\"type\":\"address\",\"name\":\"signer\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintedTo\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenIdMinted\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"mintRequest\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"quantity\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structITokenERC1155.MintRequest\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferBatch\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"values\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferSingle\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"URI\",\"inputs\":[{\"type\":\"string\",\"name\":\"value\",\"indexed\":false,\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"id\",\"indexed\":true,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfBatch\",\"inputs\":[{\"type\":\"address[]\",\"name\":\"accounts\",\"internalType\":\"address[]\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnBatch\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"values\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"contractURI\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes1\",\"name\":\"fields\",\"internalType\":\"bytes1\"},{\"type\":\"string\",\"name\":\"name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"version\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"chainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"verifyingContract\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"uint256[]\",\"name\":\"extensions\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeMetadata\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDefaultRoyaltyInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFlatPlatformFeeInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformFeeInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformFeeType\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_defaultAdmin\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_contractURI\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"_trustedForwarders\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"_primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"_royaltyBps\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint128\"},{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"type\":\"address\",\"name\":\"forwarder\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"_amount\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintWithSignature\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_req\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"quantity\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"internalType\":\"structITokenERC1155.MintRequest\"},{\"type\":\"bytes\",\"name\":\"_signature\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenIdToMint\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"platformFeeRecipient\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"primarySaleRecipient\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"royaltyInfo\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"salePrice\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyAmount\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeBatchTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"amounts\",\"internalType\":\"uint256[]\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"saleRecipientForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultRoyaltyInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_royaltyBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFlatPlatformFeeInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_flatFee\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"_newOwner\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeType\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"_feeType\",\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrimarySaleRecipient\",\"inputs\":[{\"type\":\"address\",\"name\":\"_saleRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_bps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uriFrozen\",\"inputs\":[],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_req\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"quantity\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"pricePerToken\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"internalType\":\"structITokenERC1155.MintRequest\"},{\"type\":\"bytes\",\"name\":\"_signature\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// TokenERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenERC1155MetaData.ABI instead.
var TokenERC1155ABI = TokenERC1155MetaData.ABI

// TokenERC1155 is an auto generated Go binding around an Ethereum contract.
type TokenERC1155 struct {
	TokenERC1155Caller     // Read-only binding to the contract
	TokenERC1155Transactor // Write-only binding to the contract
	TokenERC1155Filterer   // Log filterer for contract events
}

// TokenERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC1155Session struct {
	Contract     *TokenERC1155     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC1155CallerSession struct {
	Contract *TokenERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC1155TransactorSession struct {
	Contract     *TokenERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC1155Raw struct {
	Contract *TokenERC1155 // Generic contract binding to access the raw methods on
}

// TokenERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC1155CallerRaw struct {
	Contract *TokenERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC1155TransactorRaw struct {
	Contract *TokenERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC1155 creates a new instance of TokenERC1155, bound to a specific deployed contract.
func NewTokenERC1155(address common.Address, backend bind.ContractBackend) (*TokenERC1155, error) {
	contract, err := bindTokenERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155{TokenERC1155Caller: TokenERC1155Caller{contract: contract}, TokenERC1155Transactor: TokenERC1155Transactor{contract: contract}, TokenERC1155Filterer: TokenERC1155Filterer{contract: contract}}, nil
}

// NewTokenERC1155Caller creates a new read-only instance of TokenERC1155, bound to a specific deployed contract.
func NewTokenERC1155Caller(address common.Address, caller bind.ContractCaller) (*TokenERC1155Caller, error) {
	contract, err := bindTokenERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155Caller{contract: contract}, nil
}

// NewTokenERC1155Transactor creates a new write-only instance of TokenERC1155, bound to a specific deployed contract.
func NewTokenERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC1155Transactor, error) {
	contract, err := bindTokenERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155Transactor{contract: contract}, nil
}

// NewTokenERC1155Filterer creates a new log filterer instance of TokenERC1155, bound to a specific deployed contract.
func NewTokenERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC1155Filterer, error) {
	contract, err := bindTokenERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155Filterer{contract: contract}, nil
}

// bindTokenERC1155 binds a generic wrapper to an already deployed contract.
func bindTokenERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC1155 *TokenERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC1155.Contract.TokenERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC1155 *TokenERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC1155.Contract.TokenERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC1155 *TokenERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC1155.Contract.TokenERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC1155 *TokenERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC1155 *TokenERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC1155 *TokenERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC1155.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC1155 *TokenERC1155Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC1155 *TokenERC1155Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenERC1155.Contract.DEFAULTADMINROLE(&_TokenERC1155.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC1155 *TokenERC1155CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenERC1155.Contract.DEFAULTADMINROLE(&_TokenERC1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TokenERC1155.Contract.BalanceOf(&_TokenERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_TokenERC1155 *TokenERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _TokenERC1155.Contract.BalanceOf(&_TokenERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenERC1155 *TokenERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenERC1155 *TokenERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TokenERC1155.Contract.BalanceOfBatch(&_TokenERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_TokenERC1155 *TokenERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _TokenERC1155.Contract.BalanceOfBatch(&_TokenERC1155.CallOpts, accounts, ids)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC1155 *TokenERC1155Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC1155 *TokenERC1155Session) ContractType() ([32]byte, error) {
	return _TokenERC1155.Contract.ContractType(&_TokenERC1155.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC1155 *TokenERC1155CallerSession) ContractType() ([32]byte, error) {
	return _TokenERC1155.Contract.ContractType(&_TokenERC1155.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC1155 *TokenERC1155Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC1155 *TokenERC1155Session) ContractURI() (string, error) {
	return _TokenERC1155.Contract.ContractURI(&_TokenERC1155.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC1155 *TokenERC1155CallerSession) ContractURI() (string, error) {
	return _TokenERC1155.Contract.ContractURI(&_TokenERC1155.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC1155 *TokenERC1155Caller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC1155 *TokenERC1155Session) ContractVersion() (uint8, error) {
	return _TokenERC1155.Contract.ContractVersion(&_TokenERC1155.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC1155 *TokenERC1155CallerSession) ContractVersion() (uint8, error) {
	return _TokenERC1155.Contract.ContractVersion(&_TokenERC1155.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_TokenERC1155 *TokenERC1155Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_TokenERC1155 *TokenERC1155Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _TokenERC1155.Contract.Eip712Domain(&_TokenERC1155.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_TokenERC1155 *TokenERC1155CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _TokenERC1155.Contract.Eip712Domain(&_TokenERC1155.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Caller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Session) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetDefaultRoyaltyInfo(&_TokenERC1155.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155CallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetDefaultRoyaltyInfo(&_TokenERC1155.CallOpts)
}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_TokenERC1155 *TokenERC1155Caller) GetFlatPlatformFeeInfo(opts *bind.CallOpts) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getFlatPlatformFeeInfo")

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_TokenERC1155 *TokenERC1155Session) GetFlatPlatformFeeInfo() (common.Address, *big.Int, error) {
	return _TokenERC1155.Contract.GetFlatPlatformFeeInfo(&_TokenERC1155.CallOpts)
}

// GetFlatPlatformFeeInfo is a free data retrieval call binding the contract method 0xe57553da.
//
// Solidity: function getFlatPlatformFeeInfo() view returns(address, uint256)
func (_TokenERC1155 *TokenERC1155CallerSession) GetFlatPlatformFeeInfo() (common.Address, *big.Int, error) {
	return _TokenERC1155.Contract.GetFlatPlatformFeeInfo(&_TokenERC1155.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Caller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getPlatformFeeInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Session) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetPlatformFeeInfo(&_TokenERC1155.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155CallerSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetPlatformFeeInfo(&_TokenERC1155.CallOpts)
}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_TokenERC1155 *TokenERC1155Caller) GetPlatformFeeType(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getPlatformFeeType")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_TokenERC1155 *TokenERC1155Session) GetPlatformFeeType() (uint8, error) {
	return _TokenERC1155.Contract.GetPlatformFeeType(&_TokenERC1155.CallOpts)
}

// GetPlatformFeeType is a free data retrieval call binding the contract method 0xf28083c3.
//
// Solidity: function getPlatformFeeType() view returns(uint8)
func (_TokenERC1155 *TokenERC1155CallerSession) GetPlatformFeeType() (uint8, error) {
	return _TokenERC1155.Contract.GetPlatformFeeType(&_TokenERC1155.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC1155 *TokenERC1155Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC1155 *TokenERC1155Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenERC1155.Contract.GetRoleAdmin(&_TokenERC1155.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC1155 *TokenERC1155CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenERC1155.Contract.GetRoleAdmin(&_TokenERC1155.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC1155 *TokenERC1155Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC1155 *TokenERC1155Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenERC1155.Contract.GetRoleMember(&_TokenERC1155.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC1155 *TokenERC1155CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenERC1155.Contract.GetRoleMember(&_TokenERC1155.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenERC1155.Contract.GetRoleMemberCount(&_TokenERC1155.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC1155 *TokenERC1155CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenERC1155.Contract.GetRoleMemberCount(&_TokenERC1155.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Caller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155Session) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetRoyaltyInfoForToken(&_TokenERC1155.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_TokenERC1155 *TokenERC1155CallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _TokenERC1155.Contract.GetRoyaltyInfoForToken(&_TokenERC1155.CallOpts, _tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC1155 *TokenERC1155Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC1155 *TokenERC1155Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenERC1155.Contract.HasRole(&_TokenERC1155.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC1155 *TokenERC1155CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenERC1155.Contract.HasRole(&_TokenERC1155.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenERC1155 *TokenERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenERC1155 *TokenERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TokenERC1155.Contract.IsApprovedForAll(&_TokenERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_TokenERC1155 *TokenERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _TokenERC1155.Contract.IsApprovedForAll(&_TokenERC1155.CallOpts, account, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC1155 *TokenERC1155Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC1155 *TokenERC1155Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenERC1155.Contract.IsTrustedForwarder(&_TokenERC1155.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC1155 *TokenERC1155CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenERC1155.Contract.IsTrustedForwarder(&_TokenERC1155.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC1155 *TokenERC1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC1155 *TokenERC1155Session) Name() (string, error) {
	return _TokenERC1155.Contract.Name(&_TokenERC1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC1155 *TokenERC1155CallerSession) Name() (string, error) {
	return _TokenERC1155.Contract.Name(&_TokenERC1155.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC1155 *TokenERC1155Caller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC1155 *TokenERC1155Session) NextTokenIdToMint() (*big.Int, error) {
	return _TokenERC1155.Contract.NextTokenIdToMint(&_TokenERC1155.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC1155 *TokenERC1155CallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _TokenERC1155.Contract.NextTokenIdToMint(&_TokenERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC1155 *TokenERC1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC1155 *TokenERC1155Session) Owner() (common.Address, error) {
	return _TokenERC1155.Contract.Owner(&_TokenERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC1155 *TokenERC1155CallerSession) Owner() (common.Address, error) {
	return _TokenERC1155.Contract.Owner(&_TokenERC1155.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155Caller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155Session) PlatformFeeRecipient() (common.Address, error) {
	return _TokenERC1155.Contract.PlatformFeeRecipient(&_TokenERC1155.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155CallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _TokenERC1155.Contract.PlatformFeeRecipient(&_TokenERC1155.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155Caller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155Session) PrimarySaleRecipient() (common.Address, error) {
	return _TokenERC1155.Contract.PrimarySaleRecipient(&_TokenERC1155.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC1155 *TokenERC1155CallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _TokenERC1155.Contract.PrimarySaleRecipient(&_TokenERC1155.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_TokenERC1155 *TokenERC1155Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_TokenERC1155 *TokenERC1155Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _TokenERC1155.Contract.RoyaltyInfo(&_TokenERC1155.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_TokenERC1155 *TokenERC1155CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _TokenERC1155.Contract.RoyaltyInfo(&_TokenERC1155.CallOpts, tokenId, salePrice)
}

// SaleRecipientForToken is a free data retrieval call binding the contract method 0xea500d69.
//
// Solidity: function saleRecipientForToken(uint256 ) view returns(address)
func (_TokenERC1155 *TokenERC1155Caller) SaleRecipientForToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "saleRecipientForToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleRecipientForToken is a free data retrieval call binding the contract method 0xea500d69.
//
// Solidity: function saleRecipientForToken(uint256 ) view returns(address)
func (_TokenERC1155 *TokenERC1155Session) SaleRecipientForToken(arg0 *big.Int) (common.Address, error) {
	return _TokenERC1155.Contract.SaleRecipientForToken(&_TokenERC1155.CallOpts, arg0)
}

// SaleRecipientForToken is a free data retrieval call binding the contract method 0xea500d69.
//
// Solidity: function saleRecipientForToken(uint256 ) view returns(address)
func (_TokenERC1155 *TokenERC1155CallerSession) SaleRecipientForToken(arg0 *big.Int) (common.Address, error) {
	return _TokenERC1155.Contract.SaleRecipientForToken(&_TokenERC1155.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC1155 *TokenERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC1155 *TokenERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenERC1155.Contract.SupportsInterface(&_TokenERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC1155 *TokenERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenERC1155.Contract.SupportsInterface(&_TokenERC1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC1155 *TokenERC1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC1155 *TokenERC1155Session) Symbol() (string, error) {
	return _TokenERC1155.Contract.Symbol(&_TokenERC1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC1155 *TokenERC1155CallerSession) Symbol() (string, error) {
	return _TokenERC1155.Contract.Symbol(&_TokenERC1155.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Caller) TotalSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "totalSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_TokenERC1155 *TokenERC1155Session) TotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _TokenERC1155.Contract.TotalSupply(&_TokenERC1155.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_TokenERC1155 *TokenERC1155CallerSession) TotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _TokenERC1155.Contract.TotalSupply(&_TokenERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_TokenERC1155 *TokenERC1155Caller) Uri(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "uri", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_TokenERC1155 *TokenERC1155Session) Uri(_tokenId *big.Int) (string, error) {
	return _TokenERC1155.Contract.Uri(&_TokenERC1155.CallOpts, _tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string)
func (_TokenERC1155 *TokenERC1155CallerSession) Uri(_tokenId *big.Int) (string, error) {
	return _TokenERC1155.Contract.Uri(&_TokenERC1155.CallOpts, _tokenId)
}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC1155 *TokenERC1155Caller) UriFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "uriFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC1155 *TokenERC1155Session) UriFrozen() (bool, error) {
	return _TokenERC1155.Contract.UriFrozen(&_TokenERC1155.CallOpts)
}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC1155 *TokenERC1155CallerSession) UriFrozen() (bool, error) {
	return _TokenERC1155.Contract.UriFrozen(&_TokenERC1155.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC1155 *TokenERC1155Caller) Verify(opts *bind.CallOpts, _req ITokenERC1155MintRequest, _signature []byte) (bool, common.Address, error) {
	var out []interface{}
	err := _TokenERC1155.contract.Call(opts, &out, "verify", _req, _signature)

	if err != nil {
		return *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC1155 *TokenERC1155Session) Verify(_req ITokenERC1155MintRequest, _signature []byte) (bool, common.Address, error) {
	return _TokenERC1155.Contract.Verify(&_TokenERC1155.CallOpts, _req, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC1155 *TokenERC1155CallerSession) Verify(_req ITokenERC1155MintRequest, _signature []byte) (bool, common.Address, error) {
	return _TokenERC1155.Contract.Verify(&_TokenERC1155.CallOpts, _req, _signature)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_TokenERC1155 *TokenERC1155Transactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_TokenERC1155 *TokenERC1155Session) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Burn(&_TokenERC1155.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Burn(&_TokenERC1155.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_TokenERC1155 *TokenERC1155Transactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_TokenERC1155 *TokenERC1155Session) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.BurnBatch(&_TokenERC1155.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.BurnBatch(&_TokenERC1155.TransactOpts, account, ids, values)
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC1155 *TokenERC1155Transactor) FreezeMetadata(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "freezeMetadata")
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC1155 *TokenERC1155Session) FreezeMetadata() (*types.Transaction, error) {
	return _TokenERC1155.Contract.FreezeMetadata(&_TokenERC1155.TransactOpts)
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) FreezeMetadata() (*types.Transaction, error) {
	return _TokenERC1155.Contract.FreezeMetadata(&_TokenERC1155.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.GrantRole(&_TokenERC1155.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.GrantRole(&_TokenERC1155.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _primarySaleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC1155 *TokenERC1155Transactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _primarySaleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _primarySaleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _primarySaleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC1155 *TokenERC1155Session) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _primarySaleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Initialize(&_TokenERC1155.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _primarySaleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _primarySaleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _primarySaleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Initialize(&_TokenERC1155.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _primarySaleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// MintTo is a paid mutator transaction binding the contract method 0xb03f4528.
//
// Solidity: function mintTo(address _to, uint256 _tokenId, string _uri, uint256 _amount) returns()
func (_TokenERC1155 *TokenERC1155Transactor) MintTo(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _uri string, _amount *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "mintTo", _to, _tokenId, _uri, _amount)
}

// MintTo is a paid mutator transaction binding the contract method 0xb03f4528.
//
// Solidity: function mintTo(address _to, uint256 _tokenId, string _uri, uint256 _amount) returns()
func (_TokenERC1155 *TokenERC1155Session) MintTo(_to common.Address, _tokenId *big.Int, _uri string, _amount *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.MintTo(&_TokenERC1155.TransactOpts, _to, _tokenId, _uri, _amount)
}

// MintTo is a paid mutator transaction binding the contract method 0xb03f4528.
//
// Solidity: function mintTo(address _to, uint256 _tokenId, string _uri, uint256 _amount) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) MintTo(_to common.Address, _tokenId *big.Int, _uri string, _amount *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.MintTo(&_TokenERC1155.TransactOpts, _to, _tokenId, _uri, _amount)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns()
func (_TokenERC1155 *TokenERC1155Transactor) MintWithSignature(opts *bind.TransactOpts, _req ITokenERC1155MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "mintWithSignature", _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns()
func (_TokenERC1155 *TokenERC1155Session) MintWithSignature(_req ITokenERC1155MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.MintWithSignature(&_TokenERC1155.TransactOpts, _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) MintWithSignature(_req ITokenERC1155MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.MintWithSignature(&_TokenERC1155.TransactOpts, _req, _signature)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC1155 *TokenERC1155Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC1155 *TokenERC1155Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Multicall(&_TokenERC1155.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC1155 *TokenERC1155TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.Multicall(&_TokenERC1155.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.RenounceRole(&_TokenERC1155.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.RenounceRole(&_TokenERC1155.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.RevokeRole(&_TokenERC1155.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.RevokeRole(&_TokenERC1155.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenERC1155 *TokenERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SafeBatchTransferFrom(&_TokenERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SafeBatchTransferFrom(&_TokenERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenERC1155 *TokenERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SafeTransferFrom(&_TokenERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SafeTransferFrom(&_TokenERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC1155 *TokenERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetApprovalForAll(&_TokenERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetApprovalForAll(&_TokenERC1155.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC1155 *TokenERC1155Session) SetContractURI(_uri string) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetContractURI(&_TokenERC1155.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetContractURI(&_TokenERC1155.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC1155 *TokenERC1155Session) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetDefaultRoyaltyInfo(&_TokenERC1155.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetDefaultRoyaltyInfo(&_TokenERC1155.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetFlatPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setFlatPlatformFeeInfo", _platformFeeRecipient, _flatFee)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_TokenERC1155 *TokenERC1155Session) SetFlatPlatformFeeInfo(_platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetFlatPlatformFeeInfo(&_TokenERC1155.TransactOpts, _platformFeeRecipient, _flatFee)
}

// SetFlatPlatformFeeInfo is a paid mutator transaction binding the contract method 0x7e54523c.
//
// Solidity: function setFlatPlatformFeeInfo(address _platformFeeRecipient, uint256 _flatFee) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetFlatPlatformFeeInfo(_platformFeeRecipient common.Address, _flatFee *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetFlatPlatformFeeInfo(&_TokenERC1155.TransactOpts, _platformFeeRecipient, _flatFee)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC1155 *TokenERC1155Session) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetOwner(&_TokenERC1155.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetOwner(&_TokenERC1155.TransactOpts, _newOwner)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC1155 *TokenERC1155Session) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPlatformFeeInfo(&_TokenERC1155.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPlatformFeeInfo(&_TokenERC1155.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetPlatformFeeType(opts *bind.TransactOpts, _feeType uint8) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setPlatformFeeType", _feeType)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_TokenERC1155 *TokenERC1155Session) SetPlatformFeeType(_feeType uint8) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPlatformFeeType(&_TokenERC1155.TransactOpts, _feeType)
}

// SetPlatformFeeType is a paid mutator transaction binding the contract method 0xb6f10c79.
//
// Solidity: function setPlatformFeeType(uint8 _feeType) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetPlatformFeeType(_feeType uint8) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPlatformFeeType(&_TokenERC1155.TransactOpts, _feeType)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC1155 *TokenERC1155Session) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPrimarySaleRecipient(&_TokenERC1155.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetPrimarySaleRecipient(&_TokenERC1155.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC1155 *TokenERC1155Session) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetRoyaltyInfoForToken(&_TokenERC1155.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetRoyaltyInfoForToken(&_TokenERC1155.TransactOpts, _tokenId, _recipient, _bps)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC1155 *TokenERC1155Transactor) SetTokenURI(opts *bind.TransactOpts, _tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC1155.contract.Transact(opts, "setTokenURI", _tokenId, _uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC1155 *TokenERC1155Session) SetTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetTokenURI(&_TokenERC1155.TransactOpts, _tokenId, _uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC1155 *TokenERC1155TransactorSession) SetTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC1155.Contract.SetTokenURI(&_TokenERC1155.TransactOpts, _tokenId, _uri)
}

// TokenERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TokenERC1155 contract.
type TokenERC1155ApprovalForAllIterator struct {
	Event *TokenERC1155ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155ApprovalForAll represents a ApprovalForAll event raised by the TokenERC1155 contract.
type TokenERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenERC1155 *TokenERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*TokenERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155ApprovalForAllIterator{contract: _TokenERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenERC1155 *TokenERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155ApprovalForAll)
				if err := _TokenERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_TokenERC1155 *TokenERC1155Filterer) ParseApprovalForAll(log types.Log) (*TokenERC1155ApprovalForAll, error) {
	event := new(TokenERC1155ApprovalForAll)
	if err := _TokenERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the TokenERC1155 contract.
type TokenERC1155BatchMetadataUpdateIterator struct {
	Event *TokenERC1155BatchMetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155BatchMetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155BatchMetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the TokenERC1155 contract.
type TokenERC1155BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TokenERC1155 *TokenERC1155Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TokenERC1155BatchMetadataUpdateIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155BatchMetadataUpdateIterator{contract: _TokenERC1155.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TokenERC1155 *TokenERC1155Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TokenERC1155BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155BatchMetadataUpdate)
				if err := _TokenERC1155.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TokenERC1155 *TokenERC1155Filterer) ParseBatchMetadataUpdate(log types.Log) (*TokenERC1155BatchMetadataUpdate, error) {
	event := new(TokenERC1155BatchMetadataUpdate)
	if err := _TokenERC1155.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155DefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the TokenERC1155 contract.
type TokenERC1155DefaultRoyaltyIterator struct {
	Event *TokenERC1155DefaultRoyalty // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155DefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155DefaultRoyalty)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155DefaultRoyalty)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155DefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155DefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155DefaultRoyalty represents a DefaultRoyalty event raised by the TokenERC1155 contract.
type TokenERC1155DefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*TokenERC1155DefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155DefaultRoyaltyIterator{contract: _TokenERC1155.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *TokenERC1155DefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155DefaultRoyalty)
				if err := _TokenERC1155.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultRoyalty is a log parse operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) ParseDefaultRoyalty(log types.Log) (*TokenERC1155DefaultRoyalty, error) {
	event := new(TokenERC1155DefaultRoyalty)
	if err := _TokenERC1155.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the TokenERC1155 contract.
type TokenERC1155EIP712DomainChangedIterator struct {
	Event *TokenERC1155EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155EIP712DomainChanged represents a EIP712DomainChanged event raised by the TokenERC1155 contract.
type TokenERC1155EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_TokenERC1155 *TokenERC1155Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*TokenERC1155EIP712DomainChangedIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155EIP712DomainChangedIterator{contract: _TokenERC1155.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_TokenERC1155 *TokenERC1155Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *TokenERC1155EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155EIP712DomainChanged)
				if err := _TokenERC1155.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_TokenERC1155 *TokenERC1155Filterer) ParseEIP712DomainChanged(log types.Log) (*TokenERC1155EIP712DomainChanged, error) {
	event := new(TokenERC1155EIP712DomainChanged)
	if err := _TokenERC1155.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155FlatPlatformFeeUpdatedIterator is returned from FilterFlatPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for FlatPlatformFeeUpdated events raised by the TokenERC1155 contract.
type TokenERC1155FlatPlatformFeeUpdatedIterator struct {
	Event *TokenERC1155FlatPlatformFeeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155FlatPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155FlatPlatformFeeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155FlatPlatformFeeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155FlatPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155FlatPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155FlatPlatformFeeUpdated represents a FlatPlatformFeeUpdated event raised by the TokenERC1155 contract.
type TokenERC1155FlatPlatformFeeUpdated struct {
	PlatformFeeRecipient common.Address
	FlatFee              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFlatPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_TokenERC1155 *TokenERC1155Filterer) FilterFlatPlatformFeeUpdated(opts *bind.FilterOpts) (*TokenERC1155FlatPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155FlatPlatformFeeUpdatedIterator{contract: _TokenERC1155.contract, event: "FlatPlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFlatPlatformFeeUpdated is a free log subscription operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_TokenERC1155 *TokenERC1155Filterer) WatchFlatPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC1155FlatPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155FlatPlatformFeeUpdated)
				if err := _TokenERC1155.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlatPlatformFeeUpdated is a log parse operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_TokenERC1155 *TokenERC1155Filterer) ParseFlatPlatformFeeUpdated(log types.Log) (*TokenERC1155FlatPlatformFeeUpdated, error) {
	event := new(TokenERC1155FlatPlatformFeeUpdated)
	if err := _TokenERC1155.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenERC1155 contract.
type TokenERC1155InitializedIterator struct {
	Event *TokenERC1155Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155Initialized represents a Initialized event raised by the TokenERC1155 contract.
type TokenERC1155Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenERC1155 *TokenERC1155Filterer) FilterInitialized(opts *bind.FilterOpts) (*TokenERC1155InitializedIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155InitializedIterator{contract: _TokenERC1155.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenERC1155 *TokenERC1155Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenERC1155Initialized) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155Initialized)
				if err := _TokenERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenERC1155 *TokenERC1155Filterer) ParseInitialized(log types.Log) (*TokenERC1155Initialized, error) {
	event := new(TokenERC1155Initialized)
	if err := _TokenERC1155.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155MetadataFrozenIterator is returned from FilterMetadataFrozen and is used to iterate over the raw logs and unpacked data for MetadataFrozen events raised by the TokenERC1155 contract.
type TokenERC1155MetadataFrozenIterator struct {
	Event *TokenERC1155MetadataFrozen // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155MetadataFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155MetadataFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155MetadataFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155MetadataFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155MetadataFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155MetadataFrozen represents a MetadataFrozen event raised by the TokenERC1155 contract.
type TokenERC1155MetadataFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMetadataFrozen is a free log retrieval operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_TokenERC1155 *TokenERC1155Filterer) FilterMetadataFrozen(opts *bind.FilterOpts) (*TokenERC1155MetadataFrozenIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155MetadataFrozenIterator{contract: _TokenERC1155.contract, event: "MetadataFrozen", logs: logs, sub: sub}, nil
}

// WatchMetadataFrozen is a free log subscription operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_TokenERC1155 *TokenERC1155Filterer) WatchMetadataFrozen(opts *bind.WatchOpts, sink chan<- *TokenERC1155MetadataFrozen) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155MetadataFrozen)
				if err := _TokenERC1155.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataFrozen is a log parse operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_TokenERC1155 *TokenERC1155Filterer) ParseMetadataFrozen(log types.Log) (*TokenERC1155MetadataFrozen, error) {
	event := new(TokenERC1155MetadataFrozen)
	if err := _TokenERC1155.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the TokenERC1155 contract.
type TokenERC1155MetadataUpdateIterator struct {
	Event *TokenERC1155MetadataUpdate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155MetadataUpdate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155MetadataUpdate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155MetadataUpdate represents a MetadataUpdate event raised by the TokenERC1155 contract.
type TokenERC1155MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TokenERC1155 *TokenERC1155Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*TokenERC1155MetadataUpdateIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155MetadataUpdateIterator{contract: _TokenERC1155.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TokenERC1155 *TokenERC1155Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TokenERC1155MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155MetadataUpdate)
				if err := _TokenERC1155.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TokenERC1155 *TokenERC1155Filterer) ParseMetadataUpdate(log types.Log) (*TokenERC1155MetadataUpdate, error) {
	event := new(TokenERC1155MetadataUpdate)
	if err := _TokenERC1155.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155OwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the TokenERC1155 contract.
type TokenERC1155OwnerUpdatedIterator struct {
	Event *TokenERC1155OwnerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155OwnerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155OwnerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155OwnerUpdated represents a OwnerUpdated event raised by the TokenERC1155 contract.
type TokenERC1155OwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_TokenERC1155 *TokenERC1155Filterer) FilterOwnerUpdated(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*TokenERC1155OwnerUpdatedIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155OwnerUpdatedIterator{contract: _TokenERC1155.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_TokenERC1155 *TokenERC1155Filterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC1155OwnerUpdated, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155OwnerUpdated)
				if err := _TokenERC1155.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnerUpdated is a log parse operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_TokenERC1155 *TokenERC1155Filterer) ParseOwnerUpdated(log types.Log) (*TokenERC1155OwnerUpdated, error) {
	event := new(TokenERC1155OwnerUpdated)
	if err := _TokenERC1155.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155PlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the TokenERC1155 contract.
type TokenERC1155PlatformFeeInfoUpdatedIterator struct {
	Event *TokenERC1155PlatformFeeInfoUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155PlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155PlatformFeeInfoUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155PlatformFeeInfoUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155PlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155PlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155PlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the TokenERC1155 contract.
type TokenERC1155PlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	PlatformFeeBps       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_TokenERC1155 *TokenERC1155Filterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts, platformFeeRecipient []common.Address) (*TokenERC1155PlatformFeeInfoUpdatedIterator, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155PlatformFeeInfoUpdatedIterator{contract: _TokenERC1155.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_TokenERC1155 *TokenERC1155Filterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC1155PlatformFeeInfoUpdated, platformFeeRecipient []common.Address) (event.Subscription, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155PlatformFeeInfoUpdated)
				if err := _TokenERC1155.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlatformFeeInfoUpdated is a log parse operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_TokenERC1155 *TokenERC1155Filterer) ParsePlatformFeeInfoUpdated(log types.Log) (*TokenERC1155PlatformFeeInfoUpdated, error) {
	event := new(TokenERC1155PlatformFeeInfoUpdated)
	if err := _TokenERC1155.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155PlatformFeeTypeUpdatedIterator is returned from FilterPlatformFeeTypeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeTypeUpdated events raised by the TokenERC1155 contract.
type TokenERC1155PlatformFeeTypeUpdatedIterator struct {
	Event *TokenERC1155PlatformFeeTypeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155PlatformFeeTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155PlatformFeeTypeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155PlatformFeeTypeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155PlatformFeeTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155PlatformFeeTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155PlatformFeeTypeUpdated represents a PlatformFeeTypeUpdated event raised by the TokenERC1155 contract.
type TokenERC1155PlatformFeeTypeUpdated struct {
	FeeType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeTypeUpdated is a free log retrieval operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_TokenERC1155 *TokenERC1155Filterer) FilterPlatformFeeTypeUpdated(opts *bind.FilterOpts) (*TokenERC1155PlatformFeeTypeUpdatedIterator, error) {

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenERC1155PlatformFeeTypeUpdatedIterator{contract: _TokenERC1155.contract, event: "PlatformFeeTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeTypeUpdated is a free log subscription operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_TokenERC1155 *TokenERC1155Filterer) WatchPlatformFeeTypeUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC1155PlatformFeeTypeUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155PlatformFeeTypeUpdated)
				if err := _TokenERC1155.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePlatformFeeTypeUpdated is a log parse operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_TokenERC1155 *TokenERC1155Filterer) ParsePlatformFeeTypeUpdated(log types.Log) (*TokenERC1155PlatformFeeTypeUpdated, error) {
	event := new(TokenERC1155PlatformFeeTypeUpdated)
	if err := _TokenERC1155.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155PrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the TokenERC1155 contract.
type TokenERC1155PrimarySaleRecipientUpdatedIterator struct {
	Event *TokenERC1155PrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155PrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155PrimarySaleRecipientUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155PrimarySaleRecipientUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155PrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155PrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155PrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the TokenERC1155 contract.
type TokenERC1155PrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_TokenERC1155 *TokenERC1155Filterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*TokenERC1155PrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155PrimarySaleRecipientUpdatedIterator{contract: _TokenERC1155.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_TokenERC1155 *TokenERC1155Filterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC1155PrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155PrimarySaleRecipientUpdated)
				if err := _TokenERC1155.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePrimarySaleRecipientUpdated is a log parse operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_TokenERC1155 *TokenERC1155Filterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*TokenERC1155PrimarySaleRecipientUpdated, error) {
	event := new(TokenERC1155PrimarySaleRecipientUpdated)
	if err := _TokenERC1155.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TokenERC1155 contract.
type TokenERC1155RoleAdminChangedIterator struct {
	Event *TokenERC1155RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155RoleAdminChanged represents a RoleAdminChanged event raised by the TokenERC1155 contract.
type TokenERC1155RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenERC1155 *TokenERC1155Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenERC1155RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155RoleAdminChangedIterator{contract: _TokenERC1155.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenERC1155 *TokenERC1155Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenERC1155RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155RoleAdminChanged)
				if err := _TokenERC1155.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenERC1155 *TokenERC1155Filterer) ParseRoleAdminChanged(log types.Log) (*TokenERC1155RoleAdminChanged, error) {
	event := new(TokenERC1155RoleAdminChanged)
	if err := _TokenERC1155.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TokenERC1155 contract.
type TokenERC1155RoleGrantedIterator struct {
	Event *TokenERC1155RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155RoleGranted represents a RoleGranted event raised by the TokenERC1155 contract.
type TokenERC1155RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenERC1155RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155RoleGrantedIterator{contract: _TokenERC1155.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenERC1155RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155RoleGranted)
				if err := _TokenERC1155.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) ParseRoleGranted(log types.Log) (*TokenERC1155RoleGranted, error) {
	event := new(TokenERC1155RoleGranted)
	if err := _TokenERC1155.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TokenERC1155 contract.
type TokenERC1155RoleRevokedIterator struct {
	Event *TokenERC1155RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155RoleRevoked represents a RoleRevoked event raised by the TokenERC1155 contract.
type TokenERC1155RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenERC1155RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155RoleRevokedIterator{contract: _TokenERC1155.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenERC1155RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155RoleRevoked)
				if err := _TokenERC1155.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC1155 *TokenERC1155Filterer) ParseRoleRevoked(log types.Log) (*TokenERC1155RoleRevoked, error) {
	event := new(TokenERC1155RoleRevoked)
	if err := _TokenERC1155.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155RoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the TokenERC1155 contract.
type TokenERC1155RoyaltyForTokenIterator struct {
	Event *TokenERC1155RoyaltyForToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155RoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155RoyaltyForToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155RoyaltyForToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155RoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155RoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155RoyaltyForToken represents a RoyaltyForToken event raised by the TokenERC1155 contract.
type TokenERC1155RoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*TokenERC1155RoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155RoyaltyForTokenIterator{contract: _TokenERC1155.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *TokenERC1155RoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155RoyaltyForToken)
				if err := _TokenERC1155.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoyaltyForToken is a log parse operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_TokenERC1155 *TokenERC1155Filterer) ParseRoyaltyForToken(log types.Log) (*TokenERC1155RoyaltyForToken, error) {
	event := new(TokenERC1155RoyaltyForToken)
	if err := _TokenERC1155.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155TokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the TokenERC1155 contract.
type TokenERC1155TokensMintedIterator struct {
	Event *TokenERC1155TokensMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155TokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155TokensMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155TokensMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155TokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155TokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155TokensMinted represents a TokensMinted event raised by the TokenERC1155 contract.
type TokenERC1155TokensMinted struct {
	MintedTo       common.Address
	TokenIdMinted  *big.Int
	Uri            string
	QuantityMinted *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x04133ee4cb027e1c5fce5e3481289278a93bd16a65a3b65b428a6d239e706bfb.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri, uint256 quantityMinted)
func (_TokenERC1155 *TokenERC1155Filterer) FilterTokensMinted(opts *bind.FilterOpts, mintedTo []common.Address, tokenIdMinted []*big.Int) (*TokenERC1155TokensMintedIterator, error) {

	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "TokensMinted", mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155TokensMintedIterator{contract: _TokenERC1155.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x04133ee4cb027e1c5fce5e3481289278a93bd16a65a3b65b428a6d239e706bfb.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri, uint256 quantityMinted)
func (_TokenERC1155 *TokenERC1155Filterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *TokenERC1155TokensMinted, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "TokensMinted", mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155TokensMinted)
				if err := _TokenERC1155.contract.UnpackLog(event, "TokensMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMinted is a log parse operation binding the contract event 0x04133ee4cb027e1c5fce5e3481289278a93bd16a65a3b65b428a6d239e706bfb.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri, uint256 quantityMinted)
func (_TokenERC1155 *TokenERC1155Filterer) ParseTokensMinted(log types.Log) (*TokenERC1155TokensMinted, error) {
	event := new(TokenERC1155TokensMinted)
	if err := _TokenERC1155.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155TokensMintedWithSignatureIterator is returned from FilterTokensMintedWithSignature and is used to iterate over the raw logs and unpacked data for TokensMintedWithSignature events raised by the TokenERC1155 contract.
type TokenERC1155TokensMintedWithSignatureIterator struct {
	Event *TokenERC1155TokensMintedWithSignature // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155TokensMintedWithSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155TokensMintedWithSignature)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155TokensMintedWithSignature)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155TokensMintedWithSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155TokensMintedWithSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155TokensMintedWithSignature represents a TokensMintedWithSignature event raised by the TokenERC1155 contract.
type TokenERC1155TokensMintedWithSignature struct {
	Signer        common.Address
	MintedTo      common.Address
	TokenIdMinted *big.Int
	MintRequest   ITokenERC1155MintRequest
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensMintedWithSignature is a free log retrieval operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC1155 *TokenERC1155Filterer) FilterTokensMintedWithSignature(opts *bind.FilterOpts, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (*TokenERC1155TokensMintedWithSignatureIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155TokensMintedWithSignatureIterator{contract: _TokenERC1155.contract, event: "TokensMintedWithSignature", logs: logs, sub: sub}, nil
}

// WatchTokensMintedWithSignature is a free log subscription operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC1155 *TokenERC1155Filterer) WatchTokensMintedWithSignature(opts *bind.WatchOpts, sink chan<- *TokenERC1155TokensMintedWithSignature, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155TokensMintedWithSignature)
				if err := _TokenERC1155.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMintedWithSignature is a log parse operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC1155 *TokenERC1155Filterer) ParseTokensMintedWithSignature(log types.Log) (*TokenERC1155TokensMintedWithSignature, error) {
	event := new(TokenERC1155TokensMintedWithSignature)
	if err := _TokenERC1155.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the TokenERC1155 contract.
type TokenERC1155TransferBatchIterator struct {
	Event *TokenERC1155TransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155TransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155TransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155TransferBatch represents a TransferBatch event raised by the TokenERC1155 contract.
type TokenERC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenERC1155 *TokenERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenERC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155TransferBatchIterator{contract: _TokenERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenERC1155 *TokenERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *TokenERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155TransferBatch)
				if err := _TokenERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_TokenERC1155 *TokenERC1155Filterer) ParseTransferBatch(log types.Log) (*TokenERC1155TransferBatch, error) {
	event := new(TokenERC1155TransferBatch)
	if err := _TokenERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the TokenERC1155 contract.
type TokenERC1155TransferSingleIterator struct {
	Event *TokenERC1155TransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155TransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155TransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155TransferSingle represents a TransferSingle event raised by the TokenERC1155 contract.
type TokenERC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenERC1155 *TokenERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*TokenERC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155TransferSingleIterator{contract: _TokenERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenERC1155 *TokenERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *TokenERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155TransferSingle)
				if err := _TokenERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_TokenERC1155 *TokenERC1155Filterer) ParseTransferSingle(log types.Log) (*TokenERC1155TransferSingle, error) {
	event := new(TokenERC1155TransferSingle)
	if err := _TokenERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the TokenERC1155 contract.
type TokenERC1155URIIterator struct {
	Event *TokenERC1155URI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC1155URI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenERC1155URI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC1155URI represents a URI event raised by the TokenERC1155 contract.
type TokenERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenERC1155 *TokenERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*TokenERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TokenERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC1155URIIterator{contract: _TokenERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenERC1155 *TokenERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *TokenERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _TokenERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC1155URI)
				if err := _TokenERC1155.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_TokenERC1155 *TokenERC1155Filterer) ParseURI(log types.Log) (*TokenERC1155URI, error) {
	event := new(TokenERC1155URI)
	if err := _TokenERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
