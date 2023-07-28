package main

import (
	"fmt"
	"strings"

	"gitlab.com/yawning/obfs4.git/transports"

	"github.com/jmwample/nested_modules/mods/server"
)

func main() {
	var s server.Server = &server.BasicServer{}
	out := s.TryParse("AAAAAAAA")
	e := transports.Init()
	if e != nil {
		panic(e)
	}
	strs := transports.Transports()
	fmt.Printf("%+v %+v\n", out, strings.Join(strs, ","))
}
