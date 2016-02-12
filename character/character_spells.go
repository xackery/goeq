package character

import (
	"database/sql"
)

type CharacterSpells struct {
	Id       sql.NullInt64 `db:"id"`       //int(11) unsigned	NO	PRI	NULL	auto_increment
	Slot_id  int           `db:"slot_id"`  //smallint(11) unsigned	NO	PRI	0
	Spell_id int           `db:"spell_id"` //	smallint(11) unsigned	NO		0
}
