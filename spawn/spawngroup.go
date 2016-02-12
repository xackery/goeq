package spawn

import (
	"database/sql"
)

type SpawnGroup struct {
	Id            sql.NullInt64  `db:"id"`            //	int(11)	NO	PRI	NULL	auto_increment
	Name          sql.NullString `db:"name"`          //	varchar(50)	NO	UNI
	Spawn_limit   int            `db:"spawn_limit"`   //	tinyint(4)	NO		0
	Dist          float64        `db:"dist"`          //	float	NO		0
	Max_x         float64        `db:"max_x"`         //	float	NO		0
	Min_x         float64        `db:"min_x"`         //	float	NO		0
	Max_y         float64        `db:"max_y"`         //	float	NO		0
	Min_y         float64        `db:"min_y"`         //	float	NO		0
	Delay         int            `db:"delay"`         //	int(11)	NO		45000
	Mindelay      int            `db:"mindelay"`      //	int(11)	NO		15000
	Despawn       int            `db:"despawn"`       //	tinyint(3)	NO		0
	Despawn_timer int            `db:"despawn_timer"` //	int(11)	NO		100
}
