package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bkielbasa/go-ecommerce/backend/productcatalog/adapter"
	"github.com/bkielbasa/go-ecommerce/backend/productcatalog/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ecommerce",
	Short: "A CLI for the ecommerce app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute(db *sql.DB) {
	storage := adapter.NewPostgres(db)
	appServ := app.NewProductService(storage)
	rootCmd.AddCommand(newProductCatalogCmd(appServ))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
