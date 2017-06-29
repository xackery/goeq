package npc

import (
	"database/sql"
	taknpc "github.com/xackery/goeq/takp/npc"
)

type NpcTypes struct {
	Id                    sql.NullInt64  `db:"id"`                    //	int(11)	NO	PRI	NULL	auto_increment
	Name                  string         `db:"name"`                  //	text	NO		NULL
	Lastname              sql.NullString `db:"lastname"`              //	varchar(32)	YES		NULL
	Level                 int            `db:"level"`                 //	tinyint(2) unsigned	NO		0
	Race                  int            `db:"race"`                  //	smallint(5) unsigned	NO		0
	Class                 int            `db:"class"`                 //	tinyint(2) unsigned	NO		0
	Bodytype              int            `db:"bodytype"`              //	int(11)	YES		NULL
	Hp                    int            `db:"hp"`                    //	int(11)	NO		0
	Mana                  int            `db:"mana"`                  //	int(11)	NO		0
	Gender                int            `db:"gender"`                //	tinyint(2) unsigned	NO		0
	Texture               int            `db:"texture"`               //	tinyint(2) unsigned	NO		0
	Helmtexture           int            `db:"helmtexture"`           //	tinyint(2) unsigned	NO		0
	Herosforgemodel       int            `db:"herosforgemodel"`       //	int(11) NOT NULL DEFAULT '0',
	Size                  float64        `db:"size"`                  //	float	NO		0
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
	Npcspecialattks       string         `db:"npcspecialattks"`       //	varchar(36)	NO
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
	D_melee_texture1      int            `db:"d_melee_texture1"`      //	int(10) unsigned	NO		0
	D_melee_texture2      int            `db:"d_melee_texture2"`      //	int(10) unsigned	NO		0
	Ammo_idfile           string         `db:"ammo_idfile"`           //	varchar(30)	NO		IT10
	Prim_melee_type       int            `db:"prim_melee_type"`       //	tinyint(4) unsigned	NO		28
	Sec_melee_type        int            `db:"sec_melee_type"`        //	tinyint(4) unsigned	NO		28
	Ranged_type           int            `db:"ranged_type"`           //	tinyint(4) unsigned	NO		7
	Runspeed              float64        `db:"runspeed"`              //	float	NO		0
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
	Attack_speed          float64        `db:"attack_speed"`          //	float	NO		0
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
	Spellscale            float64        `db:"spellscale"`            //	float	NO		100
	Healscale             float64        `db:"healscale"`             //	float	NO		100
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
	Ignore_despawn        int            `db:"ignore_despawn"`        //    tinyint(2) NOT NULL DEFAULT '0',
}

//Convert Takp to peq
func (n *NpcTypes) DecodeFromTakp(tnpc *taknpc.NpcTypes) (err error) {
	//heroesforgemodel default ok
	//attack_speed default ok
	//chesttexture takp default ok
	//combat_hp_regen takp can be ignored, default ok
	//combat_mana_regen takp
	//aggro_pc takp
	//ignore_distance takp
	n.Id.Scan(tnpc.Id)
	n.Name = tnpc.Name
	n.Lastname = tnpc.Lastname
	n.Level = tnpc.Level
	n.Race = tnpc.Race
	n.Class = tnpc.Class
	n.Bodytype = tnpc.Bodytype
	n.Hp = tnpc.Hp
	n.Mana = tnpc.Mana
	n.Gender = tnpc.Gender
	n.Texture = tnpc.Texture
	n.Helmtexture = tnpc.Helmtexture
	n.Size = tnpc.Size
	n.Hp_regen_rate = tnpc.Hp_regen_rate
	n.Mana_regen_rate = tnpc.Mana_regen_rate
	n.Loottable_id = tnpc.Loottable_id
	n.Merchant_id = tnpc.Merchant_id
	n.Alt_currency_id = tnpc.Alt_currency_id
	n.Npc_spells_id = tnpc.Npc_spells_id
	n.Npc_spells_effects_id = tnpc.Npc_spells_effects_id
	n.Npc_faction_id = tnpc.Npc_faction_id
	n.Adventure_template_id = tnpc.Adventure_template_id
	n.Trap_template = tnpc.Trap_template
	n.Mindmg = tnpc.Mindmg
	n.Maxdmg = tnpc.Maxdmg
	n.Attack_count = tnpc.Attack_count
	n.Npcspecialattks = tnpc.Npcspecialattks
	n.Special_abilities.Scan(tnpc.Special_abilities)
	n.Aggroradius = tnpc.Aggroradius
	n.Assistradius = tnpc.Assistradius
	n.Face = tnpc.Face
	n.Luclin_hairstyle = tnpc.Luclin_hairstyle
	n.Luclin_haircolor = tnpc.Luclin_haircolor
	n.Luclin_eyecolor = tnpc.Luclin_eyecolor
	n.Luclin_eyecolor2 = tnpc.Luclin_eyecolor2
	n.Luclin_beardcolor = tnpc.Luclin_beardcolor
	n.Luclin_beard = tnpc.Luclin_beard
	n.Drakkin_heritage = tnpc.Drakkin_heritage
	n.Drakkin_tattoo = tnpc.Drakkin_tattoo
	n.Drakkin_details = tnpc.Drakkin_details
	n.Armortint_id = tnpc.Armortint_id
	n.Armortint_red = tnpc.Armortint_red
	n.Armortint_green = tnpc.Armortint_green
	n.Armortint_blue = tnpc.Armortint_blue
	n.D_melee_texture1 = tnpc.D_melee_texture1
	n.D_melee_texture2 = tnpc.D_melee_texture2
	n.Ammo_idfile = tnpc.Ammo_idfile
	n.Prim_melee_type = tnpc.Prim_melee_type
	n.Sec_melee_type = tnpc.Sec_melee_type
	n.Ranged_type = tnpc.Ranged_type
	n.Runspeed = tnpc.Runspeed
	n.Mr = tnpc.Mr
	n.Cr = tnpc.Cr
	n.Dr = tnpc.Dr
	n.Fr = tnpc.Fr
	n.Pr = tnpc.Pr
	n.Corrup = tnpc.Corrup
	n.Phr = tnpc.Phr
	n.See_invis = tnpc.See_invis
	n.See_invis_undead = tnpc.See_invis_undead
	n.Qglobal = tnpc.Qglobal
	n.Ac = tnpc.Ac
	n.Npc_aggro = tnpc.Npc_aggro
	n.Spawn_limit = tnpc.Spawn_limit
	//n.Attack_speed = tnpc.Attack_speed
	n.Attack_delay = tnpc.Attack_delay
	n.Findable = tnpc.Findable
	n.Str = tnpc.Str
	n.Sta = tnpc.Sta
	n.Dex = tnpc.Dex
	n.Agi = tnpc.Agi
	n.Int = tnpc.Int
	n.Wis = tnpc.Wis
	n.Cha = tnpc.Cha
	n.See_hide = tnpc.See_hide
	n.See_improved_hide = tnpc.See_improved_hide
	n.Trackable = tnpc.Trackable
	n.Isbot = tnpc.Isbot
	n.Exclude = tnpc.Exclude
	n.Atk = tnpc.Atk
	n.Accuracy = tnpc.Accuracy
	n.Avoidance = tnpc.Avoidance
	n.Slow_mitigation = tnpc.Slow_mitigation
	n.Version = tnpc.Version
	n.Maxlevel = tnpc.Maxlevel
	n.Scalerate = tnpc.Scalerate
	n.Private_corpse = tnpc.Private_corpse
	n.Unique_spawn_by_name = tnpc.Unique_spawn_by_name
	n.Underwater = tnpc.Underwater
	n.Isquest = tnpc.Isquest
	n.Emoteid = tnpc.Emoteid
	n.Spellscale = tnpc.Spellscale
	n.Healscale = tnpc.Healscale
	n.No_target_hotkey = tnpc.No_target_hotkey
	n.Raid_target = tnpc.Raid_target
	n.Armtexture = tnpc.Armtexture
	n.Bracertexture = tnpc.Bracertexture
	n.Handtexture = tnpc.Handtexture
	n.Legtexture = tnpc.Legtexture
	n.Feettexture = tnpc.Feettexture
	n.Light = tnpc.Light
	//n.Walkspeed = tnpc.Walkspeed
	n.Peqid = tnpc.Peqid
	n.Unique = tnpc.Unique
	n.Fixed = tnpc.Fixed
	return
}
