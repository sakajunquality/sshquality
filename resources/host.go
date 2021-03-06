package resources

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
	AddPrefix    bool
	IdentityFile string
}

func InitSshQuality() {
	usr, _ := user.Current()
	ssh_dir := usr.HomeDir + "/.ssh"

	config_dir := ssh_dir + "/conf.d"
	key_dir := ssh_dir + "/keys"

	// @todo create if not exist
	os.MkdirAll(config_dir, 0755)
	os.MkdirAll(key_dir, 0755)

	// @todo add option for this
	exec.Command("cp", "-a", ssh_dir+"/config", ssh_dir+"/config.save").Run()
}

func WriteEachConfig(hosts []Host, config HostConfig, configName string) {
	var config_string string

	for _, host := range hosts {

		if config.AddPrefix {
			config_string += "Host " + configName + "-" + host.Name + "\n"
		} else {
			config_string += "Host " + host.Name + "\n"
		}

		if config.UsePublicIp && host.PublicIpAddress != "" {
			config_string += "  HostName " + host.PublicIpAddress + "\n"
		} else {
			config_string += "  HostName " + host.PrivateIpAddress + "\n"
		}

		config_string += "  User " + config.User + "\n"
		config_string += "  Port " + config.Port + "\n"

		if config.IdentityFile != "" {
			config_string += "  IdentityFile " + config.IdentityFile + "\n"
		}

		config_string += "\n"
	}

	content := []byte(config_string)

	usr, _ := user.Current()

	config_dir := usr.HomeDir + "/.ssh/conf.d"
	file_name := config_dir + "/" + configName + ".conf"

	ioutil.WriteFile(file_name, content, 0644)

	fmt.Printf("Created ssh config to %s\n", file_name)
}

func WriteSshConfig() {
	usr, _ := user.Current()
	ssh_dir := usr.HomeDir + "/.ssh"

	// @todo add option for this
	exec.Command("cp", "-a", ssh_dir+"/config", ssh_dir+"/config.save").Run()

	// @todo refactor
	// exec.Command cannot use Redirect or pipeline
	generate_cmd := "cat " + ssh_dir + "/conf.d/* > " + ssh_dir + "/config"
	exec.Command("sh", "-c", generate_cmd).Run()

	fmt.Printf("Overwrite ~/.ssh/config\n")
}
