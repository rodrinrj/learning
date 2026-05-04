package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(delCmd)
}

var delCmd = &cobra.Command{
	Use:   "del [key]",
	Short: "Deletes the key from the store",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := Store.Delete(args[0])
		if err != nil {
			cmd.PrintErrln("an error ocurred")
			return
		}

		cmd.Println("key was deleted successfully")
	},
}
