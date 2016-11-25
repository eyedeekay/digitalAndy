#! /bin/bash
rm *.out output/*.png dandy
go build dandy.go
for ((x = 0 ; x <= 30 ; x++)); do
  for files in $(find skel -name *.txt); do
    echo "path=skel/colors.txt" > config.txg
    echo "path=$files" >> config.txg
    echo $x
    ./dandy &> /dev/null
  done
done