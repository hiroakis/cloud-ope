package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/elb"
)

type ELB struct {
	LoadBalancerName string
	Scheme           string
}

type ElbList struct {
	base Base
}

func (self *ElbList) getLoadBalancers() []ELB {
	svc := elb.New(self.base.Session)
	params := &elb.DescribeLoadBalancersInput{
		LoadBalancerNames: []*string{},
	}
	resp, err := svc.DescribeLoadBalancers(params)

	if err != nil {
		log.Fatalln(err)
	}

	elbs := []ELB{}
	for _, lbd := range resp.LoadBalancerDescriptions {
		elb := ELB{}
		elb.LoadBalancerName = *lbd.LoadBalancerName
		elb.Scheme = *lbd.Scheme
		elbs = append(elbs, elb)
	}
	return elbs
}

type AdditionalOpts struct{}

func main() {

	elbList := &ElbList{}
	elbList.base.initialize()

	elbs := elbList.getLoadBalancers()

	switch elbList.base.Opts.Output {
	case "json":
		j, err := json.Marshal(elbs)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(j))

	default:
		header := fmt.Sprintf("LoadBalancerName,Scheme")
		fmt.Println(header)
		var line string
		for _, v := range elbs {
			line = fmt.Sprintf(`"%s","%s"`, v.LoadBalancerName, v.Scheme)
			fmt.Println(line)
		}
	}
}
