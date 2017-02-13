
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
- AWS EC2


Installation
------------

### homebrew

upcoming.....

### src

upcoming...



Usage
------------

upcoming.....



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