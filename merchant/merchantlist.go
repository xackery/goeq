package merchant

import (
//"database/sql"
)

type MerchantList struct {
	Merchantid        int `db:"merchantid"`        // int(11) NOT NULL DEFAULT '0',
	Slot              int `db:"slot"`              // int(11) NOT NULL DEFAULT '0',
	Item              int `db:"item"`              // int(11) NOT NULL DEFAULT '0',
	Faction_required  int `db:"faction_required"`  // smallint(6) NOT NULL DEFAULT '-100',
	Level_required    int `db:"level_required"`    // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Alt_currency_cost int `db:"alt_currency_cost"` // smallint(5) unsigned NOT NULL DEFAULT '0',
	Classes_required  int `db:"classes_required"`  // int(11) NOT NULL DEFAULT '65535',
	Probability       int `db:"probability"`       // int(3) NOT NULL DEFAULT '100',
}
