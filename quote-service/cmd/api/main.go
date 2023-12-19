package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

const rpcPort = "5001"

type Config struct {
}

func main() {
	log.Println("Start RPC configuration on port: ", rpcPort)

	config := &Config{}

	server := &RPCServer{
		Config: config,
	}
	err := rpc.Register(server)
	if err != nil {
		log.Panic("Error registering RPC server: ", err)
	}

	log.Println("Registered RPC server ", server)
	err = config.rpcListen()
	if err != nil {
		log.Panic("Error starting RPC server: ", err)
	}
}

func (app *Config) rpcListen() error {
	log.Println("Starting RPC server on port ", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		log.Println("Error starting Auth RPC server: ", err)
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}

}
