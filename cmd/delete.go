// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a code snippet.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`..`)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
