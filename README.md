# go-example

This is simple go example project for learning

Setup development environment for developing go application and bundle it as micro services

### Requirement
    OS : Ubuntu Xenial 16.04 (64 bit)

### Install docker
    scripts/install-docker.sh
    
### Install golang
    https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
   - Download unzip and add go/bin directory in PATH variable
   - Create a directory GOROOT e.g. /root/GOROOT/ 
   - and set environment variable GOPATH e.g. export GOPATH=/root/GOROOT/

### Install gRPC,Protoc,protoc-gen-go
    
	scrips/install-grpc-deps.sh
	 
### References
   - https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/
   - https://golang.org/
