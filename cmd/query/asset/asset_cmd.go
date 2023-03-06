package asset

import (
	"github.com/spf13/cobra"
)

var assetQueryCmd = &cobra.Command{
	Use:   "asset",
	Short: "asset queries",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
}

func init() {
	assetQueryCmd.AddCommand(getAssetDetailsCmd())
}

func GetAssetCmd() *cobra.Command {
	return assetQueryCmd
}
