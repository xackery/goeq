package character

import ()

type CharacterCurrency struct {
	Id                      int `db:"id"`                      //	int(11) unsigned	NO	PRI	0
	Platinum                int `db:"platinum"`                //	int(11) unsigned	NO		0
	Gold                    int `db:"gold"`                    //	int(11) unsigned	NO		0
	Silver                  int `db:"silver"`                  //	int(11) unsigned	NO		0
	Copper                  int `db:"copper"`                  //	int(11) unsigned	NO		0
	Platinum_bank           int `db:"platinum_bank"`           //	int(11) unsigned	NO		0
	Gold_bank               int `db:"gold_bank"`               //	int(11) unsigned	NO		0
	Silver_bank             int `db:"silver_bank"`             //	int(11) unsigned	NO		0
	Copper_bank             int `db:"copper_bank"`             //	int(11) unsigned	NO		0
	Platinum_cursor         int `db:"platinum_cursor"`         //	int(11) unsigned	NO		0
	Gold_cursor             int `db:"gold_cursor"`             //	int(11) unsigned	NO		0
	Silver_cursor           int `db:"silver_cursor"`           //	int(11) unsigned	NO		0
	Copper_cursor           int `db:"copper_cursor"`           //	int(11) unsigned	NO		0
	Radiant_crystals        int `db:"radiant_crystals"`        //	int(11) unsigned	NO		0
	Career_radiant_crystals int `db:"career_radiant_crystals"` //	int(11) unsigned	NO		0
	Ebon_crystals           int `db:"ebon_crystals"`           //	int(11) unsigned	NO		0
	Career_ebon_crystals    int `db:"career_ebon_crystals"`    //	int(11) unsigned	NO		0
}
