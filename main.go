package main

import (
	"hadoop-raft/raft"
)

func main() {
	r := raft.Server()
	r.Run(":8080")
}
