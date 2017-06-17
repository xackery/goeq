package loot

import (
//"database/sql"
)

type LootTable struct {
	Id      int    `db:"id"`      //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Name    string `db:"name"`    //	varchar(255)	NO
	Mincash int    `db:"mincash"` //	int(11) unsigned	NO		0
	Maxcash int    `db:"maxcash"` //	int(11) unsigned	NO		0
	Avgcoin int    `db:"avgcoin"` //	smallint(4) unsigned	NO		0
	Done    int    `db:"done"`    //	tinyint(3)	NO		0
}
