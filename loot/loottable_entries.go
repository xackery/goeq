package loot

import (
	takloot "github.com/xackery/goeq/takp/loot"
)

type LootTableEntries struct {
	Loottable_id int     `db:"loottable_id"` //	int(11) unsigned	NO	PRI	0
	Lootdrop_id  int     `db:"lootdrop_id"`  //	int(11) unsigned	NO	PRI	0
	Multiplier   int     `db:"multiplier"`   //	tinyint(2) unsigned	NO		1
	Droplimit    int     `db:"droplimit"`    //	tinyint(2) unsigned	NO		0
	Mindrop      int     `db:"mindrop"`      //	tinyint(2) unsigned	NO		0
	Probability  float64 `db:"probability"`  //	float	NO		100
}

//Convert Takp to peq
func (n *LootTableEntries) DecodeFromTakp(tl *takloot.LootTableEntries) (err error) {
	n.Loottable_id = tl.Loottable_id
	n.Lootdrop_id = tl.Lootdrop_id
	n.Multiplier = tl.Multiplier
	n.Droplimit = tl.Droplimit
	n.Mindrop = tl.Mindrop
	n.Probability = tl.Probability
	return
}
