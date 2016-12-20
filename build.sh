#! /bin/bash
rm *.out output/*.png dandy
rm skel/*.txt.folder/ -rf
go build dandy.go main.go
