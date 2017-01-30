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

package resources

import (
	"fmt"
	"os"
	"os/user"
	"io/ioutil"
)

type Host struct {
	Name string
	PrivateIpAddress string
	PublicIpAddress string
}

type HostConfig struct {
	User string
	UsePublicIp bool
	Port string
	IdentityFile string
}

func WriteSshConfig(hosts []Host, config HostConfig, config_name string) {
	var config_string string

	for _, host := range hosts {
		config_string += "Host " + host.Name + "\n"
		config_string += "  HostName " + host.PrivateIpAddress + "\n"
		config_string += "  User "+ config.User +"\n"
		config_string += "  Port "+config.Port+"\n\n"
	}

	content := []byte(config_string)

	usr, _ := user.Current()

	config_dir := usr.HomeDir + "/.ssh/conf.d" 
    file_name := config_dir +"/"+ config_name+".conf"

	os.MkdirAll(config_dir, 0755)
	ioutil.WriteFile(file_name, content, 0644)

	fmt.Printf("created ssh config to %s\n", file_name)
}