# Introduction

An SQL-liked Database written by Go.

This project is a practice of: 
- Attempt to use Thrift as a RPC for client-server communication.
- Attempt to implement LSM-Tree as a database storage engine.

# Tutorial 
 
## How to compile thrift source into Go language ?

To compile thrift source into Go, from root folder, please run the below command\
`thrift -r -v --gen go --out vendor thrift/main.thrift`

There would be a /vendor/thrift/ folder compiled into root folder

# How to test

To run all the test inside **go_client** module, please run the below command\
`go test ./gosql_client/... -count=1`
