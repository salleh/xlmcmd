package query

import (
	"github.com/salleh/xlmcmd/cmd/query/account"
	"github.com/salleh/xlmcmd/cmd/query/asset"
	"github.com/salleh/xlmcmd/cmd/query/lumens"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "get",
	Short: "query Stellar blockchain",
	Long:  "query objects in Stellar blockchain",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
}

func init() {
	queryCmd.AddCommand(account.GetAccountCmd())
	queryCmd.AddCommand(asset.GetAssetCmd())
	queryCmd.AddCommand(lumens.GetLumensCmd())
}

func GetQueryCmd() *cobra.Command {
	return queryCmd
}
