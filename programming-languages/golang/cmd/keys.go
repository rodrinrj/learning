package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(keysCmd)
}

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Lists the keys saved in the store",
	Run: func(cmd *cobra.Command, args []string) {
		keys, _ := Store.Keys()
		cmd.Println("keys stored: ", keys)
	},
}
