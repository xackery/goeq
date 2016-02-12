package character

import ()

type CharacterPotionBelt struct {
	Id        int `db:id"`        //	int(11) unsigned	NO	PRI	0
	Potion_id int `db:potion_id"` //	tinyint(11) unsigned	NO	PRI	0
	Item_id   int `db:item_id"`   //	int(11) unsigned	NO		0
	Icon      int `db:icon"`      //	int(11) unsigned	NO		0
}
