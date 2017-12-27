#!/bin/bash

############################################################################
#
#						README
#
# @Purpose: Utitlity script to setup and install gRPC dependancies
# @Requirement : Ubuntu Xenial 16.04 (64 bit)	
# @Auther: Amit Wankhede
############################################################################

apt-get install -y unzip 

cd /

wget https://github.com/google/protobuf/releases/download/v3.5.1/protoc-3.5.1-linux-x86_64.zip  

unzip protoc-3.5.1-linux-x86_64.zip

go get -u github.com/golang/protobuf/{proto,protoc-gen-go} 

go get -u google.golang.org/grpc