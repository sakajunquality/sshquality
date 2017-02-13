package cmd

import (
	"github.com/sakajunquality/sshquality/datasources"
	"github.com/sakajunquality/sshquality/resources"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate configuration",
	Long:  `generate configuration`,
	Run: func(cmd *cobra.Command, args []string) {

		clouds := viper.GetStringMap("clouds")
		for _, c := range clouds {
			cc := c.(map[string]interface{})

			var hosts []resources.Host
			var config resources.HostConfig

			if cc["type"] == "ec2" {
				hosts = datasources.GetEc2Instances()
				config = *datasources.GetEc2DefaultConfig()
			} else {
				continue
			}

			// TODO: order alphabetically

			config_name, _ := cc["name"].(string)

			resources.WriteEachConfig(hosts, config, config_name)
		}

		resources.WriteSshConfig()
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
