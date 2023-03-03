package root

import (
	"os"

	"github.com/salleh/xlmcmd/cmd/query"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "xmlcmd",
	Short:   "xmlcmd provides command line utility to interact with Stellar blockchain",
	Long:    "xmlcmd provides command line utility to interact with Stellar blockchain based on the given options as parameters.",
	Args:    cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
	Version: "0.1.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Add all child commands
	rootCmd.AddCommand(query.GetQueryCmd())

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
