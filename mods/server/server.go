package server

import (
	"encoding/hex"
	"fmt"

	"gitlab.com/yawning/obfs4.git/transports/obfs4"
)

type Server interface {
	TryParse(string) []byte
}

type BasicServer struct{}

func (b *BasicServer) TryParse(s string) []byte {
	out, _ := hex.DecodeString(s)
	t := obfs4.Transport{}
	fmt.Printf("%+v", t)
	return out
}
