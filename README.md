
sshquality
=================================


Overview
------------
SSH config manager for cloud, like AWS. This app will create individual ssh config (you can add manual config too), and combine them into one file `~/.ssh/config`.

### Naive Idea

Organize ssh configs, generated both manually and automatically, and keep `~/.ssh` directory clean.

#### Example

```
~/.ssh
├── conf.d
│   ├── 000-manual.conf // manual
│   ├── aws.conf // auto
│   ├── aws-2.conf // auto
│   └── gcp.conf // auto
├── config // all the configs in conf.d are combined
└── keys
    ├── id_rsa
    └── id_rsa.pub
```

#### SupportedPlatform
- AWS EC2
- GCP (Upcoming)
- Azure (Upcoming)




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

First you need to create config file for the app. The file must be located at your home directory and named ".sshquality.yaml".

#### Config Example

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

#### config details

| name          | accepted value                     | default        | mandatory |
| ------------- | ---------------------------------- | -------------- | --------- |
| type          | `ec2`                              | -              | Yes       |
| name          | any string                         | -              | Yes       |
| credential    | any string (See ~/.aws/credential) | default        | No        |
| region        | AWS region                         | ap-northeast-1 | No        |
| add_prefix    | bool                               | false          | No        |
| use_public_ip | bool                               | false          | No        |



### Init

By the command `init`, the app will create necessary directories and files.

```bash
$ sshquality init
```



### Manual Configuretion

After initilization, you can add some manual configs. 

#### Example

```
// ~/.ssh/conf.d/000-manual.conf

Host *
  forwardagent yes
  
Host my-personal-server
   hostname test.domain.local
   user hello
```



### Generate

Finally generate the `~/.ssh/config`

**Note:** This will override your `~/.ssh/config`.

```bash
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