package spawn

import ()

type SpawnEntry struct {
	Spawngroupid int `db:"spawngroupID"` //	int(11)	NO	PRI	0
	Npcid        int `db:"npcID"`        //	int(11)	NO	PRI	0
	Chance       int `db:"chance"`       //	smallint(4)	NO		0
	Mintime      int `db:"mintime"`      // smallint(4) NOT NULL DEFAULT '0',
	Maxtime      int `db:"maxtime"`      // smallint(4) NOT NULL DEFAULT '0',
}
