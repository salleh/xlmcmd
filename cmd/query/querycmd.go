package query

import (
	"github.com/spf13/cobra"
)

var stellarNetwork string

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "xmlcmd provides command line utility to query Stellar blockchain",
	Long:  "xmlcmd provides command line utility to query Stellar blockchain based on the given options as parameters.",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
}

func init() {
	queryCmd.PersistentFlags().StringVarP(
		&stellarNetwork,
		"network",
		"n",
		"public",
		`Stellar network to connect to: public | testnet`)

	queryCmd.AddCommand(GetBalanceCmd())
}

func GetQueryCmd() *cobra.Command {
	return queryCmd
}
