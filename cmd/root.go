package cmd

import (
	"fmt"
	"log"

	"github.com/Anondo/graphql-and-go/cmd/migration"
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/spf13/cobra"
)

var (
	// RootCMD ...
	RootCMD = &cobra.Command{
		Use:   "graphql-and-go",
		Short: "Play with graphql & Go",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// connect to DB
			if err := conn.Connect(); err != nil {
				return fmt.Errorf("can't connect to DB: %v", err)
			}

			return nil
		},
	}
)

func init() {
	RootCMD.AddCommand(migration.RootCMD)
	RootCMD.AddCommand(serveCMD)
}

// Execute ...
func Execute() {
	if err := RootCMD.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
