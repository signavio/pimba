package cmd

import (
	"fmt"

	"github.com/signavio/pimba/pkg/push"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push static files to the Pimba server",
	Long:  "Push static files to the Pimba server.",
	Run: func(cmd *cobra.Command, args []string) {
		serverURL := viper.GetString("server")
		bucketName := viper.GetString("name")
		token := viper.GetString("token")

		fmt.Printf("Pushing files to the Pimba server %v...\n", serverURL)
		pushResp, err := push.PushCurrentDirFiles(serverURL, bucketName, token)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Files published successfuly.")
		fmt.Printf("Your push address: http://%v\n", pushResp[0])
		if token == "" {
			fmt.Printf(
				"Use the following token to push to the created bucket: %v\n",
				pushResp[1],
			)
		}
	},
}

func init() {
	RootCmd.AddCommand(pushCmd)

	pushCmd.PersistentFlags().StringP("server", "s", "", "The Pimba API server URL")
	pushCmd.PersistentFlags().StringP("name", "n", "", "Name for the Pimba bucket")
	pushCmd.PersistentFlags().StringP("token", "t", "", "Token for the Pimba bucket")

	viper.BindPFlag("server", pushCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("name", pushCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("token", pushCmd.PersistentFlags().Lookup("token"))
}
