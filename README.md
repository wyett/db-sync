## source

### package
- collector
- `common` connect mysql, session, config parser

## collector

support reading data from mysql database, and write it into different middle ware. The source cluster includes
- single mysql instance
- mysql replication with no high availability 
- mysql replication with keepalived+lvs
- pxc cluster with keepalived
- sharding-shpere

charset support
- utf8
- gbk

sync from
- instance with no administrator databases
- databases
- tables


sync mode
- full
- incr
- all

write to
- kafka
- rocketmq
- file
- rabbitmq
- database



## dev schedule

|Function|Description|Version |Status|
| ---- | ---- | ---- | ---- |
|simple|sync whole instance to file|v1.0.0|todo| 
