package types

import (
	"encoding/json"
	"strconv"
	"testing"

	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/stretchr/testify/assert"
)

var (
	TestNormal          = 0
	TestDuplicatePeerID = 1
	TestInvalidPeerID   = 2
	TestInvalidString   = 3
)

func TestGovernanceTypeTransaction(t *testing.T) {
	const testSender = "AmPNYHyzyh9zweLwDyuoiUuTVCdrdksxkRWDjVJS76WQLExa2Jr4"
	const testReceiver = "AmNhXiU3s2BN26v5B5hT2bbEjvSjqyrBY7DGnD9UqVcwkTrDYyJN"
	account, err := DecodeAddress(testSender)
	assert.NoError(t, err, "should success to decode test address")
	recipient, err := DecodeAddress(testReceiver)
	assert.NoError(t, err, "should success to decode test address")
	tx := &Tx{
		Body: &TxBody{
			Account: account,
			Payload: []byte(`{"Name":"v1unstake"}`),
			Type:    TxType_GOVERNANCE,
		},
	}
	transaction := NewTransaction(tx)
	err = transaction.Validate()
	assert.EqualError(t, ErrTxHasInvalidHash, err.Error(), "empty hash")

	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, ErrTxInvalidRecipient, err.Error(), "recipient null")

	transaction.GetTx().GetBody().Recipient = tx.Body.Payload
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, ErrTxInvalidRecipient, err.Error(), "wrong recipient case")

	transaction.GetTx().GetBody().Recipient = recipient
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, ErrTxInvalidRecipient, err.Error(), "recipient should be aergo.*")

	transaction.GetTx().GetBody().Recipient = []byte(AergoSystem)
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, ErrTooSmallAmount, err.Error(), "recipient should be aergo.*")

	transaction.GetTx().GetBody().Amount = StakingMinimum.Bytes()
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.NoError(t, err, "should success")

	transaction.GetTx().GetBody().Payload = buildVoteBPPayloadEx(2, TestInvalidString)
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, err, ErrTxInvalidPayload.Error(), "invalid string")

	transaction.GetTx().GetBody().Payload = buildVoteBPPayloadEx(2, TestInvalidPeerID)
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, err, ErrTxInvalidPayload.Error(), "invalid string")

	transaction.GetTx().GetBody().Payload = buildVoteBPPayloadEx(2, TestDuplicatePeerID)
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.EqualError(t, err, ErrTxInvalidPayload.Error(), "invalid string")

	transaction.GetTx().GetBody().Payload = buildVoteBPPayloadEx(2, TestNormal)
	transaction.GetTx().Hash = transaction.CalculateTxHash()
	err = transaction.Validate()
	assert.Error(t, err, "should success")
}

func buildVoteBPPayloadEx(count int, err int) []byte {
	var ci CallInfo
	ci.Name = VoteBP
	_, pub, _ := crypto.GenerateKeyPair(crypto.Secp256k1, 256)
	peerid, _ := peer.IDFromPublicKey(pub)
	for i := 0; i < count; i++ {
		if err == TestDuplicatePeerID {
			ci.Args = append(ci.Args, peer.IDB58Encode(peerid))
		} else if err == TestInvalidString {
			ci.Args = append(ci.Args, (i + 1))
		} else if err == TestInvalidPeerID {
			ci.Args = append(ci.Args, string(i+1))
		} else {
			ci.Args = append(ci.Args, strconv.Itoa(i+1))
		}
	}
	payload, _ := json.Marshal(ci)
	return payload
}