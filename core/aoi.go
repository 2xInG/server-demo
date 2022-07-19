package core

import "fmt"

// AOI AOI区域
type AOI struct {
	//最小X轴
	MinX int
	//最大X轴
	MaxX int
	//最小Y轴
	MinY int
	//最大Y轴
	MaxY int
	//X轴格子数量
	CountX int
	//Y轴格子数量
	CountY int
	//总的格子集合
	grids map[int]*Grid
}

// AoiInit 初始化Aoi
func AoiInit(minX, maxX, minY, maxY, countX, countY int) *AOI {
	aoiM := &AOI{
		MinX:   minX,
		MaxX:   maxX,
		MinY:   minY,
		MaxY:   maxY,
		CountX: countX,
		CountY: countY,
		grids:  make(map[int]*Grid),
	}
	//初始化计算每个格子的坐标
	//这一块是AOI的重点的地方
	//X y 分别为格子的编号

	//平均每个格子x的宽度
	avgX := (maxX - minX) / countX
	//平均每个格子y的高度
	avgY := (maxY - minY) / countY

	for y := 0; y < countY; y++ {
		for x := 0; x < countX; x++ {

			gid := y*countX + x
			//初始化gid格子
			aoiM.grids[gid] = GridInit(
				gid,
				avgX*x,
				avgX*(x+1),
				avgY*y,
				avgY*(y+1))
		}

	}
	return aoiM
}

// GetAvgWidth 获取每个格子的平均宽度
func (a *AOI) GetAvgWidth() int {
	return (a.MaxX - a.MinX) / a.CountX
}

// GetAvgLength 获取每个格子的平均长度
func (a *AOI) GetAvgLength() int {
	return (a.MaxY - a.MinY) / a.CountY
}
func (a *AOI) String() string {

	str := fmt.Sprintf("AOI:\n minX:%d,maxX:%d,minY:%d,maxY:%d,countX:%d,countY:%d Grids in AOI \n", a.MinX, a.MaxX, a.MinY, a.MaxY, a.CountX, a.CountY)
	for _, g := range a.grids {
		str += fmt.Sprintln(g)
	}
	return str
}

// GetRoundGridsByGid AOI九宫格算法 这一块不懂 用的现成的
// 个人感觉更快的解决方案是使用redis
// 用户加入的时候 redis 记录这个用户处于哪个格子 每次移动 修改相应数据
// 地图初始化的时候可以直接存储所有AOI九宫格数据到Redis 因为很多场景下大的地图是不会变的
// 每次查询可以直接走redis进行广播 这一部分数据也要入库 要做到高一致性 否则用户可能会瞬移
func (a *AOI) GetRoundGridsByGid(gID int) (grIDs []*Grid) {
	//判断gID是否存在
	if _, ok := a.grids[gID]; !ok {
		return
	}

	//将当前gID添加到九宫格中
	grIDs = append(grIDs, a.grids[gID])

	// 根据gID, 得到格子所在的坐标
	x, y := gID%a.CountX, gID/a.CountX

	// 新建一个临时存储周围格子的数组
	surroundGID := make([]int, 0)

	// 新建8个方向向量: 左上: (-1, -1), 左中: (-1, 0), 左下: (-1,1), 中上: (0,-1), 中下: (0,1), 右上:(1, -1)
	// 右中: (1, 0), 右下: (1, 1), 分别将这8个方向的方向向量按顺序写入x, y的分量数组
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	// 根据8个方向向量, 得到周围点的相对坐标, 挑选出没有越界的坐标, 将坐标转换为gID
	for i := 0; i < 8; i++ {
		newX := x + dx[i]
		newY := y + dy[i]

		if newX >= 0 && newX < a.CountX && newY >= 0 && newY < a.CountY {
			surroundGID = append(surroundGID, newY*a.CountX+newX)
		}
	}

	// 根据没有越界的gID, 得到格子信息
	for _, gID := range surroundGID {
		grIDs = append(grIDs, a.grids[gID])
	}

	return
}

// GetRoundPlayerIdsByPods 通过坐标获取九宫格内所有的玩家ID
func (a *AOI) GetRoundPlayerIdsByPos(x, y float32) (playerIds []int) {
	//获取格子ID
	gid := a.GetGidByPos(x, y)
	//通过格子ID获取九宫格
	grids := a.GetRoundGridsByGid(gid)
	//通过九宫格获取Players
	for _, v := range grids {
		playerIds = append(playerIds, v.GetAllPlayer()...)
		fmt.Println("------>Grid : %d ,pids : %v<--------------", v.Gid, v.GetAllPlayer())
	}
	return
}

// GetGidByPos 通过坐标获取格子ID
func (a *AOI) GetGidByPos(x, y float32) (gid int) {
	idx := (int(x) - a.MinX) / a.GetAvgWidth()
	idy := (int(y) - a.MinY) / a.GetAvgLength()
	return idy*a.CountX + idx
}
