package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jessevdk/go-flags"
)

type Base struct {
	Session   client.ConfigProvider
	AwsConfig AwsConfig
	Opts      Opts
}

type (
	AwsConfig struct {
		AwsAccessKeyId     string `json:"aws_access_key_id"`
		AwsSecretAccessKey string `json:"aws_secret_access_key"`
		AwsRegion          string `json:"aws_region"`
	}
	Opts struct {
		ConfigFilePath string `short:"c" long:"config" default:"config.json" description:"The config file path"`
		Output         string `short:"o" long:"output" default:"" description:"optput format"`
		AdditionalOpts AdditionalOpts
	}
)

func (base *Base) parseOptions(args []string) {
	_, err := flags.ParseArgs(&base.Opts, args)
	if err != nil {
		log.Fatalln(err)
	}
}

func (base *Base) setAdditionalOpts(additionalOpts AdditionalOpts) {
	base.Opts.AdditionalOpts = additionalOpts
}

func (base *Base) initialize() {
	base.parseOptions(os.Args[1:])
	base.loadConfig(base.Opts.ConfigFilePath)
	base.setSession(base.AwsConfig.AwsAccessKeyId, base.AwsConfig.AwsSecretAccessKey, base.AwsConfig.AwsRegion)
}

func (base *Base) loadConfig(confPath string) {
	var file []byte
	var err error
	var awsConfig AwsConfig

	file, err = ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(file, &awsConfig)
	if err != nil {
		log.Fatalln(err)
	}
	base.AwsConfig = awsConfig
}

func (base *Base) setSession(awsAccessKeyId, awsSecretAccessKey, awsRegion string) {
	creds := credentials.NewStaticCredentials(awsAccessKeyId, awsSecretAccessKey, "")
	sess := session.New(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
	})
	base.Session = sess
}
