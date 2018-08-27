# GoApiRestBasicAuth
API REST for Raspberry Pi using Basic Authent written with Golang. 
This project uses gin-gonic (https://github.com/gin-gonic/gin) and Dave Cheney GPIO library (https://github.com/davecheney/gpio)

This API has been written for a Raspberry Pi 3 model B v1.2
![RPI3 GPIO](pictures/RPi3-GPIO.jpg)

Available gpios:

{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 16, 17, 18, 19, 20, 21, 22, 23, 24, 26}

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
* use PUT rather than GET to update gpio status
* add URL to get gpio status

## GET API Status (without authentication):
* curl -X GET http://<localhost or Pi@IP>:8088/status
* return HTTP 200 OK response with this message "This server is up. Plz contact sys admin to use it!"

## Get GPIO Status (with authentication):
* curl --user foo:bar -X GET http://<localhost or Pi@IP>:8088/admin/gpiostatus/<gpio_nb>
* return HTTP 200 OK response with this message: "gpio status <gpio_nb> has been called by authenticated user foo. GPIO status is: <true/false>"

### PUT Switch ON del on any GPIO (with authentication):
* curl --user foo:bar -X PUT http://<localhost or Pi@IP>:8088/admin/switchon/<gpio_nb>
* return HTTP 200 OK response with this message: "switchon gpio <gpio_nb> has been called by authenticated user: foo"

### PUT Switch OFF del on any GPIO (with authentication):
* curl --user foo:bar -X PUT http://<localhost or Pi@IP>:8088/admin/switchoff/<gpio_nb>
* return HTTP 200 OK response with this message: "switchoff gpio <gpio_nb> has been called by authenticated user: foo"
