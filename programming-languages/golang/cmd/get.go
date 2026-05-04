package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Shows the current value of a key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		value, err := Store.Get(args[0])
		if err != nil {
			cmd.PrintErrln("an error ocurred")
			return
		}

		cmd.Println(string(value))
	},
}
