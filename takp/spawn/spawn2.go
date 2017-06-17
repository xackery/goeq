package spawn

import (
	"database/sql"
)

type Spawn2 struct {
	Id                 sql.NullInt64  `db:"id"`                 //  int(11) NO  PRI NULL  auto_increment
	Spawngroupid       int            `db:"spawngroupID"`       //  int(11) NO  MUL 0
	Zone               sql.NullString `db:"zone"`               //  varchar(32) YES MUL NULL
	Version            int            `db:"version"`            //  smallint(5) unsigned  NO    0
	X                  float64        `db:"x"`                  //  float(14,6) NO    0.000000
	Y                  float64        `db:"y"`                  //  float(14,6) NO    0.000000
	Z                  float64        `db:"z"`                  //  float(14,6) NO    0.000000
	Heading            float64        `db:"heading"`            //  float(14,6) NO    0.000000
	Respawntime        int            `db:"respawntime"`        //  int(11) NO    0
	Variance           int            `db:"variance"`           //  int(11) NO    0
	Pathgrid           int            `db:"pathgrid"`           //  int(10) NO    0
	Condition          int            `db:"_condition"`         //  mediumint(8) unsigned NO    0
	Cond_value         int            `db:"cond_value"`         //  mediumint(9)  NO    1
	Enabled            int            `db:"enabled"`            //  tinyint(3) unsigned NO    1
	Animation          int            `db:"animation"`          //  tinyint(3) unsigned NO    0
	Boot_respawntime   int            `db:"boot_respawntime"`   // int(11) NOT NULL DEFAULT '0',
	Clear_timer_onboot int            `db:"clear_timer_onboot"` // tinyint(3) NOT NULL DEFAULT '0',
}
