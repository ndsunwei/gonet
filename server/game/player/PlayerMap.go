package player

import (
	"gonet/server/game"
)

func (this *Player) SendToZone(funcName string, params  ...interface{}) {
	game.SendToZone(this.PlayerId, this.Raft.ZClusterId, funcName, params...)
}

func (this *Player) AddMap() {
	this.SendToZone("LoginMap", 200000, this.PlayerId, this.GetGateClusterId(), game.SERVER.GetCluster().Id())
}

func (this *Player) LeaveMap() {
	this.SendToZone("LogoutMap", 200000, this.PlayerId)
}

func (this *Player) ReloginMap() {
	this.SendToZone("ReloginMap", this.PlayerId, this.GetGateClusterId())
}

//添加buff
func (this *Player) AddBuff(Orgint int, BuffId int) {
	if BuffId < 0{
		return
	}
	this.SendToZone("AddBuff", this.PlayerId, Orgint, BuffId)
}

//删除buff
func (this *Player) RemoveBuff(BuffId int) {
	if BuffId < 0{
		return
	}
	this.SendToZone("RemoveBuff", this.PlayerId, BuffId)
}

//批量添加buff
func (this *Player) AddBuffS(Orgint int, BuffId []int) {
	BuffIds :=  []int{}
	for i := 0; i < len(BuffId); i++{
		if BuffId[i] < 0{
			continue
		}
		BuffIds = append(BuffIds, int(BuffId[i]))
	}

	this.SendToZone("AddBuffS", this.PlayerId, Orgint, BuffIds)
}

//批量删除buff
func (this *Player) RemoveBuffS(BuffId []int) {
	BuffIds :=  []int{}
	for i := 0; i < len(BuffId); i++{
		if BuffId[i] < 0{
			continue
		}
		BuffIds = append(BuffIds, int(BuffId[i]))
	}
	this.SendToZone("RemoveBuffS", this.PlayerId, BuffIds)
}