package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

type ElbDeregister struct {
	base Base
}

func (self *ElbDeregister) deregisterInstances(name string, instanceIds []string) (bool, error) {
	svc := elb.New(self.base.Session)

	var elbInstances []*elb.Instance
	for _, instanceId := range instanceIds {
		elbInstance := &elb.Instance{
			InstanceId: aws.String(instanceId),
		}
		elbInstances = append(elbInstances, elbInstance)
	}

	params := &elb.DeregisterInstancesFromLoadBalancerInput{
		Instances:        elbInstances,
		LoadBalancerName: aws.String(name),
	}

	_, err := svc.DeregisterInstancesFromLoadBalancer(params)

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

	elbDereg := &ElbDeregister{}
	additionalOpts := AdditionalOpts{}
	elbDereg.base.setAdditionalOpts(additionalOpts)
	elbDereg.base.initialize()

	name := elbDereg.base.Opts.AdditionalOpts.Name
	instanceIds := elbDereg.base.Opts.AdditionalOpts.InstanceIds

	ok, err := elbDereg.deregisterInstances(name, instanceIds)
	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("fail!")
		fmt.Println(err)
	}

}
