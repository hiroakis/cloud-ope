package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type EC2Terminate struct {
	base Base
}

func (self *EC2Terminate) terminateInstances(instanceIds []string, dryRun bool) (bool, error) {
	svc := ec2.New(self.base.Session)

	var instances []*string
	for _, instanceId := range instanceIds {
		instance := aws.String(instanceId)
		instances = append(instances, instance)
	}

	params := &ec2.TerminateInstancesInput{
		InstanceIds: instances,
		DryRun:      aws.Bool(dryRun),
	}

	_, err := svc.TerminateInstances(params)
	if err != nil {
		if dryRun {
			fmt.Println(err)
			fmt.Println(fmt.Sprintf("terminate targets: %v", instanceIds))
			return true, err
		} else {
			return false, err
		}
	}

	return true, nil
}

type AdditionalOpts struct {
	InstanceIds []string `short:"i" long:"instance-ids" default:"" description:"The list of the EC2 instance-id"`
	DryRun      bool     `short:"d" long:"dry-run" default:"false" description:"Dry-run"`
}

func main() {

	ec2Term := &EC2Terminate{}
	additionalOpts := AdditionalOpts{}
	ec2Term.base.setAdditionalOpts(additionalOpts)
	ec2Term.base.initialize()

	instanceIds := ec2Term.base.Opts.AdditionalOpts.InstanceIds
	dryRun := ec2Term.base.Opts.AdditionalOpts.DryRun

	ok, err := ec2Term.terminateInstances(instanceIds, dryRun)
	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("fail!")
		fmt.Println(err)
	}

}
