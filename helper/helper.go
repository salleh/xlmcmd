package helper

func IsValidStellarNetwork(network string) bool {
	return network == "public" || network == "testnet"
}
