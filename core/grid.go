package core

import (
	"fmt"
	"sync"
)

// Grid AOI地图格子类型/*
type Grid struct {
	//格子ID
	Gid int
	//格子左上角边界坐标
	MinX int
	//格子右上角边界坐标
	MaxX int
	//格子左下角边界坐标
	MinY int
	//格子右下角边界坐标
	MaxY int
	//玩家或者物体的集合
	PlayerIDs map[int]bool

	PLock sync.RWMutex
}

// GridInit 初始化格子/创建格子
func GridInit(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		Gid:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		PlayerIDs: make(map[int]bool),
	}
}

// AddPlayer 添加玩家
func (g *Grid) AddPlayer(playerID int) {
	g.PLock.Lock()
	defer g.PLock.Unlock()
	g.PlayerIDs[playerID] = true
}

// RemovePlayer 移除玩家
func (g *Grid) RemovePlayer(playerID int) {
	g.PLock.Lock()
	defer g.PLock.Unlock()
	delete(g.PlayerIDs, playerID)
}

// GetAllPlayer 获取当前格子所有玩家
func (g *Grid) GetAllPlayer() (playerIDS []int) {
	g.PLock.RLock()
	defer g.PLock.RUnlock()
	for k, _ := range g.PlayerIDs {
		playerIDS = append(playerIDS, k)
	}
	return
}

// String 调试输出Grid相关数据
func (g *Grid) String() string {
	return fmt.Sprintf("Grid:%d,minx:%d,maxx:%d,miny:%d,maxy:%d，playerIDS:%v", g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY, g.PlayerIDs)
}
