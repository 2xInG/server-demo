package core

import (
	"fmt"
	"testing"
)

func TestAOI(t *testing.T) {
	aoiMG := AoiInit(100, 1000, 300, 1000, 9, 9)
	fmt.Println(aoiMG)
}

func TestAOIM(t *testing.T) {
	aoiMG := AoiInit(100, 1000, 300, 1000, 9, 9)
	for k, _ := range aoiMG.grids {
		//得到当前格子周边的九宫格
		grIDs := aoiMG.GetRoundGridsByGid(k)
		//得到九宫格所有的IDs
		fmt.Println("gID : ", k, " grIDs len = ", len(grIDs))
		gIDs := make([]int, 0, len(grIDs))
		for _, grID := range grIDs {
			gIDs = append(gIDs, grID.Gid)
		}
		fmt.Printf("grID ID: %d, surrounding grID IDs are %v\n", k, gIDs)
	}
}
