package loot

import ()

type LootTableEntries struct {
	Loottable_id   int     `db:"loottable_id"`   //	int(11) unsigned NOT NULL DEFAULT '0',
	Lootdrop_id    int     `db:"lootdrop_id"`    //	int(11) unsigned NOT NULL DEFAULT '0',
	Multiplier     int     `db:"multiplier"`     //	tinyint(2) unsigned NOT NULL DEFAULT '1',
	Droplimit      int     `db:"droplimit"`      //	tinyint(2) unsigned NOT NULL DEFAULT '0',
	Mindrop        int     `db:"mindrop"`        //	tinyint(2) unsigned NOT NULL DEFAULT '0',
	Probability    float64 `db:"probability"`    //	tinyint(2) unsigned NOT NULL DEFAULT '100',
	Multiplier_min int     `db:"multiplier_min"` // tinyint(2) unsigned NOT NULL DEFAULT '0',
}
