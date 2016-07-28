# cloud-ope

My AWS Operation tools for work. 自分がよく使うオペレーションをコマンド化したもの。

## Installation インストール

```
git clone git@github.com:hiroakis/cloud-ope.git
cd cloud-ope

# For OSX
make
sudo make install
 -> binaries move to /usr/local/bin/

# For Linux
export GOOS=linux
make
sudo make install
 -> binaries move to /usr/local/bin/
```

## Usage 使い方

### Common options

```
-c The config file path(default: ./config.json)
```

### Configuration file

```
{
    "aws_access_key_id": "xxxxxxx",
    "aws_secret_access_key": "xxxxxx",
    "aws_region": "ap-northeast-1"
}
```

### ec2-list

List EC2 instances.

```
ec2-list
```

### ec2-terminate

Terminate EC2 instances

```
ec2-terminate -i [instanceid] -i [instanceid] -i ...
```

### elb-list

List ELBs

```
elb-list
```

### elb-describe

Show instances which are attached to the ELB.

```
elb-describe -n [The-Name-Of-ELB]
```

### elb-dereg

Detach specified instances to the ELB.

```
elb-dereg -n [The-Name-Of-ELB] -i [instanceid] -i [instanceid] -i ...
```

### elb-reg

Attach specified instances to the ELB.

```
elb-reg -n [The-Name-Of-ELB] -i [instanceid] -i [instanceid] -i ...
```

## License

MIT