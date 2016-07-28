package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elb"
)

type (
	InstanceState struct {
		InstanceId  string
		State       string
		Description string
		ReasonCode  string
	}
)

type ElbDescribe struct {
	base Base
}

type AdditionalOpts struct {
	Name string `short:"n" long:"name" default:"" description:"The name of the ELB"`
}

func (self *ElbDescribe) describeLoadbalancer(name string) []InstanceState {
	svc := elb.New(self.base.Session)
	params := &elb.DescribeInstanceHealthInput{
		LoadBalancerName: aws.String(name),
		Instances: []*elb.Instance{
			{},
		},
	}

	resp, err := svc.DescribeInstanceHealth(params)
	if err != nil {
		log.Fatalln(err)
	}

	instanceStates := []InstanceState{}
	for _, is := range resp.InstanceStates {
		instanceState := InstanceState{}
		instanceState.InstanceId = *is.InstanceId
		instanceState.State = *is.State
		instanceState.Description = *is.Description
		instanceState.ReasonCode = *is.ReasonCode
		instanceStates = append(instanceStates, instanceState)
	}

	return instanceStates
}

func main() {

	elbd := &ElbDescribe{}
	additionalOpts := AdditionalOpts{}
	elbd.base.setAdditionalOpts(additionalOpts)
	elbd.base.initialize()

	name := elbd.base.Opts.AdditionalOpts.Name
	dl := elbd.describeLoadbalancer(name)

	switch elbd.base.Opts.Output {
	case "json":
		j, err := json.Marshal(dl)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(j))

	default:
		header := fmt.Sprintf("InstanceId,State,Description,ReasonCode")
		fmt.Println(header)
		var line string
		for _, v := range dl {
			line = fmt.Sprintf(`"%s","%s","%s","%s"`, v.InstanceId, v.State, v.Description, v.ReasonCode)
			fmt.Println(line)
		}
	}
}
