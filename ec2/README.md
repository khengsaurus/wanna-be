## Set up EC2

1. Launch an AWS EC2 instance with a key pair
2. Set appropriate connection config - ssh & inbound rules
3. Download private-key (.pem)

<hr/>

### Basic setup

Default user = ec2-user

```bash
# change pw
sudo passwd <user>

# install zsh + oh-my-zsh
# https://blog.devops.dev/installing-zsh-oh-my-zsh-on-amazon-ec2-amazon-linux-2-ami-88b5fc83109
sudo yum -y install zsh
sudo yum install util-linux-user
sudo chsh -s $(which zsh) $(whoami)
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# install homebrew
# https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-homebrew.html#install-homebrew-instructions
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

# add homebrew to path
test -d ~/.linuxbrew && eval $(~/.linuxbrew/bin/brew shellenv)
test -d /home/linuxbrew/.linuxbrew && eval $(/home/linuxbrew/.linuxbrew/bin/brew shellenv)
test -r ~/.bash_profile && echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.bash_profile
echo "eval \$($(brew --prefix)/bin/brew shellenv)" >>~/.profile

# install docker
# https://docs.aws.amazon.com/AmazonECS/latest/developerguide/create-container-image.html
sudo amazon-linux-extras install docker
sudo service docker start
sudo systemctl enable docker
sudo usermod -a -G docker ec2-user

# install nginx
sudo amazon-linux-extras list | grep nginx
sudo amazon-linux-extras enable nginx<1>
sudo yum install nginx
```

<hr/>

### Util

```bash
# ssh function
function _ssh {
  ssh -i <path–to-private-key.pem> <user>@<aws-ec2-public-dns>
}

# scp function - keys, zshrc, config, .env files
function _scp {
  if [ "$#" -ne 2 ]
  then
    echo "Usage: <path to local file> <destination path in remote>" >&2
  else
    scp -i \
      <path–to-private-key.pem> $1 \
      <user>@<aws-ec2-public-dns>:$2
  fi
}

# disk space usage
df -hT /dev/xvda1
```

<hr/>

### Nice-to-have's

```bash
# neovim
brew install neovim

# starship
brew install starship

# zsh autosuggestions
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
```

<hr/>

## Serving multiple apps via nginx

`../init_apps.sh` and `../init_nginx.sh` assume the following folder structure:

```
~
├── git
│ ├── wanna-be
│ │ └── app
│ └── mj-cms
```
