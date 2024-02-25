# Introduction

An SQL-liked Database written by Go.

This project is a practice of: 
- Attempt to use Thrift as a RPC for client-server communication.
- Attempt to implement LSM-Tree as a database storage engine.

# Tutorial 
 
## How to compile thrift source into Go language ?

To compile thrift source into Go, from root folder, please run 2 below commands seperately
`thrift -r --gen go ./thrift/query/command/select/service.thrift`
`thrift -r --gen go ./thrift/query/command/insert/service.thrift`

There would be /gen-go folder compiled into root folder