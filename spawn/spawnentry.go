package spawn

import (
	takspawn "github.com/xackery/goeq/takp/spawn"
)

type SpawnEntry struct {
	Spawngroupid int `db:"spawngroupID"` //	int(11)	NO	PRI	0
	Npcid        int `db:"npcID"`        //	int(11)	NO	PRI	0
	Chance       int `db:"chance"`       //	smallint(4)	NO		0
}

//Convert Takp to peq
func (n *SpawnEntry) DecodeFromTakp(tnpc *takspawn.SpawnEntry) (err error) {
	n.Spawngroupid = tnpc.Spawngroupid
	n.Chance = tnpc.Chance
	n.Npcid = tnpc.Npcid
	return
}
