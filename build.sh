#! /bin/sh
rm *.log *.out *.png
go build lair-image-gen.go 2> err.log
#./a.out 1> test.log 2> err.run.log
