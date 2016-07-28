package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

type ElbRegister struct {
	base Base
}

func (self *ElbRegister) registerInstances(name string, instanceIds []string) (bool, error) {
	svc := elb.New(self.base.Session)

	var elbInstances []*elb.Instance
	for _, instanceId := range instanceIds {
		elbInstance := &elb.Instance{
			InstanceId: aws.String(instanceId),
		}
		elbInstances = append(elbInstances, elbInstance)
	}

	params := &elb.RegisterInstancesWithLoadBalancerInput{
		Instances:        elbInstances,
		LoadBalancerName: aws.String(name),
	}

	_, err := svc.RegisterInstancesWithLoadBalancer(params)

	if err != nil {
		return false, err
	}

	return true, nil
}

type AdditionalOpts struct {
	Name        string   `short:"n" long:"name" default:"" description:"The name of the ELB"`
	InstanceIds []string `short:"i" long:"instance-ids" default:"" description:"The list of the EC2 instance-id"`
}

func main() {

	elbReg := &ElbRegister{}
	additionalOpts := AdditionalOpts{}
	elbReg.base.setAdditionalOpts(additionalOpts)
	elbReg.base.initialize()

	name := elbReg.base.Opts.AdditionalOpts.Name
	instanceIds := elbReg.base.Opts.AdditionalOpts.InstanceIds

	ok, err := elbReg.registerInstances(name, instanceIds)
	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("fail!")
		fmt.Println(err)
	}

}
