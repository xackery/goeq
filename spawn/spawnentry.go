package spawn

import ()

type SpawnEntry struct {
	spawngroupID int `db:"spawngroupID"` //	int(11)	NO	PRI	0
	npcID        int `db:"npcID"`        //	int(11)	NO	PRI	0
	chance       int `db:"chance"`       //	smallint(4)	NO		0
}
