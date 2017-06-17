package loot

import (
	"database/sql"
	takloot "github.com/xackery/goeq/takp/loot"
)

type LootTable struct {
	Id      sql.NullInt64  `db:"id"`      //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Name    sql.NullString `db:"name"`    //	varchar(255)	NO
	Mincash int            `db:"mincash"` //	int(11) unsigned	NO		0
	Maxcash int            `db:"maxcash"` //	int(11) unsigned	NO		0
	Avgcoin int            `db:"avgcoin"` //	smallint(4) unsigned	NO		0
	Done    int            `db:"done"`    //	tinyint(3)	NO		0
}

//Convert Takp to peq
func (n *LootTable) DecodeFromTakp(tl *takloot.LootTable) (err error) {
	n.Id.Scan(tl.Id)
	n.Name.Scan(tl.Name)
	n.Mincash = tl.Mincash
	n.Maxcash = tl.Maxcash
	n.Avgcoin = tl.Avgcoin
	n.Done = tl.Done
	return
}
