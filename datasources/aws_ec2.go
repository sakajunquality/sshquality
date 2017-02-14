package datasources

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/ec2"

	"github.com/sakajunquality/sshquality/resources"
)

const EC2DefaultRegion = "ap-northeast-1"
const EC2DefaultCredential = "default"
const EC2DefaultUser = "ec2-user"

func GetEc2Instances(awsCredential string, awsRegion string) []resources.Host {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Printf("failed to create session %v\n", err)
	}

	svc := ec2.New(sess, &aws.Config{
		Credentials: credentials.NewSharedCredentials("", awsCredential),
		Region:      aws.String(awsRegion),
	})

	res, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("there was an error listing instances in", awsRegion, err.Error())
	}

	var hosts []resources.Host

	for _, r := range res.Reservations {
		for _, i := range r.Instances {
			var new_host resources.Host
			var tag_name string
			var private_ip_address string
			var public_ip_address string

			for _, t := range i.Tags {
				if *t.Key == "Name" {
					tag_name = *t.Value
				}
			}

			// ignore terminated instances
			if i.PrivateIpAddress == nil {
				continue
			}
			private_ip_address = *i.PrivateIpAddress

			if i.PublicIpAddress != nil {
				public_ip_address = *i.PublicIpAddress
			}

			new_host = resources.Host{Name: tag_name, PrivateIpAddress: private_ip_address, PublicIpAddress: public_ip_address}
			hosts = append(hosts, new_host)
		}
	}

	return hosts
}

func GetEc2DefaultConfig() *resources.HostConfig {
	return &resources.HostConfig{User: EC2DefaultUser, UsePublicIp: false, Port: "22"}
}
