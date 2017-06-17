package loot

import (
//"database/sql"
//takloot "github.com/xackery/goeq/takp/loot"
)

type LootDropEntries struct {
	Lootdrop_id     int     `db:"lootdrop_id"`     // int(11) unsigned NOT NULL DEFAULT '0',
	Item_id         int     `db:"item_id"`         // int(11) NOT NULL DEFAULT '0',
	Item_charges    int     `db:"item_charges"`    // smallint(2) unsigned NOT NULL DEFAULT '1',
	Equip_item      int     `db:"equip_item"`      // tinyint(2) unsigned NOT NULL DEFAULT '0',
	Chance          float64 `db:"chance"`          // float NOT NULL DEFAULT '1',
	Disabled_chance float64 `db:"disabled_chance"` // float NOT NULL DEFAULT '0',
	Minlevel        int     `db:"minlevel"`        // tinyint(3) NOT NULL DEFAULT '0',
	Maxlevel        int     `db:"maxlevel"`        // tinyint(3) NOT NULL DEFAULT '127',
	Multiplier      int     `db:"multiplier"`      // tinyint(2) unsigned NOT NULL DEFAULT '1',
}
