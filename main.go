package main

import "flag"

func main(){
	opts := flag.String("conf","config.txg","path to a configuration file")
	confErr, configArray := readConfig(*opts)
	var colors [][]string
	loopConfigs(confErr, configArray, &colors)
}
