#! /bin/bash
rm *.out *.png dandy
go build dandy.go
./dandy | tee log.out
for ((x = 0 ; x <= 100 ; x++)); do
  ./dandy
done