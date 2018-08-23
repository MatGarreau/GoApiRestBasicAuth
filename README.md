# GoApiRestBasicAuth
API REST for Raspberry Pi using Basic Authent written with Golang. 
This project uses gin-gonic (https://github.com/gin-gonic/gin) and Dave Cheney GPIO library (https://github.com/davecheney/gpio)


# PreRequisite :
* installation of Go
* go get -u github.com/gin-gonic/gin
* go get -u github.com/davecheney/gpio

# Installation of Go

// download : wget https://dl.google.com/go/go1.10.3.linux-armv6l.tar.gz

// extract : sudo tar -C /usr/local/ -xzf go1.10.3.linux-armv6l.tar.gz

// remove archive : rm go1.10.3.linux-armv6l.tar.gz

// update profile : export PATH=$PATH:/usr/local/go/bin # put into ~/.profile

// check version : go version

//set environment variables at the end of .bashrc (or in /etc/profile)

export GOPATH=$HOME/go

export PATH=$PATH:/usr/local/go/bin

// $PATH must look like:

/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/games:/usr/games:/usr/local/go/bin
