package migration

import (
	"fmt"

	_ "github.com/Anondo/graphql-and-go/config" // for config init
	"github.com/Anondo/graphql-and-go/conn"
	"github.com/spf13/cobra"
)

var (
	// RootCMD ...
	RootCMD = &cobra.Command{
		Use:   "migration",
		Short: "Runs migration on the database",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

			if err := conn.Connect(); err != nil {
				return fmt.Errorf("Cant't connect database: %v", err)
			}
			return nil
		},
	}
)

func init() {
	RootCMD.AddCommand(upCMD)
	RootCMD.AddCommand(downCMD)
}
