package inventory

import (
	"database/sql"
)

//Inventory is awesome
type Inventory struct {
	Charid              int            `db:"charid"`      //	int(11) unsigned	NO	PRI	0
	Slotid              int            `db:"slotid"`      //	mediumint(7) unsigned	NO	PRI	0
	Itemid              int            `db:"itemid"`      //	int(11) unsigned	YES		0
	Charges             int            `db:"charges"`     //	smallint(3) unsigned	YES		0
	Color               int            `db:"color"`       //	int(11) unsigned	NO		0
	Augslot1            int            `db:"augslot1"`    //	mediumint(7) unsigned	NO		0
	Augslot2            int            `db:"augslot2"`    //	mediumint(7) unsigned	NO		0
	Augslot3            int            `db:"augslot3"`    //	mediumint(7) unsigned	NO		0
	Augslot4            int            `db:"augslot4"`    //	mediumint(7) unsigned	NO		0
	Augslot5            int            `db:"augslot5"`    //	mediumint(7) unsigned	YES		0
	Augslot6            int            `db:"augslot6"`    //	mediumint(7) unsigned	YES		0
	Instnodrop          int            `db:"instnodrop"`  //	tinyint(1) unsigned	NO		0
	Customdata          sql.NullString `db:"custom_data"` //	text	YES		NULL
	Ornamenticon        int            `db:"ornamenticon"`
	Ornamentidfile      int            `db:"ornamentidfile"`
	Ornament_hero_model int            `db:"ornament_hero_model"`
}
