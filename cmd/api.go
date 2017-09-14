package cmd

import (
	"fmt"
	"os"

	"github.com/signavio/pimba/pkg/api"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start a new API server",
	Long:  "Start a new API server.",
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetInt("port")
		storagePath := viper.GetString("storage")
		secret := viper.GetString("secret")

		if secret == "" {
			os.Stderr.WriteString("Error: Flag --secret is mandatory\n")
			os.Exit(2)
		}

		fmt.Println("Starting the Pimba server...")
		fmt.Printf("Serving on port %v. ˆC to stop", port)

		api.Serve(port, storagePath, secret)
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)
	apiCmd.PersistentFlags().IntP("port", "p", 8080,
		"The port to serve the API")
	apiCmd.PersistentFlags().StringP("storage", "d", "/tmp/pimba-storage",
		"Path to the directory to store the data received by the API server")
	apiCmd.PersistentFlags().StringP("secret", "k", "",
		"The secret necessary to sign and parse bucket tokens")

	viper.BindPFlag("port", apiCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("storage", apiCmd.PersistentFlags().Lookup("storage"))
	viper.BindPFlag("secret", apiCmd.PersistentFlags().Lookup("secret"))
}
