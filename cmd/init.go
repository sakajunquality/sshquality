package cmd

import (
	"github.com/sakajunquality/sshquality/resources"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize files and directories",
	Long:  "initialize files and directories",
	Run: func(cmd *cobra.Command, args []string) {
		resources.InitSshQuality()
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}
