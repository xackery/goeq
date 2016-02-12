package character

import (
	"database/sql"
)

type CharacterSkills struct {
	Id       sql.NullInt64 `db:"id"`       //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Skill_id int           `db:"skill_id"` //	smallint(11) unsigned	NO	PRI	0
	Value    int           `db:"value"`    //	smallint(11) unsigned	NO		0
}
