// Copyright © 2017 sakajunquality
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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

			resources.WriteSshConfig(hosts, config, config_name)
		}	
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}
