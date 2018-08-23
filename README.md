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

# Short description

## GET Status (without authentication):
* curl -X GET http://<localhost or Pi@IP>:8088/status
* return HTTP 200 OK response with this message "This server is up. Plz contact sys admin to use it!"

## POST Switch ON del on GPIO N°17 (with authentication):
* curl --user foo:bar -X POST http://<localhost or Pi@IP>:8088/admin/switchon
* return HTTP 200 OK response with this message: "switchon has been called by: foo"

## POST Switch OFF del on GPIO N°17 (with authentication):
* curl --user foo:bar -X POST http://<localhost or Pi@IP>:8088/admin/switchoff
* return HTTP 200 OK response with this message: "switchoff has been called by: foo"
