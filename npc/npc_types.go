package npc

import (
	"database/sql"
)

type NpcTypes struct {
	Id                    sql.NullInt64  `db:"id"`                    //	int(11)	NO	PRI	NULL	auto_increment
	Name                  sql.NullString `db:"name"`                  //	text	NO		NULL
	Lastname              sql.NullString `db:"lastname"`              //	varchar(32)	YES		NULL
	Level                 int            `db:"level"`                 //	tinyint(2) unsigned	NO		0
	Race                  int            `db:"race"`                  //	smallint(5) unsigned	NO		0
	Class                 int            `db:"class"`                 //	tinyint(2) unsigned	NO		0
	Bodytype              sql.NullInt64  `db:"bodytype"`              //	int(11)	YES		NULL
	Hp                    int            `db:"hp"`                    //	int(11)	NO		0
	Mana                  int            `db:"mana"`                  //	int(11)	NO		0
	Gender                int            `db:"gender"`                //	tinyint(2) unsigned	NO		0
	Texture               int            `db:"texture"`               //	tinyint(2) unsigned	NO		0
	Helmtexture           int            `db:"helmtexture"`           //	tinyint(2) unsigned	NO		0
	Size                  int            `db:"size"`                  //	float	NO		0
	Hp_regen_rate         int            `db:"hp_regen_rate"`         //	int(11) unsigned	NO		0
	Mana_regen_rate       int            `db:"mana_regen_rate"`       //	int(11) unsigned	NO		0
	Loottable_id          int            `db:"loottable_id"`          //	int(11) unsigned	NO		0
	Merchant_id           int            `db:"merchant_id"`           //	int(11) unsigned	NO		0
	Alt_currency_id       int            `db:"alt_currency_id"`       //	int(11) unsigned	NO		0
	Npc_spells_id         int            `db:"npc_spells_id"`         //	int(11) unsigned	NO		0
	Npc_spells_effects_id int            `db:"npc_spells_effects_id"` //	int(11) unsigned	NO		0
	Npc_faction_id        int            `db:"npc_faction_id"`        //	int(11)	NO		0
	Adventure_template_id int            `db:"adventure_template_id"` //	int(10) unsigned	NO		0
	Trap_template         int            `db:"trap_template"`         //	int(10) unsigned	YES		0
	Mindmg                int            `db:"mindmg"`                //	int(10) unsigned	NO		0
	Maxdmg                int            `db:"maxdmg"`                //	int(10) unsigned	NO		0
	Attack_count          int            `db:"attack_count"`          //	smallint(6)	NO		-1
	Npcspecialattks       int            `db:"npcspecialattks"`       //	varchar(36)	NO
	Special_abilities     sql.NullString `db:"special_abilities"`     //	text	NO		NULL
	Aggroradius           int            `db:"aggroradius"`           //	int(10) unsigned	NO		0
	Assistradius          int            `db:"assistradius"`          //	int(10) unsigned	NO		0
	Face                  int            `db:"face"`                  //	int(10) unsigned	NO		1
	Luclin_hairstyle      int            `db:"luclin_hairstyle"`      //	int(10) unsigned	NO		1
	Luclin_haircolor      int            `db:"luclin_haircolor"`      //	int(10) unsigned	NO		1
	Luclin_eyecolor       int            `db:"luclin_eyecolor"`       //	int(10) unsigned	NO		1
	Luclin_eyecolor2      int            `db:"luclin_eyecolor2"`      //	int(10) unsigned	NO		1
	Luclin_beardcolor     int            `db:"luclin_beardcolor"`     //	int(10) unsigned	NO		1
	Luclin_beard          int            `db:"luclin_beard"`          //	int(10) unsigned	NO		0
	Drakkin_heritage      int            `db:"drakkin_heritage"`      //	int(10)	NO		0
	Drakkin_tattoo        int            `db:"drakkin_tattoo"`        //	int(10)	NO		0
	Drakkin_details       int            `db:"drakkin_details"`       //	int(10)	NO		0
	Armortint_id          int            `db:"armortint_id"`          //	int(10) unsigned	NO		0
	Armortint_red         int            `db:"armortint_red"`         //	tinyint(3) unsigned	NO		0
	Armortint_green       int            `db:"armortint_green"`       //	tinyint(3) unsigned	NO		0
	Armortint_blue        int            `db:"armortint_blue"`        //	tinyint(3) unsigned	NO		0
	D_meele_texture1      int            `db:"d_meele_texture1"`      //	int(10) unsigned	NO		0
	D_meele_texture2      int            `db:"d_meele_texture2"`      //	int(10) unsigned	NO		0
	Ammo_idfile           int            `db:"ammo_idfile"`           //	varchar(30)	NO		IT10
	Prim_melee_type       int            `db:"prim_melee_type"`       //	tinyint(4) unsigned	NO		28
	Sec_melee_type        int            `db:"sec_melee_type"`        //	tinyint(4) unsigned	NO		28
	Ranged_type           int            `db:"ranged_type"`           //	tinyint(4) unsigned	NO		7
	Runspeed              int            `db:"runspeed"`              //	float	NO		0
	Mr                    int            `db:"MR"`                    //	smallint(5)	NO		0
	Cr                    int            `db:"CR"`                    //	smallint(5)	NO		0
	Dr                    int            `db:"DR"`                    //	smallint(5)	NO		0
	Fr                    int            `db:"FR"`                    //	smallint(5)	NO		0
	Pr                    int            `db:"PR"`                    //	smallint(5)	NO		0
	Corrup                int            `db:"Corrup"`                //	smallint(5)	NO		0
	Phr                   int            `db:"PhR"`                   //	smallint(5) unsigned	NO		0
	See_invis             int            `db:"see_invis"`             //	smallint(4)	NO		0
	See_invis_undead      int            `db:"see_invis_undead"`      //	smallint(4)	NO		0
	Qglobal               int            `db:"qglobal"`               //	int(2) unsigned	NO		0
	Ac                    int            `db:"AC"`                    //	smallint(5)	NO		0
	Npc_aggro             int            `db:"npc_aggro"`             //	tinyint(4)	NO		0
	Spawn_limit           int            `db:"spawn_limit"`           //	tinyint(4)	NO		0
	Attack_speed          int            `db:"attack_speed"`          //	float	NO		0
	Attack_delay          int            `db:"attack_delay"`          //	tinyint(3) unsigned	NO		30
	Findable              int            `db:"findable"`              //	tinyint(4)	NO		0
	Str                   int            `db:"STR"`                   //	mediumint(8) unsigned	NO		75
	Sta                   int            `db:"STA"`                   //	mediumint(8) unsigned	NO		75
	Dex                   int            `db:"DEX"`                   //	mediumint(8) unsigned	NO		75
	Agi                   int            `db:"AGI"`                   //	mediumint(8) unsigned	NO		75
	Int                   int            `db:"_INT"`                  //	mediumint(8) unsigned	NO		80
	Wis                   int            `db:"WIS"`                   //	mediumint(8) unsigned	NO		75
	Cha                   int            `db:"CHA"`                   //	mediumint(8) unsigned	NO		75
	See_hide              int            `db:"see_hide"`              //	tinyint(4)	NO		0
	See_improved_hide     int            `db:"see_improved_hide"`     //	tinyint(4)	NO		0
	Trackable             int            `db:"trackable"`             //	tinyint(4)	NO		1
	Isbot                 int            `db:"isbot"`                 //	tinyint(4)	NO		0
	Exclude               int            `db:"exclude"`               //	tinyint(4)	NO		1
	Atk                   int            `db:"ATK"`                   //	mediumint(9)	NO		0
	Accuracy              int            `db:"Accuracy"`              //	mediumint(9)	NO		0
	Avoidance             int            `db:"Avoidance"`             //	mediumint(9) unsigned	NO		0
	Slow_mitigation       int            `db:"slow_mitigation"`       //	smallint(4)	NO		0
	Version               int            `db:"version"`               //	smallint(5) unsigned	NO		0
	Maxlevel              int            `db:"maxlevel"`              //	tinyint(3)	NO		0
	Scalerate             int            `db:"scalerate"`             //	int(11)	NO		100
	Private_corpse        int            `db:"private_corpse"`        //	tinyint(3) unsigned	NO		0
	Unique_spawn_by_name  int            `db:"unique_spawn_by_name"`  //	tinyint(3) unsigned	NO		0
	Underwater            int            `db:"underwater"`            //	tinyint(3) unsigned	NO		0
	Isquest               int            `db:"isquest"`               //	tinyint(3)	NO		0
	Emoteid               int            `db:"emoteid"`               //	int(10) unsigned	NO		0
	Spellscale            int            `db:"spellscale"`            //	float	NO		100
	Healscale             int            `db:"healscale"`             //	float	NO		100
	No_target_hotkey      int            `db:"no_target_hotkey"`      //	tinyint(1) unsigned	NO		0
	Raid_target           int            `db:"raid_target"`           //	tinyint(1) unsigned	NO		0
	Armtexture            int            `db:"armtexture"`            //	tinyint(2)	NO		0
	Bracertexture         int            `db:"bracertexture"`         //	tinyint(2)	NO		0
	Handtexture           int            `db:"handtexture"`           //	tinyint(2)	NO		0
	Legtexture            int            `db:"legtexture"`            //	tinyint(2)	NO		0
	Feettexture           int            `db:"feettexture"`           //	tinyint(2)	NO		0
	Light                 int            `db:"light"`                 //	tinyint(2)	NO		0
	Walkspeed             int            `db:"walkspeed"`             //	tinyint(2)	NO		0
	Peqid                 int            `db:"peqid"`                 //	int(11)	NO		0
	Unique                int            `db:"unique_"`               //	tinyint(2)	NO		0
	Fixed                 int            `db:"fixed"`                 //	tinyint(2)	NO		0
}

//Convert Takp to peq
func (n *NpcTypes) DecodeFromTakp(npc *NpcTypes) {
	//heroesforgemodel default ok
	//attack_speed default ok
	//chesttexture takp default ok
	//combat_hp_regen takp can be ignored, default ok
	//combat_mana_regen takp
	//aggro_pc takp
	//ignore_distance takp
}
