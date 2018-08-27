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

# Short description of v_0_0_3
This version include following improvments :
* remove unused functions
* move gpio func into business
* use PUT rather than GET to update pin status
* add URL to get gpio status

## GET API Status (without authentication):
* curl -X GET http://<localhost or Pi@IP>:8088/status
* return HTTP 200 OK response with this message "This server is up. Plz contact sys admin to use it!"

## Get GPIO Status:
* curl --user foo:bar -X GET http://<localhost or Pi@IP>:8088/admin/pinstatus/<gpio_nb>
* return HTTP 200 OK response with this message: "pin status <gpio_nb> has been called by authenticated user foo. Pin status is: <true/false>"

### PUT Switch ON del on any GPIO (with authentication):
* curl --user foo:bar -X PUT http://<localhost or Pi@IP>:8088/admin/switchon/<gpio_nb>
* return HTTP 200 OK response with this message: "switchon pin <gpio_nb> has been called by authenticated user: foo"

### PUT Switch OFF del on any GPIO (with authentication):
* curl --user foo:bar -X PUT http://<localhost or Pi@IP>:8088/admin/switchoff/<gpio_nb>
* return HTTP 200 OK response with this message: "switchoff pin <gpio_nb> has been called by authenticated user: foo"
