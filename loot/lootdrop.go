package loot

import ()

type LootDrop struct {
	Id   int    `db:"id"`   // `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	Name string `db:"name"` // `name` varchar(255) NOT NULL DEFAULT '',
}
