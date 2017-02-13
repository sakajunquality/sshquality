package resources

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
)

type Host struct {
	Name             string
	PrivateIpAddress string
	PublicIpAddress  string
}

type HostConfig struct {
	User         string
	UsePublicIp  bool
	Port         string
	IdentityFile string
}

func WriteSshConfig(hosts []Host, config HostConfig, config_name string) {
	var config_string string

	for _, host := range hosts {
		config_string += "Host " + host.Name + "\n"
		config_string += "  HostName " + host.PrivateIpAddress + "\n"
		config_string += "  User " + config.User + "\n"
		config_string += "  Port " + config.Port + "\n\n"
	}

	content := []byte(config_string)

	usr, _ := user.Current()

	config_dir := usr.HomeDir + "/.ssh/conf.d"
	file_name := config_dir + "/" + config_name + ".conf"

	os.MkdirAll(config_dir, 0755)
	ioutil.WriteFile(file_name, content, 0644)

	fmt.Printf("created ssh config to %s\n", file_name)
}
