#! /bin/sh
rm *.out *.png dandy
go build dandy.go
./dandy | tee log.out
