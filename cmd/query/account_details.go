package query

import (
	"encoding/json"
	"fmt"

	"github.com/salleh/xlmcmd/helper"
	"github.com/spf13/cobra"

	hClient "github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon"
)

var detailsCmd = &cobra.Command{
	Use:   "details <Account ID>",
	Short: "get account details",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		networkFlag := cmd.Parent().PersistentFlags().Lookup("network")
		fmt.Printf("Working on network: %v\n", networkFlag.Value)
		if !helper.IsValidStellarNetwork(networkFlag.Value.String()) {
			return fmt.Errorf("invalid network name")
		}

		var client *hClient.Client

		if networkFlag.Value.String() == "public" {
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

		printAccountDetails(account)

		return nil
	},
}

func init() {
	balanceCmd.Flags().StringVarP(
		&assetCode,
		"asset",
		"a",
		"",
		`Asset ID/Name`)
}

func GetDetailsCmd() *cobra.Command {
	return detailsCmd
}

func printAccountDetails(account horizon.Account) {
	accJSON, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(accJSON))
}
