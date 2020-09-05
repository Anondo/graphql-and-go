package migration

import (
	"log"

	"github.com/Anondo/graphql-and-go/conn"
	"github.com/Anondo/graphql-and-go/database/migration"
	"github.com/spf13/cobra"
)

var upCMD = &cobra.Command{
	Use:   "up",
	Short: "Populate tables in database",
	Long:  `Populate tables in database`,
	RunE:  upDatabase,
}

func upDatabase(cmd *cobra.Command, args []string) error {
	log.Println("Populating database...")
	db := conn.Default()

	if err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(migration.Models...).Error; err != nil {
		return err
	}

	log.Println("Database populated successfully!")
	return nil
}
