package raft

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var serverCfg *config

// Hello hello world
func Hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "hello world!",
	})
}

// StartNodes 初始化网络以及节点
func StartNodes(c *gin.Context) {
	if serverCfg != nil {
		c.JSON(200, gin.H{
			"msg": "already started",
		})
	}
	s := c.Query("servers")
	servers, _ := strconv.ParseInt(s, 10, 64)
	serverCfg = make_config(nil, int(servers), false)
	c.JSON(200, gin.H{
		"msg": "success!",
	})
}

// CleanNodes 删除所有节点并释放网络资源
func CleanNodes(c *gin.Context) {
	// serverCfg.end()
	serverCfg.cleanup()
	serverCfg = nil
	c.JSON(200, gin.H{
		"msg": "success!",
	})
}

// DisconnectNode 断开编号为number的节点
func DisconnectNode(c *gin.Context) {
	s := c.Query("number")
	number := 0
	fmt.Sscanf(s, "%d", &number)
	serverCfg.disconnect(number)
	c.JSON(200, gin.H{
		"msg": "success!",
	})
}

// ReconnectNode 重连编号为number的节点
func ReconnectNode(c *gin.Context) {
	s := c.Query("number")
	number := 0
	fmt.Sscanf(s, "%d", &number)
	serverCfg.connect(number)
	c.JSON(200, gin.H{
		"msg": "success!",
	})
}

// GetState 获取编号为number的节点状态
func GetState(c *gin.Context) {
	s := c.Query("number")
	number := 0
	fmt.Sscanf(s, "%d", &number)
	serverCfg.rafts[number].mu.Lock()
	defer serverCfg.rafts[number].mu.Unlock()
	// term, leader := serverCfg.rafts[number].GetState()

	c.JSON(200, gin.H{
		"number":      number,
		"term":        serverCfg.rafts[number].CurrentTerm,
		"votedFor":    serverCfg.rafts[number].VotedFor,
		"state":       serverCfg.rafts[number].state,
		"votedCount":  serverCfg.rafts[number].votedCount,
		"leaderId":    serverCfg.rafts[number].leaderId,
		"logs":        serverCfg.rafts[number].Log,
		"commitIndex": serverCfg.rafts[number].commitIndex,
		"lastApplied": serverCfg.rafts[number].lastApplied,
	})
}

// StartCommand 向某一节点发送command请求
func StartCommand(c *gin.Context) {
	command := c.Query("command")
	s := c.Query("number")
	number := 0
	fmt.Sscanf(s, "%d", &number)
	cmd, _ := strconv.ParseInt(command, 10, 64)
	index, term, isLeader := serverCfg.rafts[number].Start(int(cmd))
	c.JSON(200, gin.H{
		"index":    index,
		"term":     term,
		"isLeader": isLeader,
	})
}

// Server 创建Server
func Server() *gin.Engine {
	serverCfg = nil
	r := gin.Default()
	r.GET("/api/startnodes", StartNodes)
	r.GET("/api/cleannodes", CleanNodes)
	r.GET("/api/disconnect", DisconnectNode)
	r.GET("/api/reconnect", ReconnectNode)
	r.GET("/api/getstate", GetState)
	r.GET("/api/startcommand", StartCommand)
	r.Static("/index", "./frontend")
	return r
}
