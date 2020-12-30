#!/bin/bash
#source /home/jakekausler/go/src/github.com/jakekausler/Library-Organizer-2.0/dev/libraries.env;go run /home/jakekausler/go/src/github.com/jakekausler/Library-Organizer-2.0/server/main/server.go -username jakekausler -password Jake021f2f1! -mysqlport 3306 -host localhost -appport 8181
PATH=$PATH:/usr/local/go/bin/go
GOPATH=/home/jakekausler/go
source /home/jakekausler/go/src/github.com/jakekausler/Library-Organizer-2.0/dev/libraries.env
/home/jakekausler/go/src/github.com/jakekausler/Library-Organizer-2.0/server/main/server 2>&1 | tee -a /home/jakekausler/test 
