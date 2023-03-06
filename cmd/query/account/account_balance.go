package account

import (
	"fmt"

	"github.com/salleh/xlmcmd/helper"
	"github.com/spf13/cobra"

	hClient "github.com/stellar/go/clients/horizonclient"
)

var assetCode string

var accountBalanceCmd = &cobra.Command{
	Use:   "balance <Account ID>",
	Short: "get account balance",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		selectedNetwork := cmd.Flags().Lookup("network").Value.String()
		fmt.Printf("Working on network: %v\n", selectedNetwork)
		if !helper.IsValidStellarNetwork(selectedNetwork) {
			return fmt.Errorf("invalid network name")
		}

		var client *hClient.Client

		if selectedNetwork == "public" {
			client = hClient.DefaultPublicNetClient
		} else {
			client = hClient.DefaultTestNetClient
		}

		// Create an account request
		accountRequest := hClient.AccountRequest{AccountID: args[0]}
		// Load the account detail from the network
		account, err := client.AccountDetail(accountRequest)
		if err != nil {
			if horizonErr, ok := err.(*hClient.Error); ok {
				if horizonErr.Problem.Status == 404 {
					return fmt.Errorf("account not found")
				}
			}
			return fmt.Errorf("unable to fetch account balance: %v", err.Error())
		}

		fmt.Printf("Account's asset balances: %s\n", account.AccountID)
		for _, balance := range account.Balances {
			balanceAssetCode := balance.Asset.Code
			if len(balanceAssetCode) == 0 {
				balanceAssetCode = "XLM"
			}

			if len(assetCode) == 0 {
				// asset code not specified, print all balances
				fmt.Printf("  - %s: %s\n", balanceAssetCode, balance.Balance)
			} else {
				if assetCode == balanceAssetCode {
					fmt.Printf("  - %s: %s\n", balanceAssetCode, balance.Balance)
				}

				// just in case user put native as the asset code
				if assetCode == "native" && balanceAssetCode == "XLM" {
					fmt.Printf("  - %s: %s\n", balanceAssetCode, balance.Balance)
				}
			}
		}

		return nil
	},
}

func init() {
	accountBalanceCmd.Flags().StringVarP(
		&assetCode,
		"asset",
		"a",
		"",
		`Asset Code`)
}

func getAccountBalanceCmd() *cobra.Command {
	return accountBalanceCmd
}
