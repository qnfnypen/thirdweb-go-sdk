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

// ITokenERC721MintRequest is an auto generated low-level Go binding around an user-defined struct.
type ITokenERC721MintRequest struct {
	To                     common.Address
	RoyaltyRecipient       common.Address
	RoyaltyBps             *big.Int
	PrimarySaleRecipient   common.Address
	Uri                    string
	Price                  *big.Int
	Currency               common.Address
	ValidityStartTimestamp *big.Int
	ValidityEndTimestamp   *big.Int
	Uid                    [32]byte
}

// TokenERC721MetaData contains all meta data concerning the TokenERC721 contract.
var TokenERC721MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"name\":\"\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"approved\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"indexed\":false,\"internalType\":\"bool\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_fromTokenId\",\"indexed\":false,\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"_toTokenId\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultRoyalty\",\"inputs\":[{\"type\":\"address\",\"name\":\"newRoyaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"newRoyaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FlatPlatformFeeUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":false,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"flatFee\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"indexed\":false,\"internalType\":\"uint8\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataFrozen\",\"inputs\":[],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnerUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"prevOwner\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"newOwner\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeInfoUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"platformFeeRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"platformFeeBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PlatformFeeTypeUpdated\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"feeType\",\"indexed\":false,\"internalType\":\"enumIPlatformFee.PlatformFeeType\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrimarySaleRecipientUpdated\",\"inputs\":[{\"type\":\"address\",\"name\":\"recipient\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"previousAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"newAdminRole\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"sender\",\"indexed\":true,\"internalType\":\"address\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoyaltyForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"indexed\":false,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensMinted\",\"inputs\":[{\"type\":\"address\",\"name\":\"mintedTo\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenIdMinted\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"uri\",\"indexed\":false,\"internalType\":\"string\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensMintedWithSignature\",\"inputs\":[{\"type\":\"address\",\"name\":\"signer\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintedTo\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenIdMinted\",\"indexed\":true,\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"mintRequest\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structITokenERC721.MintRequest\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"indexed\":true,\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"indexed\":true,\"internalType\":\"uint256\"}],\"outputs\":[],\"anonymous\":false},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contractType\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"contractURI\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contractVersion\",\"inputs\":[],\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"type\":\"bytes1\",\"name\":\"fields\",\"internalType\":\"bytes1\"},{\"type\":\"string\",\"name\":\"name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"version\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"chainId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"verifyingContract\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"uint256[]\",\"name\":\"extensions\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"freezeMetadata\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDefaultRoyaltyInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformFeeInfo\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"uint16\",\"name\":\"\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_defaultAdmin\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_contractURI\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"_trustedForwarders\",\"internalType\":\"address[]\"},{\"type\":\"address\",\"name\":\"_saleRecipient\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"_royaltyBps\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint128\"},{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTrustedForwarder\",\"inputs\":[{\"type\":\"address\",\"name\":\"forwarder\",\"internalType\":\"address\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintTo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_to\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintWithSignature\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_req\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"internalType\":\"structITokenERC721.MintRequest\"},{\"type\":\"bytes\",\"name\":\"_signature\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"tokenIdMinted\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"type\":\"bytes[]\",\"name\":\"results\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextTokenIdToMint\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"platformFeeRecipient\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"primarySaleRecipient\",\"inputs\":[],\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"role\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"royaltyInfo\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"salePrice\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"address\",\"name\":\"receiver\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyAmount\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractURI\",\"inputs\":[{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDefaultRoyaltyInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_royaltyBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOwner\",\"inputs\":[{\"type\":\"address\",\"name\":\"_newOwner\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFeeInfo\",\"inputs\":[{\"type\":\"address\",\"name\":\"_platformFeeRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_platformFeeBps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPrimarySaleRecipient\",\"inputs\":[{\"type\":\"address\",\"name\":\"_saleRecipient\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoyaltyInfoForToken\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"_recipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"_bps\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenByIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"_tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uriFrozen\",\"inputs\":[],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"_req\",\"components\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyRecipient\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"royaltyBps\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"primarySaleRecipient\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"address\"},{\"type\":\"uint128\",\"name\":\"validityStartTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"uint128\",\"name\":\"validityEndTimestamp\",\"internalType\":\"uint128\"},{\"type\":\"bytes32\",\"name\":\"uid\",\"internalType\":\"bytes32\"}],\"internalType\":\"structITokenERC721.MintRequest\"},{\"type\":\"bytes\",\"name\":\"_signature\",\"internalType\":\"bytes\"}],\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"},{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// TokenERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenERC721MetaData.ABI instead.
var TokenERC721ABI = TokenERC721MetaData.ABI

// TokenERC721 is an auto generated Go binding around an Ethereum contract.
type TokenERC721 struct {
	TokenERC721Caller     // Read-only binding to the contract
	TokenERC721Transactor // Write-only binding to the contract
	TokenERC721Filterer   // Log filterer for contract events
}

// TokenERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type TokenERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenERC721Session struct {
	Contract     *TokenERC721      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenERC721CallerSession struct {
	Contract *TokenERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenERC721TransactorSession struct {
	Contract     *TokenERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type TokenERC721Raw struct {
	Contract *TokenERC721 // Generic contract binding to access the raw methods on
}

// TokenERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenERC721CallerRaw struct {
	Contract *TokenERC721Caller // Generic read-only contract binding to access the raw methods on
}

// TokenERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenERC721TransactorRaw struct {
	Contract *TokenERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenERC721 creates a new instance of TokenERC721, bound to a specific deployed contract.
func NewTokenERC721(address common.Address, backend bind.ContractBackend) (*TokenERC721, error) {
	contract, err := bindTokenERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenERC721{TokenERC721Caller: TokenERC721Caller{contract: contract}, TokenERC721Transactor: TokenERC721Transactor{contract: contract}, TokenERC721Filterer: TokenERC721Filterer{contract: contract}}, nil
}

// NewTokenERC721Caller creates a new read-only instance of TokenERC721, bound to a specific deployed contract.
func NewTokenERC721Caller(address common.Address, caller bind.ContractCaller) (*TokenERC721Caller, error) {
	contract, err := bindTokenERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC721Caller{contract: contract}, nil
}

// NewTokenERC721Transactor creates a new write-only instance of TokenERC721, bound to a specific deployed contract.
func NewTokenERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*TokenERC721Transactor, error) {
	contract, err := bindTokenERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenERC721Transactor{contract: contract}, nil
}

// NewTokenERC721Filterer creates a new log filterer instance of TokenERC721, bound to a specific deployed contract.
func NewTokenERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*TokenERC721Filterer, error) {
	contract, err := bindTokenERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenERC721Filterer{contract: contract}, nil
}

// bindTokenERC721 binds a generic wrapper to an already deployed contract.
func bindTokenERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC721 *TokenERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC721.Contract.TokenERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC721 *TokenERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC721.Contract.TokenERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC721 *TokenERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC721.Contract.TokenERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenERC721 *TokenERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenERC721 *TokenERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenERC721 *TokenERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenERC721.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC721 *TokenERC721Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC721 *TokenERC721Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenERC721.Contract.DEFAULTADMINROLE(&_TokenERC721.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TokenERC721 *TokenERC721CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenERC721.Contract.DEFAULTADMINROLE(&_TokenERC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenERC721 *TokenERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenERC721.Contract.BalanceOf(&_TokenERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TokenERC721.Contract.BalanceOf(&_TokenERC721.CallOpts, owner)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC721 *TokenERC721Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC721 *TokenERC721Session) ContractType() ([32]byte, error) {
	return _TokenERC721.Contract.ContractType(&_TokenERC721.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_TokenERC721 *TokenERC721CallerSession) ContractType() ([32]byte, error) {
	return _TokenERC721.Contract.ContractType(&_TokenERC721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC721 *TokenERC721Caller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC721 *TokenERC721Session) ContractURI() (string, error) {
	return _TokenERC721.Contract.ContractURI(&_TokenERC721.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_TokenERC721 *TokenERC721CallerSession) ContractURI() (string, error) {
	return _TokenERC721.Contract.ContractURI(&_TokenERC721.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC721 *TokenERC721Caller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC721 *TokenERC721Session) ContractVersion() (uint8, error) {
	return _TokenERC721.Contract.ContractVersion(&_TokenERC721.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_TokenERC721 *TokenERC721CallerSession) ContractVersion() (uint8, error) {
	return _TokenERC721.Contract.ContractVersion(&_TokenERC721.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_TokenERC721 *TokenERC721Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "eip712Domain")

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
func (_TokenERC721 *TokenERC721Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _TokenERC721.Contract.Eip712Domain(&_TokenERC721.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_TokenERC721 *TokenERC721CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _TokenERC721.Contract.Eip712Domain(&_TokenERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.GetApproved(&_TokenERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.GetApproved(&_TokenERC721.CallOpts, tokenId)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_TokenERC721 *TokenERC721Caller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

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
func (_TokenERC721 *TokenERC721Session) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetDefaultRoyaltyInfo(&_TokenERC721.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_TokenERC721 *TokenERC721CallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetDefaultRoyaltyInfo(&_TokenERC721.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_TokenERC721 *TokenERC721Caller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getPlatformFeeInfo")

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
func (_TokenERC721 *TokenERC721Session) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetPlatformFeeInfo(&_TokenERC721.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_TokenERC721 *TokenERC721CallerSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetPlatformFeeInfo(&_TokenERC721.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC721 *TokenERC721Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC721 *TokenERC721Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenERC721.Contract.GetRoleAdmin(&_TokenERC721.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TokenERC721 *TokenERC721CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenERC721.Contract.GetRoleAdmin(&_TokenERC721.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC721 *TokenERC721Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC721 *TokenERC721Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.GetRoleMember(&_TokenERC721.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.GetRoleMember(&_TokenERC721.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC721 *TokenERC721Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenERC721.Contract.GetRoleMemberCount(&_TokenERC721.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenERC721.Contract.GetRoleMemberCount(&_TokenERC721.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_TokenERC721 *TokenERC721Caller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

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
func (_TokenERC721 *TokenERC721Session) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetRoyaltyInfoForToken(&_TokenERC721.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_TokenERC721 *TokenERC721CallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _TokenERC721.Contract.GetRoyaltyInfoForToken(&_TokenERC721.CallOpts, _tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC721 *TokenERC721Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC721 *TokenERC721Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenERC721.Contract.HasRole(&_TokenERC721.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TokenERC721 *TokenERC721CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenERC721.Contract.HasRole(&_TokenERC721.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenERC721 *TokenERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenERC721 *TokenERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TokenERC721.Contract.IsApprovedForAll(&_TokenERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TokenERC721 *TokenERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TokenERC721.Contract.IsApprovedForAll(&_TokenERC721.CallOpts, owner, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC721 *TokenERC721Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC721 *TokenERC721Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenERC721.Contract.IsTrustedForwarder(&_TokenERC721.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TokenERC721 *TokenERC721CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TokenERC721.Contract.IsTrustedForwarder(&_TokenERC721.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC721 *TokenERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC721 *TokenERC721Session) Name() (string, error) {
	return _TokenERC721.Contract.Name(&_TokenERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TokenERC721 *TokenERC721CallerSession) Name() (string, error) {
	return _TokenERC721.Contract.Name(&_TokenERC721.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC721 *TokenERC721Session) NextTokenIdToMint() (*big.Int, error) {
	return _TokenERC721.Contract.NextTokenIdToMint(&_TokenERC721.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _TokenERC721.Contract.NextTokenIdToMint(&_TokenERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC721 *TokenERC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC721 *TokenERC721Session) Owner() (common.Address, error) {
	return _TokenERC721.Contract.Owner(&_TokenERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) Owner() (common.Address, error) {
	return _TokenERC721.Contract.Owner(&_TokenERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.OwnerOf(&_TokenERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TokenERC721.Contract.OwnerOf(&_TokenERC721.CallOpts, tokenId)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC721 *TokenERC721Caller) PlatformFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "platformFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC721 *TokenERC721Session) PlatformFeeRecipient() (common.Address, error) {
	return _TokenERC721.Contract.PlatformFeeRecipient(&_TokenERC721.CallOpts)
}

// PlatformFeeRecipient is a free data retrieval call binding the contract method 0xeb13554f.
//
// Solidity: function platformFeeRecipient() view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) PlatformFeeRecipient() (common.Address, error) {
	return _TokenERC721.Contract.PlatformFeeRecipient(&_TokenERC721.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC721 *TokenERC721Caller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC721 *TokenERC721Session) PrimarySaleRecipient() (common.Address, error) {
	return _TokenERC721.Contract.PrimarySaleRecipient(&_TokenERC721.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_TokenERC721 *TokenERC721CallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _TokenERC721.Contract.PrimarySaleRecipient(&_TokenERC721.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_TokenERC721 *TokenERC721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_TokenERC721 *TokenERC721Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _TokenERC721.Contract.RoyaltyInfo(&_TokenERC721.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_TokenERC721 *TokenERC721CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _TokenERC721.Contract.RoyaltyInfo(&_TokenERC721.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC721 *TokenERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC721 *TokenERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenERC721.Contract.SupportsInterface(&_TokenERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenERC721 *TokenERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenERC721.Contract.SupportsInterface(&_TokenERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC721 *TokenERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC721 *TokenERC721Session) Symbol() (string, error) {
	return _TokenERC721.Contract.Symbol(&_TokenERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TokenERC721 *TokenERC721CallerSession) Symbol() (string, error) {
	return _TokenERC721.Contract.Symbol(&_TokenERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721Session) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TokenERC721.Contract.TokenByIndex(&_TokenERC721.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _TokenERC721.Contract.TokenByIndex(&_TokenERC721.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721Session) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TokenERC721.Contract.TokenOfOwnerByIndex(&_TokenERC721.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _TokenERC721.Contract.TokenOfOwnerByIndex(&_TokenERC721.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_TokenERC721 *TokenERC721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_TokenERC721 *TokenERC721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _TokenERC721.Contract.TokenURI(&_TokenERC721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_TokenERC721 *TokenERC721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _TokenERC721.Contract.TokenURI(&_TokenERC721.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC721 *TokenERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC721 *TokenERC721Session) TotalSupply() (*big.Int, error) {
	return _TokenERC721.Contract.TotalSupply(&_TokenERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TokenERC721 *TokenERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _TokenERC721.Contract.TotalSupply(&_TokenERC721.CallOpts)
}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC721 *TokenERC721Caller) UriFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "uriFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC721 *TokenERC721Session) UriFrozen() (bool, error) {
	return _TokenERC721.Contract.UriFrozen(&_TokenERC721.CallOpts)
}

// UriFrozen is a free data retrieval call binding the contract method 0x274e4a1d.
//
// Solidity: function uriFrozen() view returns(bool)
func (_TokenERC721 *TokenERC721CallerSession) UriFrozen() (bool, error) {
	return _TokenERC721.Contract.UriFrozen(&_TokenERC721.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0xde903774.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC721 *TokenERC721Caller) Verify(opts *bind.CallOpts, _req ITokenERC721MintRequest, _signature []byte) (bool, common.Address, error) {
	var out []interface{}
	err := _TokenERC721.contract.Call(opts, &out, "verify", _req, _signature)

	if err != nil {
		return *new(bool), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// Verify is a free data retrieval call binding the contract method 0xde903774.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC721 *TokenERC721Session) Verify(_req ITokenERC721MintRequest, _signature []byte) (bool, common.Address, error) {
	return _TokenERC721.Contract.Verify(&_TokenERC721.CallOpts, _req, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xde903774.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool, address)
func (_TokenERC721 *TokenERC721CallerSession) Verify(_req ITokenERC721MintRequest, _signature []byte) (bool, common.Address, error) {
	return _TokenERC721.Contract.Verify(&_TokenERC721.CallOpts, _req, _signature)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.Approve(&_TokenERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.Approve(&_TokenERC721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.Burn(&_TokenERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.Burn(&_TokenERC721.TransactOpts, tokenId)
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC721 *TokenERC721Transactor) FreezeMetadata(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "freezeMetadata")
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC721 *TokenERC721Session) FreezeMetadata() (*types.Transaction, error) {
	return _TokenERC721.Contract.FreezeMetadata(&_TokenERC721.TransactOpts)
}

// FreezeMetadata is a paid mutator transaction binding the contract method 0xd111515d.
//
// Solidity: function freezeMetadata() returns()
func (_TokenERC721 *TokenERC721TransactorSession) FreezeMetadata() (*types.Transaction, error) {
	return _TokenERC721.Contract.FreezeMetadata(&_TokenERC721.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.GrantRole(&_TokenERC721.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.GrantRole(&_TokenERC721.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC721 *TokenERC721Transactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC721 *TokenERC721Session) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.Initialize(&_TokenERC721.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_TokenERC721 *TokenERC721TransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.Initialize(&_TokenERC721.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address _to, string _uri) returns(uint256)
func (_TokenERC721 *TokenERC721Transactor) MintTo(opts *bind.TransactOpts, _to common.Address, _uri string) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "mintTo", _to, _uri)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address _to, string _uri) returns(uint256)
func (_TokenERC721 *TokenERC721Session) MintTo(_to common.Address, _uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.MintTo(&_TokenERC721.TransactOpts, _to, _uri)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address _to, string _uri) returns(uint256)
func (_TokenERC721 *TokenERC721TransactorSession) MintTo(_to common.Address, _uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.MintTo(&_TokenERC721.TransactOpts, _to, _uri)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x2c4510f8.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns(uint256 tokenIdMinted)
func (_TokenERC721 *TokenERC721Transactor) MintWithSignature(opts *bind.TransactOpts, _req ITokenERC721MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "mintWithSignature", _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x2c4510f8.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns(uint256 tokenIdMinted)
func (_TokenERC721 *TokenERC721Session) MintWithSignature(_req ITokenERC721MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.MintWithSignature(&_TokenERC721.TransactOpts, _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x2c4510f8.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) payable returns(uint256 tokenIdMinted)
func (_TokenERC721 *TokenERC721TransactorSession) MintWithSignature(_req ITokenERC721MintRequest, _signature []byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.MintWithSignature(&_TokenERC721.TransactOpts, _req, _signature)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC721 *TokenERC721Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC721 *TokenERC721Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.Multicall(&_TokenERC721.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TokenERC721 *TokenERC721TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.Multicall(&_TokenERC721.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.RenounceRole(&_TokenERC721.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.RenounceRole(&_TokenERC721.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.RevokeRole(&_TokenERC721.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TokenERC721 *TokenERC721TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.RevokeRole(&_TokenERC721.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SafeTransferFrom(&_TokenERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SafeTransferFrom(&_TokenERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenERC721 *TokenERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenERC721 *TokenERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.SafeTransferFrom0(&_TokenERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenERC721.Contract.SafeTransferFrom0(&_TokenERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC721 *TokenERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC721 *TokenERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetApprovalForAll(&_TokenERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetApprovalForAll(&_TokenERC721.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC721 *TokenERC721Transactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC721 *TokenERC721Session) SetContractURI(_uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetContractURI(&_TokenERC721.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetContractURI(&_TokenERC721.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC721 *TokenERC721Transactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC721 *TokenERC721Session) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetDefaultRoyaltyInfo(&_TokenERC721.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetDefaultRoyaltyInfo(&_TokenERC721.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC721 *TokenERC721Transactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC721 *TokenERC721Session) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetOwner(&_TokenERC721.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetOwner(&_TokenERC721.TransactOpts, _newOwner)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC721 *TokenERC721Transactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC721 *TokenERC721Session) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetPlatformFeeInfo(&_TokenERC721.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetPlatformFeeInfo(&_TokenERC721.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC721 *TokenERC721Transactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC721 *TokenERC721Session) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetPrimarySaleRecipient(&_TokenERC721.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetPrimarySaleRecipient(&_TokenERC721.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC721 *TokenERC721Transactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC721 *TokenERC721Session) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetRoyaltyInfoForToken(&_TokenERC721.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetRoyaltyInfoForToken(&_TokenERC721.TransactOpts, _tokenId, _recipient, _bps)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC721 *TokenERC721Transactor) SetTokenURI(opts *bind.TransactOpts, _tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "setTokenURI", _tokenId, _uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC721 *TokenERC721Session) SetTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetTokenURI(&_TokenERC721.TransactOpts, _tokenId, _uri)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 _tokenId, string _uri) returns()
func (_TokenERC721 *TokenERC721TransactorSession) SetTokenURI(_tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _TokenERC721.Contract.SetTokenURI(&_TokenERC721.TransactOpts, _tokenId, _uri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.TransferFrom(&_TokenERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TokenERC721 *TokenERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TokenERC721.Contract.TransferFrom(&_TokenERC721.TransactOpts, from, to, tokenId)
}

// TokenERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TokenERC721 contract.
type TokenERC721ApprovalIterator struct {
	Event *TokenERC721Approval // Event containing the contract specifics and raw log

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
func (it *TokenERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721Approval)
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
		it.Event = new(TokenERC721Approval)
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
func (it *TokenERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721Approval represents a Approval event raised by the TokenERC721 contract.
type TokenERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TokenERC721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721ApprovalIterator{contract: _TokenERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721Approval)
				if err := _TokenERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) ParseApproval(log types.Log) (*TokenERC721Approval, error) {
	event := new(TokenERC721Approval)
	if err := _TokenERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TokenERC721 contract.
type TokenERC721ApprovalForAllIterator struct {
	Event *TokenERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TokenERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721ApprovalForAll)
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
		it.Event = new(TokenERC721ApprovalForAll)
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
func (it *TokenERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721ApprovalForAll represents a ApprovalForAll event raised by the TokenERC721 contract.
type TokenERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenERC721 *TokenERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TokenERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721ApprovalForAllIterator{contract: _TokenERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenERC721 *TokenERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TokenERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721ApprovalForAll)
				if err := _TokenERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TokenERC721 *TokenERC721Filterer) ParseApprovalForAll(log types.Log) (*TokenERC721ApprovalForAll, error) {
	event := new(TokenERC721ApprovalForAll)
	if err := _TokenERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721BatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the TokenERC721 contract.
type TokenERC721BatchMetadataUpdateIterator struct {
	Event *TokenERC721BatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TokenERC721BatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721BatchMetadataUpdate)
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
		it.Event = new(TokenERC721BatchMetadataUpdate)
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
func (it *TokenERC721BatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721BatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721BatchMetadataUpdate represents a BatchMetadataUpdate event raised by the TokenERC721 contract.
type TokenERC721BatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TokenERC721 *TokenERC721Filterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TokenERC721BatchMetadataUpdateIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenERC721BatchMetadataUpdateIterator{contract: _TokenERC721.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TokenERC721 *TokenERC721Filterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TokenERC721BatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721BatchMetadataUpdate)
				if err := _TokenERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseBatchMetadataUpdate(log types.Log) (*TokenERC721BatchMetadataUpdate, error) {
	event := new(TokenERC721BatchMetadataUpdate)
	if err := _TokenERC721.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721DefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the TokenERC721 contract.
type TokenERC721DefaultRoyaltyIterator struct {
	Event *TokenERC721DefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *TokenERC721DefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721DefaultRoyalty)
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
		it.Event = new(TokenERC721DefaultRoyalty)
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
func (it *TokenERC721DefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721DefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721DefaultRoyalty represents a DefaultRoyalty event raised by the TokenERC721 contract.
type TokenERC721DefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_TokenERC721 *TokenERC721Filterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*TokenERC721DefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721DefaultRoyaltyIterator{contract: _TokenERC721.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_TokenERC721 *TokenERC721Filterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *TokenERC721DefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721DefaultRoyalty)
				if err := _TokenERC721.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseDefaultRoyalty(log types.Log) (*TokenERC721DefaultRoyalty, error) {
	event := new(TokenERC721DefaultRoyalty)
	if err := _TokenERC721.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the TokenERC721 contract.
type TokenERC721EIP712DomainChangedIterator struct {
	Event *TokenERC721EIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *TokenERC721EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721EIP712DomainChanged)
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
		it.Event = new(TokenERC721EIP712DomainChanged)
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
func (it *TokenERC721EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721EIP712DomainChanged represents a EIP712DomainChanged event raised by the TokenERC721 contract.
type TokenERC721EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_TokenERC721 *TokenERC721Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*TokenERC721EIP712DomainChangedIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &TokenERC721EIP712DomainChangedIterator{contract: _TokenERC721.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_TokenERC721 *TokenERC721Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *TokenERC721EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721EIP712DomainChanged)
				if err := _TokenERC721.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseEIP712DomainChanged(log types.Log) (*TokenERC721EIP712DomainChanged, error) {
	event := new(TokenERC721EIP712DomainChanged)
	if err := _TokenERC721.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721FlatPlatformFeeUpdatedIterator is returned from FilterFlatPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for FlatPlatformFeeUpdated events raised by the TokenERC721 contract.
type TokenERC721FlatPlatformFeeUpdatedIterator struct {
	Event *TokenERC721FlatPlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *TokenERC721FlatPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721FlatPlatformFeeUpdated)
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
		it.Event = new(TokenERC721FlatPlatformFeeUpdated)
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
func (it *TokenERC721FlatPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721FlatPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721FlatPlatformFeeUpdated represents a FlatPlatformFeeUpdated event raised by the TokenERC721 contract.
type TokenERC721FlatPlatformFeeUpdated struct {
	PlatformFeeRecipient common.Address
	FlatFee              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFlatPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_TokenERC721 *TokenERC721Filterer) FilterFlatPlatformFeeUpdated(opts *bind.FilterOpts) (*TokenERC721FlatPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenERC721FlatPlatformFeeUpdatedIterator{contract: _TokenERC721.contract, event: "FlatPlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFlatPlatformFeeUpdated is a free log subscription operation binding the contract event 0xf8086cee80709bd44c82f89dbca54115ebd05e840a88ab81df9cf5be9754eb63.
//
// Solidity: event FlatPlatformFeeUpdated(address platformFeeRecipient, uint256 flatFee)
func (_TokenERC721 *TokenERC721Filterer) WatchFlatPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC721FlatPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "FlatPlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721FlatPlatformFeeUpdated)
				if err := _TokenERC721.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseFlatPlatformFeeUpdated(log types.Log) (*TokenERC721FlatPlatformFeeUpdated, error) {
	event := new(TokenERC721FlatPlatformFeeUpdated)
	if err := _TokenERC721.contract.UnpackLog(event, "FlatPlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenERC721 contract.
type TokenERC721InitializedIterator struct {
	Event *TokenERC721Initialized // Event containing the contract specifics and raw log

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
func (it *TokenERC721InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721Initialized)
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
		it.Event = new(TokenERC721Initialized)
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
func (it *TokenERC721InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721Initialized represents a Initialized event raised by the TokenERC721 contract.
type TokenERC721Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenERC721 *TokenERC721Filterer) FilterInitialized(opts *bind.FilterOpts) (*TokenERC721InitializedIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenERC721InitializedIterator{contract: _TokenERC721.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenERC721 *TokenERC721Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenERC721Initialized) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721Initialized)
				if err := _TokenERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseInitialized(log types.Log) (*TokenERC721Initialized, error) {
	event := new(TokenERC721Initialized)
	if err := _TokenERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721MetadataFrozenIterator is returned from FilterMetadataFrozen and is used to iterate over the raw logs and unpacked data for MetadataFrozen events raised by the TokenERC721 contract.
type TokenERC721MetadataFrozenIterator struct {
	Event *TokenERC721MetadataFrozen // Event containing the contract specifics and raw log

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
func (it *TokenERC721MetadataFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721MetadataFrozen)
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
		it.Event = new(TokenERC721MetadataFrozen)
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
func (it *TokenERC721MetadataFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721MetadataFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721MetadataFrozen represents a MetadataFrozen event raised by the TokenERC721 contract.
type TokenERC721MetadataFrozen struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMetadataFrozen is a free log retrieval operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_TokenERC721 *TokenERC721Filterer) FilterMetadataFrozen(opts *bind.FilterOpts) (*TokenERC721MetadataFrozenIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return &TokenERC721MetadataFrozenIterator{contract: _TokenERC721.contract, event: "MetadataFrozen", logs: logs, sub: sub}, nil
}

// WatchMetadataFrozen is a free log subscription operation binding the contract event 0xeef043febddf4e1d1cf1f72ff1407b84e036e805aa0934418cb82095da8d7164.
//
// Solidity: event MetadataFrozen()
func (_TokenERC721 *TokenERC721Filterer) WatchMetadataFrozen(opts *bind.WatchOpts, sink chan<- *TokenERC721MetadataFrozen) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "MetadataFrozen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721MetadataFrozen)
				if err := _TokenERC721.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseMetadataFrozen(log types.Log) (*TokenERC721MetadataFrozen, error) {
	event := new(TokenERC721MetadataFrozen)
	if err := _TokenERC721.contract.UnpackLog(event, "MetadataFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721MetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the TokenERC721 contract.
type TokenERC721MetadataUpdateIterator struct {
	Event *TokenERC721MetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TokenERC721MetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721MetadataUpdate)
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
		it.Event = new(TokenERC721MetadataUpdate)
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
func (it *TokenERC721MetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721MetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721MetadataUpdate represents a MetadataUpdate event raised by the TokenERC721 contract.
type TokenERC721MetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TokenERC721 *TokenERC721Filterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*TokenERC721MetadataUpdateIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TokenERC721MetadataUpdateIterator{contract: _TokenERC721.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TokenERC721 *TokenERC721Filterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TokenERC721MetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721MetadataUpdate)
				if err := _TokenERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseMetadataUpdate(log types.Log) (*TokenERC721MetadataUpdate, error) {
	event := new(TokenERC721MetadataUpdate)
	if err := _TokenERC721.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721OwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the TokenERC721 contract.
type TokenERC721OwnerUpdatedIterator struct {
	Event *TokenERC721OwnerUpdated // Event containing the contract specifics and raw log

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
func (it *TokenERC721OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721OwnerUpdated)
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
		it.Event = new(TokenERC721OwnerUpdated)
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
func (it *TokenERC721OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721OwnerUpdated represents a OwnerUpdated event raised by the TokenERC721 contract.
type TokenERC721OwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_TokenERC721 *TokenERC721Filterer) FilterOwnerUpdated(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*TokenERC721OwnerUpdatedIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721OwnerUpdatedIterator{contract: _TokenERC721.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_TokenERC721 *TokenERC721Filterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC721OwnerUpdated, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721OwnerUpdated)
				if err := _TokenERC721.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseOwnerUpdated(log types.Log) (*TokenERC721OwnerUpdated, error) {
	event := new(TokenERC721OwnerUpdated)
	if err := _TokenERC721.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721PlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the TokenERC721 contract.
type TokenERC721PlatformFeeInfoUpdatedIterator struct {
	Event *TokenERC721PlatformFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *TokenERC721PlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721PlatformFeeInfoUpdated)
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
		it.Event = new(TokenERC721PlatformFeeInfoUpdated)
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
func (it *TokenERC721PlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721PlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721PlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the TokenERC721 contract.
type TokenERC721PlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	PlatformFeeBps       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_TokenERC721 *TokenERC721Filterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts, platformFeeRecipient []common.Address) (*TokenERC721PlatformFeeInfoUpdatedIterator, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721PlatformFeeInfoUpdatedIterator{contract: _TokenERC721.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_TokenERC721 *TokenERC721Filterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC721PlatformFeeInfoUpdated, platformFeeRecipient []common.Address) (event.Subscription, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721PlatformFeeInfoUpdated)
				if err := _TokenERC721.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParsePlatformFeeInfoUpdated(log types.Log) (*TokenERC721PlatformFeeInfoUpdated, error) {
	event := new(TokenERC721PlatformFeeInfoUpdated)
	if err := _TokenERC721.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721PlatformFeeTypeUpdatedIterator is returned from FilterPlatformFeeTypeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeTypeUpdated events raised by the TokenERC721 contract.
type TokenERC721PlatformFeeTypeUpdatedIterator struct {
	Event *TokenERC721PlatformFeeTypeUpdated // Event containing the contract specifics and raw log

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
func (it *TokenERC721PlatformFeeTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721PlatformFeeTypeUpdated)
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
		it.Event = new(TokenERC721PlatformFeeTypeUpdated)
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
func (it *TokenERC721PlatformFeeTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721PlatformFeeTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721PlatformFeeTypeUpdated represents a PlatformFeeTypeUpdated event raised by the TokenERC721 contract.
type TokenERC721PlatformFeeTypeUpdated struct {
	FeeType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeTypeUpdated is a free log retrieval operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_TokenERC721 *TokenERC721Filterer) FilterPlatformFeeTypeUpdated(opts *bind.FilterOpts) (*TokenERC721PlatformFeeTypeUpdatedIterator, error) {

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenERC721PlatformFeeTypeUpdatedIterator{contract: _TokenERC721.contract, event: "PlatformFeeTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeTypeUpdated is a free log subscription operation binding the contract event 0xd246da9440709ce0dd3f4fd669abc85ada012ab9774b8ecdcc5059ba1486b9c1.
//
// Solidity: event PlatformFeeTypeUpdated(uint8 feeType)
func (_TokenERC721 *TokenERC721Filterer) WatchPlatformFeeTypeUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC721PlatformFeeTypeUpdated) (event.Subscription, error) {

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "PlatformFeeTypeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721PlatformFeeTypeUpdated)
				if err := _TokenERC721.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParsePlatformFeeTypeUpdated(log types.Log) (*TokenERC721PlatformFeeTypeUpdated, error) {
	event := new(TokenERC721PlatformFeeTypeUpdated)
	if err := _TokenERC721.contract.UnpackLog(event, "PlatformFeeTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721PrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the TokenERC721 contract.
type TokenERC721PrimarySaleRecipientUpdatedIterator struct {
	Event *TokenERC721PrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *TokenERC721PrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721PrimarySaleRecipientUpdated)
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
		it.Event = new(TokenERC721PrimarySaleRecipientUpdated)
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
func (it *TokenERC721PrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721PrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721PrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the TokenERC721 contract.
type TokenERC721PrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_TokenERC721 *TokenERC721Filterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*TokenERC721PrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721PrimarySaleRecipientUpdatedIterator{contract: _TokenERC721.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_TokenERC721 *TokenERC721Filterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *TokenERC721PrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721PrimarySaleRecipientUpdated)
				if err := _TokenERC721.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*TokenERC721PrimarySaleRecipientUpdated, error) {
	event := new(TokenERC721PrimarySaleRecipientUpdated)
	if err := _TokenERC721.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TokenERC721 contract.
type TokenERC721RoleAdminChangedIterator struct {
	Event *TokenERC721RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TokenERC721RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721RoleAdminChanged)
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
		it.Event = new(TokenERC721RoleAdminChanged)
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
func (it *TokenERC721RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721RoleAdminChanged represents a RoleAdminChanged event raised by the TokenERC721 contract.
type TokenERC721RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenERC721 *TokenERC721Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenERC721RoleAdminChangedIterator, error) {

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

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721RoleAdminChangedIterator{contract: _TokenERC721.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TokenERC721 *TokenERC721Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenERC721RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721RoleAdminChanged)
				if err := _TokenERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseRoleAdminChanged(log types.Log) (*TokenERC721RoleAdminChanged, error) {
	event := new(TokenERC721RoleAdminChanged)
	if err := _TokenERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TokenERC721 contract.
type TokenERC721RoleGrantedIterator struct {
	Event *TokenERC721RoleGranted // Event containing the contract specifics and raw log

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
func (it *TokenERC721RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721RoleGranted)
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
		it.Event = new(TokenERC721RoleGranted)
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
func (it *TokenERC721RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721RoleGranted represents a RoleGranted event raised by the TokenERC721 contract.
type TokenERC721RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC721 *TokenERC721Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenERC721RoleGrantedIterator, error) {

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

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721RoleGrantedIterator{contract: _TokenERC721.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC721 *TokenERC721Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenERC721RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721RoleGranted)
				if err := _TokenERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseRoleGranted(log types.Log) (*TokenERC721RoleGranted, error) {
	event := new(TokenERC721RoleGranted)
	if err := _TokenERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TokenERC721 contract.
type TokenERC721RoleRevokedIterator struct {
	Event *TokenERC721RoleRevoked // Event containing the contract specifics and raw log

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
func (it *TokenERC721RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721RoleRevoked)
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
		it.Event = new(TokenERC721RoleRevoked)
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
func (it *TokenERC721RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721RoleRevoked represents a RoleRevoked event raised by the TokenERC721 contract.
type TokenERC721RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC721 *TokenERC721Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenERC721RoleRevokedIterator, error) {

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

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721RoleRevokedIterator{contract: _TokenERC721.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TokenERC721 *TokenERC721Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenERC721RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721RoleRevoked)
				if err := _TokenERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseRoleRevoked(log types.Log) (*TokenERC721RoleRevoked, error) {
	event := new(TokenERC721RoleRevoked)
	if err := _TokenERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721RoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the TokenERC721 contract.
type TokenERC721RoyaltyForTokenIterator struct {
	Event *TokenERC721RoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *TokenERC721RoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721RoyaltyForToken)
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
		it.Event = new(TokenERC721RoyaltyForToken)
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
func (it *TokenERC721RoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721RoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721RoyaltyForToken represents a RoyaltyForToken event raised by the TokenERC721 contract.
type TokenERC721RoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_TokenERC721 *TokenERC721Filterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*TokenERC721RoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721RoyaltyForTokenIterator{contract: _TokenERC721.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_TokenERC721 *TokenERC721Filterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *TokenERC721RoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721RoyaltyForToken)
				if err := _TokenERC721.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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
func (_TokenERC721 *TokenERC721Filterer) ParseRoyaltyForToken(log types.Log) (*TokenERC721RoyaltyForToken, error) {
	event := new(TokenERC721RoyaltyForToken)
	if err := _TokenERC721.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721TokensMintedIterator is returned from FilterTokensMinted and is used to iterate over the raw logs and unpacked data for TokensMinted events raised by the TokenERC721 contract.
type TokenERC721TokensMintedIterator struct {
	Event *TokenERC721TokensMinted // Event containing the contract specifics and raw log

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
func (it *TokenERC721TokensMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721TokensMinted)
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
		it.Event = new(TokenERC721TokensMinted)
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
func (it *TokenERC721TokensMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721TokensMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721TokensMinted represents a TokensMinted event raised by the TokenERC721 contract.
type TokenERC721TokensMinted struct {
	MintedTo      common.Address
	TokenIdMinted *big.Int
	Uri           string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensMinted is a free log retrieval operation binding the contract event 0x9d89e36eadf856db0ad9ffb5a569e07f95634dddd9501141ecf04820484ad0dc.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri)
func (_TokenERC721 *TokenERC721Filterer) FilterTokensMinted(opts *bind.FilterOpts, mintedTo []common.Address, tokenIdMinted []*big.Int) (*TokenERC721TokensMintedIterator, error) {

	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "TokensMinted", mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721TokensMintedIterator{contract: _TokenERC721.contract, event: "TokensMinted", logs: logs, sub: sub}, nil
}

// WatchTokensMinted is a free log subscription operation binding the contract event 0x9d89e36eadf856db0ad9ffb5a569e07f95634dddd9501141ecf04820484ad0dc.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri)
func (_TokenERC721 *TokenERC721Filterer) WatchTokensMinted(opts *bind.WatchOpts, sink chan<- *TokenERC721TokensMinted, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "TokensMinted", mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721TokensMinted)
				if err := _TokenERC721.contract.UnpackLog(event, "TokensMinted", log); err != nil {
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

// ParseTokensMinted is a log parse operation binding the contract event 0x9d89e36eadf856db0ad9ffb5a569e07f95634dddd9501141ecf04820484ad0dc.
//
// Solidity: event TokensMinted(address indexed mintedTo, uint256 indexed tokenIdMinted, string uri)
func (_TokenERC721 *TokenERC721Filterer) ParseTokensMinted(log types.Log) (*TokenERC721TokensMinted, error) {
	event := new(TokenERC721TokensMinted)
	if err := _TokenERC721.contract.UnpackLog(event, "TokensMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721TokensMintedWithSignatureIterator is returned from FilterTokensMintedWithSignature and is used to iterate over the raw logs and unpacked data for TokensMintedWithSignature events raised by the TokenERC721 contract.
type TokenERC721TokensMintedWithSignatureIterator struct {
	Event *TokenERC721TokensMintedWithSignature // Event containing the contract specifics and raw log

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
func (it *TokenERC721TokensMintedWithSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721TokensMintedWithSignature)
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
		it.Event = new(TokenERC721TokensMintedWithSignature)
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
func (it *TokenERC721TokensMintedWithSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721TokensMintedWithSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721TokensMintedWithSignature represents a TokensMintedWithSignature event raised by the TokenERC721 contract.
type TokenERC721TokensMintedWithSignature struct {
	Signer        common.Address
	MintedTo      common.Address
	TokenIdMinted *big.Int
	MintRequest   ITokenERC721MintRequest
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensMintedWithSignature is a free log retrieval operation binding the contract event 0x110d160a1bedeea919a88fbc4b2a9fb61b7e664084391b6ca2740db66fef80fe.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC721 *TokenERC721Filterer) FilterTokensMintedWithSignature(opts *bind.FilterOpts, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (*TokenERC721TokensMintedWithSignatureIterator, error) {

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

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721TokensMintedWithSignatureIterator{contract: _TokenERC721.contract, event: "TokensMintedWithSignature", logs: logs, sub: sub}, nil
}

// WatchTokensMintedWithSignature is a free log subscription operation binding the contract event 0x110d160a1bedeea919a88fbc4b2a9fb61b7e664084391b6ca2740db66fef80fe.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC721 *TokenERC721Filterer) WatchTokensMintedWithSignature(opts *bind.WatchOpts, sink chan<- *TokenERC721TokensMintedWithSignature, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721TokensMintedWithSignature)
				if err := _TokenERC721.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
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

// ParseTokensMintedWithSignature is a log parse operation binding the contract event 0x110d160a1bedeea919a88fbc4b2a9fb61b7e664084391b6ca2740db66fef80fe.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_TokenERC721 *TokenERC721Filterer) ParseTokensMintedWithSignature(log types.Log) (*TokenERC721TokensMintedWithSignature, error) {
	event := new(TokenERC721TokensMintedWithSignature)
	if err := _TokenERC721.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TokenERC721 contract.
type TokenERC721TransferIterator struct {
	Event *TokenERC721Transfer // Event containing the contract specifics and raw log

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
func (it *TokenERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenERC721Transfer)
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
		it.Event = new(TokenERC721Transfer)
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
func (it *TokenERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenERC721Transfer represents a Transfer event raised by the TokenERC721 contract.
type TokenERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TokenERC721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TokenERC721TransferIterator{contract: _TokenERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TokenERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenERC721Transfer)
				if err := _TokenERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TokenERC721 *TokenERC721Filterer) ParseTransfer(log types.Log) (*TokenERC721Transfer, error) {
	event := new(TokenERC721Transfer)
	if err := _TokenERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
