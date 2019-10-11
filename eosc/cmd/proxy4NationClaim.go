// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import "github.com/spf13/cobra"

// proxy4NationClaimCmd represents the `proxy4nation::claim` command
var proxy4NationClaimCmd = &cobra.Command{
	Use:   "claim [account]",
	Short: "Claims all rewards.",
	Long:  "Claims all rewards from EOS4Nation for using their proxy",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		account := toAccount(args[0], "account")

		pushEOSCActions(getAPI(), NewClaimProxy4Nation(
			account,
		))
	},
}

func init() {
	proxy4NationCmd.AddCommand(proxy4NationClaimCmd)
}
