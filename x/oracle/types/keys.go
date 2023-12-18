package types

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName is the name of the oracle module
	ModuleName = "oracle"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouterKey is the msg router key for the oracle module
	RouterKey = ModuleName

	// QuerierRoute is the query router key for the oracle module
	QuerierRoute = ModuleName
)

// Keys for oracle store
// Items are stored with the following key: values
//
// - 0x01<denom_Bytes>: sdk.Dec
//
// - 0x02<valAddress_Bytes>: accAddress
//
// - 0x03<valAddress_Bytes>: int64
//
// - 0x04<valAddress_Bytes>: AggregateExchangeRatePrevote
//
// - 0x05<valAddress_Bytes>: AggregateExchangeRateVote
//
// - 0x06<denom_Bytes>: sdk.Dec
var (
	// Keys for store prefixes
	ExchangeRateKey                 = []byte{0x01} // prefix for each key to a rate
	FeederDelegationKey             = []byte{0x02} // prefix for each key to a feeder delegation
	MissCounterKey                  = []byte{0x03} // prefix for each key to a miss counter
	AggregateExchangeRatePrevoteKey = []byte{0x04} // prefix for each key to a aggregate prevote
	AggregateExchangeRateVoteKey    = []byte{0x05} // prefix for each key to a aggregate vote
	ParamsKey                       = []byte{0x06}
	WhitelistKey                    = []byte{0x07}
	CounterKey                      = collections.NewPrefix(8)
	PricesKey                       = collections.NewPrefix(9)
)

// GetExchangeRateKey - stored by *denom*
func GetExchangeRateKey(denom string) []byte {
	return append(ExchangeRateKey, []byte(denom)...)
}

// GetFeederDelegationKey - stored by *Validator* address
func GetFeederDelegationKey(v sdk.ValAddress) []byte {
	return append(FeederDelegationKey, address.MustLengthPrefix(v)...)
}

// GetMissCounterKey - stored by *Validator* address
func GetMissCounterKey(v sdk.ValAddress) []byte {
	return append(MissCounterKey, address.MustLengthPrefix(v)...)
}

// GetAggregateExchangeRatePrevoteKey - stored by *Validator* address
func GetAggregateExchangeRatePrevoteKey(v sdk.ValAddress) []byte {
	return append(AggregateExchangeRatePrevoteKey, address.MustLengthPrefix(v)...)
}

// GetAggregateExchangeRateVoteKey - stored by *Validator* address
func GetAggregateExchangeRateVoteKey(v sdk.ValAddress) []byte {
	return append(AggregateExchangeRateVoteKey, address.MustLengthPrefix(v)...)
}
