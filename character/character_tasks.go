package character

import (
	"database/sql"
)

type CharacterTasks struct {
	Charid       int           `db"charid"`       //	int(11) unsigned	NO	PRI	0
	Taskid       int           `db"taskid"`       //	int(11) unsigned	NO	PRI	0
	Slot         int           `db"slot"`         //	int(11) unsigned	NO		0
	Acceptedtime sql.NullInt64 `db"acceptedtime"` //	int(11) unsigned	YES		NULL
}
