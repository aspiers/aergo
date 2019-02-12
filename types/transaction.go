package types

import (
	"bytes"
	"math/big"
)

type Transaction interface {
	GetTx() *Tx
	GetBody() *TxBody
	GetHash() []byte
	CalculateTxHash() []byte
	Validate() error
	ValidateWithSenderState(senderState *State, fee *big.Int) error
	HasVerifedAccount() bool
	GetVerifedAccount() Address
	SetVerifedAccount(account Address) bool
	RemoveVerifedAccount() bool
}

type transaction struct {
	Tx              *Tx
	VerifiedAccount Address
}

func NewTransaction(tx *Tx) Transaction {
	return &transaction{Tx: tx}
}

func (tx *transaction) GetTx() *Tx {
	return tx.Tx
}

func (tx *transaction) GetBody() *TxBody {
	return tx.Tx.Body
}

func (tx *transaction) GetHash() []byte {
	return tx.Tx.Hash
}

func (tx *transaction) CalculateTxHash() []byte {
	return tx.Tx.CalculateTxHash()
}

func (tx *transaction) Validate() error {
	account := tx.GetBody().GetAccount()
	if account == nil {
		return ErrTxFormatInvalid
	}

	if !bytes.Equal(tx.GetHash(), tx.CalculateTxHash()) {
		return ErrTxHasInvalidHash
	}

	amount := tx.GetBody().GetAmountBigInt()
	if amount.Cmp(MaxAER) > 0 {
		return ErrTxInvalidAmount
	}

	price := tx.GetBody().GetPriceBigInt()
	if price.Cmp(MaxAER) > 0 {
		return ErrTxInvalidPrice
	}

	if len(tx.GetBody().GetAccount()) > AddressLength {
		return ErrTxInvalidAccount
	}

	if len(tx.GetBody().GetRecipient()) > AddressLength {
		return ErrTxInvalidRecipient
	}

	switch tx.GetBody().Type {
	case TxType_NORMAL:
		if tx.GetBody().GetRecipient() == nil && len(tx.GetBody().GetPayload()) == 0 {
			//contract deploy
			return ErrTxInvalidRecipient
		}
	case TxType_GOVERNANCE:
		if len(tx.GetBody().Payload) <= 0 {
			return ErrTxFormatInvalid
		}
		switch string(tx.GetBody().GetRecipient()) {
		case AergoSystem:
			if (tx.GetBody().GetPayload()[0] == 's' || tx.GetBody().GetPayload()[0] == 'u') &&
				amount.Cmp(StakingMinimum) < 0 {
				return ErrTooSmallAmount
			}
		case AergoName:
			if tx.GetBody().GetPayload()[0] != 'c' &&
				tx.GetBody().GetPayload()[0] != 'b' &&
				tx.GetBody().GetPayload()[0] != 'u' {
				return ErrTxFormatInvalid
			}
			if new(big.Int).SetUint64(1000000000000000000).Cmp(tx.GetBody().GetAmountBigInt()) > 0 {
				return ErrTooSmallAmount
			}
		default:
			return ErrTxInvalidRecipient
		}
	default:
		return ErrTxInvalidType
	}
	return nil
}

func (tx *transaction) ValidateWithSenderState(senderState *State, fee *big.Int) error {
	if (senderState.GetNonce() + 1) > tx.GetBody().GetNonce() {
		return ErrTxNonceTooLow
	}
	amount := tx.GetBody().GetAmountBigInt()
	balance := senderState.GetBalanceBigInt()
	switch tx.GetBody().GetType() {
	case TxType_NORMAL:
		spending := new(big.Int).Add(amount, fee)
		if spending.Cmp(balance) > 0 {
			return ErrInsufficientBalance
		}
	case TxType_GOVERNANCE:
		switch string(tx.GetBody().GetRecipient()) {
		case AergoSystem:
			if tx.GetBody().GetPayload()[0] == 's' &&
				amount.Cmp(balance) > 0 {
				return ErrInsufficientBalance
			}
		case AergoName:
			if (tx.GetBody().GetPayload()[0] == 'c' ||
				tx.GetBody().GetPayload()[0] == 'b') &&
				amount.Cmp(balance) > 0 {
				return ErrInsufficientBalance
			}
		default:
			return ErrTxInvalidRecipient
		}
	}
	if (senderState.GetNonce() + 1) < tx.GetBody().GetNonce() {
		return ErrTxNonceToohigh
	}
	return nil
}

//TODO : refoctor after ContractState move to types
func (tx *Tx) ValidateWithContractState(contractState *State) error {
	//in system.ValidateSystemTx
	//in name.ValidateNameTx
	return nil
}

func (tx *transaction) GetVerifedAccount() Address {
	return tx.VerifiedAccount
}

func (tx *transaction) SetVerifedAccount(account Address) bool {
	tx.VerifiedAccount = account
	return true
}

func (tx *transaction) HasVerifedAccount() bool {
	return len(tx.VerifiedAccount) != 0
}

func (tx *transaction) RemoveVerifedAccount() bool {
	return tx.SetVerifedAccount(nil)
}

func (tx *transaction) Clone() *transaction {
	if tx == nil {
		return nil
	}
	if tx.GetBody() == nil {
		return &transaction{}
	}
	body := &TxBody{
		Nonce:     tx.GetBody().Nonce,
		Account:   Clone(tx.GetBody().Account).([]byte),
		Recipient: Clone(tx.GetBody().Recipient).([]byte),
		Amount:    Clone(tx.GetBody().Amount).([]byte),
		Payload:   Clone(tx.GetBody().Payload).([]byte),
		Limit:     tx.GetBody().Limit,
		Price:     Clone(tx.GetBody().Price).([]byte),
		Type:      tx.GetBody().Type,
		Sign:      Clone(tx.GetBody().Sign).([]byte),
	}
	res := &transaction{
		Tx: &Tx{Body: body},
	}
	res.Tx.Hash = res.CalculateTxHash()
	return res
}