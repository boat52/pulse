package params

import "github.com/cosmos/cosmos-sdk/types"

const (
	Bech32PrefixAccAddr  = "pulse"
	Bech32PrefixAccPub   = "pulsepub"
	Bech32PrefixValAddr  = "pulsevaloper"
	Bech32PrefixValPub   = "pulsevaloperpub"
	Bech32PrefixConsAddr = "pulsevalcons"
	Bech32PrefixConsPub  = "pulsevalconspub"
)

func SetAddressPrefixes() {
	config := types.GetConfig()
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	config.Seal()
}
