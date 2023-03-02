package rootcmd

import (
	"os"

	"github.com/spf13/cobra"
)

var StellarNetwork string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xmlcmd [-n public|testnet]",
	Short: "xmlcmd provides command line utility to query Stellar blockchain",
	Long:  "xmlcmd provides command line utility to query Stellar blockchain based on the given options as parameters.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Args:    cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
	Version: "1.1.3",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(
		&StellarNetwork,
		"network",
		"n",
		"public",
		`Stellar network to interact with: public | testnet`)
}
