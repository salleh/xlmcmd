package account

import (
	"github.com/spf13/cobra"
)

var accountQueryCmd = &cobra.Command{
	Use:   "account",
	Short: "account queries",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
}

func init() {
	accountQueryCmd.AddCommand(getAccountBalanceCmd())
	accountQueryCmd.AddCommand(getAccountDetailsCmd())
}

func GetAccountCmd() *cobra.Command {
	return accountQueryCmd
}
