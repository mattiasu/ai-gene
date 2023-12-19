package main

import (
	"log"
)

type RPCServer struct {
	Config *Config
}

type RPCPayload struct {
	Promt string
}

func (s *RPCServer) QuoteRPC(payload RPCPayload, reply *string) error {
	log.Println("RPC Quote service called with payload: ", payload.Promt)
	return nil
}
