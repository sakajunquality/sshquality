
sshquality
=================================


Overview
------------
SSH config manager for cloud. This CLI app will create individual ssh config (you can add manual config too), and combine them into one file `~/.ssh/config`.



### Naive Idea

Organize and keep `~/.ssh` directory clean.

```
~/.ssh
├── conf.d
│   ├── 000-manual.conf
│   ├── aws.conf
│   ├── aws-2.conf
│   └── gcp.conf
├── config
└── keys
    ├── id_rsa
    └── id_rsa.pub
```



#### SupportedPlatform
- AWS EC2 (single account and single region only....)




Multiple accounts and regions will be implemented



Installation
------------

### src

```bash
$ go isntall github.com/sakajunquality/sshquality
```



### homebrew

upcoming.....



Usage
------------

### config file (~/.sshquality.yaml)

```yaml
---
clouds:
  1:
    type: ec2
    name: aws
    credential: default
    region: ap-northeast-1
    add_prefix: false
  2:
    type: ec2
    name: another-aws
    credential: cred2
    region: us-east-1
    add_prefix: true
    use_public_ip: true
```

#### config

| name          | accepted value                     | default        | mandatory |
| ------------- | ---------------------------------- | -------------- | --------- |
| type          | `ec2`                              | -              | Yes       |
| name          | any string                         | -              | Yes       |
| credential    | any string (See ~/.aws/credential) | default        | No        |
| region        | AWS region                         | ap-northeast-1 | No        |
| add_prefix    | bool                               | false          | No        |
| use_public_ip | bool                               | false          | No        |



### init & generate .ssh/config

```bash
$ sshquality init
$ sshquality generate
```



## Eco-System

### fzf (+zsh)

```
sssh() { ## whatever the name is
  local servers server
  servers=$(grep -iE "^host[[:space:]]+[^*]" ~/.ssh/config | awk "{print \$2}") &&
  server=$(echo "$servers" |
           fzf-tmux -d $(( 2 + $(wc -l <<< "$servers") )) +m) &&
  echo "ssh to $server ...\n" && ssh $(echo "$server")
}
```

### peco (+zsh)

upcoming...




Contributing
------------
Always welcome for your contribution




License & Authors
------------
- Author:: @sakajunquality
- License:: MIT