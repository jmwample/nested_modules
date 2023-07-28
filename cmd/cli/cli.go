package main

import nm "github.com/jmwample/nested_mods"

func main() {
	_, err := nm.Dial()
	if err != nil {
		panic(err)
	}
}
