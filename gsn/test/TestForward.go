// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package test

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/eywa-protocol/wrappers"
	"github.com/eywa-protocol/wrappers/gsn"
	"github.com/sirupsen/logrus"
)

// TestForwardForwardRequest is an auto generated low-level Go binding around an user-defined struct.
type TestForwardForwardRequest struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Gas   *big.Int
	Nonce *big.Int
	Data  []byte
}

func GsnBridgeFoo(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, _val *big.Int, _str string) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("foo", _val, _str)
	if err != nil {
		return
	}

	__forwarder, err := __gsnCaller.GetForwarder(__chainId)
	if err != nil {
		return
	}

	__forwarderAddress, err := __gsnCaller.GetForwarderAddress(__chainId)
	if err != nil {
		return
	}

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

	__nonce, err := __forwarder.GetNonce(&bind.CallOpts{}, __signerAddress)
	if err != nil {

		return
	}

	__req := &wrappers.IForwarderForwardRequest{
		From:  __signerAddress,
		To:    __contractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: __nonce,
		Data:  __fRequest,
	}

	__typedData, err := gsn.NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"foo", _val, _str)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := gsn.NewSignature(__typedData, __signer)
	if err != nil {
		return
	}

	__domainSeparatorHash, err := gsn.NewDomainSeparatorHash(__typedData)
	if err != nil {
		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {
		return
	}

	__reqTypeHash, err := gsn.NewRequestTypeHash(__genericParams)
	if err != nil {
		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}

func GsnBridgeTestExecute(
	__gsnCaller gsn.GsnCaller,
	__chainId *big.Int,
	__signer *ecdsa.PrivateKey,
	__contractAddress common.Address, req TestForwardForwardRequest, domainSeparator [32]byte, requestTypeHash [32]byte, suffixData []byte, sig []byte) (txHash common.Hash, err error) {

	__contractABI, err := abi.JSON(strings.NewReader(wrappers.BridgeABI))
	if err != nil {
		return common.Hash{}, fmt.Errorf("could not parse ABI: %w", err)
	}

	__fRequest, err := __contractABI.Pack("testExecute", req, domainSeparator, requestTypeHash, suffixData, sig)
	if err != nil {
		return
	}

	__forwarder, err := __gsnCaller.GetForwarder(__chainId)
	if err != nil {
		return
	}

	__forwarderAddress, err := __gsnCaller.GetForwarderAddress(__chainId)
	if err != nil {
		return
	}

	logrus.Infof("forwarderAddress: %s", __forwarderAddress.String())
	logrus.Infof("nodeRegistryAddress: %s", __contractAddress.String())

	__signerAddress := crypto.PubkeyToAddress(__signer.PublicKey)

	logrus.Infof("ownerAddress: %s", __signerAddress.String())

	__nonce, err := __forwarder.GetNonce(&bind.CallOpts{}, __signerAddress)
	if err != nil {

		return
	}

	__req := &wrappers.IForwarderForwardRequest{
		From:  __signerAddress,
		To:    __contractAddress,
		Value: big.NewInt(0),
		Gas:   big.NewInt(1e6),
		Nonce: __nonce,
		Data:  __fRequest,
	}

	__typedData, err := gsn.NewForwardRequestTypedData(
		__req,
		__forwarderAddress.String(),
		wrappers.BridgeABI,
		"testExecute", req, domainSeparator, requestTypeHash, suffixData, sig)
	if err != nil {
		return
	}

	__typedDataSignature, _, err := gsn.NewSignature(__typedData, __signer)
	if err != nil {
		return
	}

	__domainSeparatorHash, err := gsn.NewDomainSeparatorHash(__typedData)
	if err != nil {
		return
	}

	__genericParams, err := __forwarder.GENERICPARAMS(&bind.CallOpts{})
	if err != nil {
		return
	}

	__reqTypeHash, err := gsn.NewRequestTypeHash(__genericParams)
	if err != nil {
		return
	}

	return __gsnCaller.Execute(__chainId, *__req, __domainSeparatorHash, __reqTypeHash, nil, __typedDataSignature)
}