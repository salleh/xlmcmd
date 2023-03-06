package lumens

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type LumenSupply struct {
	UpdatedAt         time.Time `json:"updatedAt"`
	OriginalSupply    string    `json:"originalSupply"`
	InflationLumens   string    `json:"inflationLumens"`
	BurnedLumens      string    `json:"burnedLumens"`
	TotalSupply       string    `json:"totalSupply"`
	UpgradeReserve    string    `json:"upgradeReserve"`
	FeePool           string    `json:"feePool"`
	SdfMandate        string    `json:"sdfMandate"`
	CirculatingSupply string    `json:"circulatingSupply"`
	Details           string    `json:"_details"`
}

var lumensDetailsCmd = &cobra.Command{
	Use:   "details",
	Short: "get asset details",
	Args:  cobra.MatchAll(cobra.NoArgs, cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {

		resp, err := http.Get("https://dashboard.stellar.org/api/v2/lumens")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var supply LumenSupply
		err = json.NewDecoder(resp.Body).Decode(&supply)
		if err != nil {
			panic(err)
		}

		printLumensDetails(supply)

		return nil
	},
}

func getLumensDetailsCmd() *cobra.Command {
	return lumensDetailsCmd
}

func printLumensDetails(lumenDetail LumenSupply) {
	ldJSON, err := json.MarshalIndent(lumenDetail, "", "  ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(ldJSON))
}
