package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Sets a key with a value on the store",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := Store.Set(args[0], []byte(args[1]))
		if err != nil {
			cmd.PrintErrln("an error ocurred")
			return
		}

		cmd.Println("key was set correctly")
	},
}
