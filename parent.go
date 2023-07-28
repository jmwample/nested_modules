package nestedmodules

import (
	"fmt"
	"net"
	"github.com/jmwample/nested_mods/mods/client"
)

func check() bool {

	c := client.New()

	err := c.Function()
	if err != nil {
		return true
	}

	return false
}

func Dial() (net.Conn, error) {
	if !check() {
		return nil, fmt.Errorf("whatever")
	}
	return nil, fmt.Errorf("no winning")
}
