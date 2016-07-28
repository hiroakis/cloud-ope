package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2 struct {
	ID          string
	PrivateIP   string
	PublicIP    string
	Service     string
	Role        string
	Name        string
	Environment string
	State       string
}

type EC2List struct {
	base Base
}

func (self *EC2List) getInstances() []EC2 {
	cli := ec2.New(self.base.Session)
	resp, err := cli.DescribeInstances(nil)
	if err != nil {
		log.Fatalln(err)
	}

	instances := []EC2{}
	for _, v := range resp.Reservations {
		instance := EC2{}

		for _, inst := range v.Instances {
			instance.ID = *inst.InstanceId

			if inst.PrivateIpAddress != nil {
				instance.PrivateIP = *inst.PrivateIpAddress
			}
			if inst.PublicIpAddress != nil {
				instance.PublicIP = *inst.PublicIpAddress
			}

			for _, tag := range inst.Tags {
				switch *tag.Key {
				case "Service":
					instance.Service = *tag.Value
				case "Role":
					instance.Role = *tag.Value
				case "Name":
					instance.Name = *tag.Value
				case "Environment":
					instance.Environment = *tag.Value
				}
			}
			instance.State = *inst.State.Name
			instances = append(instances, instance)
		}
	}
	return instances
}

type AdditionalOpts struct{}

func main() {

	ec2List := EC2List{}
	ec2List.base.initialize()
	instances := ec2List.getInstances()

	switch ec2List.base.Opts.Output {
	case "json":
		j, err := json.Marshal(instances)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(j))

	default:
		header := fmt.Sprintf("InstanceId,PrivateIP,PublicIP,Service,Role,Name,Environment")
		fmt.Println(header)
		var line string
		for _, v := range instances {
			line = fmt.Sprintf(`"%s","%s","%s","%s","%s","%s","%s","%s"`, v.ID, v.PrivateIP, v.PublicIP, v.Service, v.Role, v.Name, v.Environment, v.State)
			fmt.Println(line)
		}
	}
}
