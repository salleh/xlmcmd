package lumens

import (
	"github.com/spf13/cobra"
)

var lumensQueryCmd = &cobra.Command{
	Use:   "lumens",
	Short: "Native asset (XLM - Lumens) queries",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
}

func init() {
	lumensQueryCmd.AddCommand(getLumensDetailsCmd())
}

func GetLumensCmd() *cobra.Command {
	return lumensQueryCmd
}
