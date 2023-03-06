package asset

import (
	"encoding/json"
	"fmt"

	"github.com/salleh/xlmcmd/helper"
	"github.com/spf13/cobra"

	hClient "github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon"
)

var assetCode string
var issuerId string

var assetDetailsCmd = &cobra.Command{
	Use:   "details -c <Asset Code> [-i <Issuer ID>]",
	Short: "get asset details",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		selectedNetwork := cmd.Flags().Lookup("network").Value.String()
		fmt.Printf("Working on network: %v\n", selectedNetwork)
		if !helper.IsValidStellarNetwork(selectedNetwork) {
			return fmt.Errorf("invalid network name")
		}

		if len(assetCode) == 0 {
			return fmt.Errorf("asset code must not be blank/empty")
		}

		var client *hClient.Client

		if selectedNetwork == "public" {
			client = hClient.DefaultPublicNetClient
		} else {
			client = hClient.DefaultTestNetClient
		}

		// Create an asset request
		assetRequest := hClient.AssetRequest{
			ForAssetCode:   assetCode,
			ForAssetIssuer: issuerId,
		}
		// Load the asset detail from the network
		assetsPage, err := client.Assets(assetRequest)
		if err != nil {
			if horizonErr, ok := err.(*hClient.Error); ok {
				if horizonErr.Problem.Status == 404 {
					return fmt.Errorf("asset not found")
				}
			}
			return fmt.Errorf("unable to fetch asset details: %v", err.Error())
		}

		printAssetDetails(assetsPage)

		return nil
	},
}

func init() {
	assetDetailsCmd.Flags().StringVarP(
		&assetCode,
		"code",
		"c",
		"",
		`Asset code to query. Use 'get lumens' command to query on the native XLM information`)
	assetDetailsCmd.Flags().StringVarP(
		&issuerId,
		"issuer",
		"i",
		"",
		`Asset issuer ID to query`)

	assetDetailsCmd.MarkFlagRequired("code")
}

func getAssetDetailsCmd() *cobra.Command {
	return assetDetailsCmd
}

func printAssetDetails(account horizon.AssetsPage) {
	apJSON, err := json.MarshalIndent(account, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(apJSON))
}
