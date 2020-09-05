package migration

import (
	"log"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/migration"
	"github.com/spf13/cobra"
)

var downCMD = &cobra.Command{
	Use:   "down",
	Short: "Drop tables from database",
	Long:  `Drop tables from database`,
	RunE:  downDatabase,
}

func downDatabase(cmd *cobra.Command, args []string) error {
	log.Println("Dropping database table...")
	db := conn.Default()

	if err := db.DropTableIfExists(migration.Models...).Error; err != nil {
		return err
	}

	log.Println("Database dopped successfully!")

	return nil
}
