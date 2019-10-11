// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	eos "github.com/eoscanada/eos-go"
	"github.com/spf13/cobra"
)

var (
	proxy4NationAn    = eos.AN("proxy4nation")
	proxy4NationClaim = eos.ActN("claim")
	proxy4NationCmd   = &cobra.Command{
		Use:   "proxy4nation",
		Short: "Proxy4Nation interactions",
	}
)

func init() {
	eos.RegisterAction(proxy4NationAn, proxy4NationClaim, ClaimProxy4Nation{})
	RootCmd.AddCommand(proxy4NationCmd)
}

func NewClaimProxy4Nation(owner eos.AccountName) *eos.Action {
	return &eos.Action{
		Account: proxy4NationAn,
		Name:    proxy4NationClaim,
		Authorization: []eos.PermissionLevel{
			{Actor: owner, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(ClaimProxy4Nation{
			Owner: owner,
		}),
	}
}

type ClaimProxy4Nation struct {
	Owner eos.AccountName
}
