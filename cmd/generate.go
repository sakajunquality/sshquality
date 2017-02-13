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
				// default
				region := datasources.EC2DefaultRegion
				credential := datasources.EC2DefaultCredential

				if cc["region"] != nil {
					region, _ = cc["region"].(string)
				}
				if cc["credential"] != nil {
					credential, _ = cc["credential"].(string)
				}

				hosts = datasources.GetEc2Instances(credential, region)
				config = *datasources.GetEc2DefaultConfig()
			} else {
				continue
			}

			// TODO: order alphabetically

			config_name, _ := cc["name"].(string)
			default_user, _ := cc["default_user"].(string)
			default_key, _ := cc["default_key"].(string)
			add_prefix, _ := cc["add_prefix"].(bool)
			use_public_ip, _ := cc["use_public_ip"].(bool)

			if use_public_ip {
				config.UsePublicIp = true
			}

			if add_prefix {
				config.AddPrefix = true
			}

			if add_prefix {
				config.AddPrefix = true
			}

			if default_user != "" {
				config.User = default_user
			}

			if default_key != "" {
				config.IdentityFile = default_key
			}

			resources.WriteEachConfig(hosts, config, config_name)
		}

		resources.WriteSshConfig()
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
