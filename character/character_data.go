package character

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type CharacterData struct {
	Id                      int            `db:"id"`                      //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Account_id              int            `db:"account_id"`              //	int(11)	NO	MUL	0
	Name                    string         `db:"name"`                    //	varchar(64)	NO	UNI
	Last_name               string         `db:"last_name"`               //	varchar(64)	NO
	Title                   string         `db:"title"`                   //	varchar(32)	NO
	Suffix                  string         `db:"suffix"`                  //	varchar(32)	NO
	Zone_id                 int            `db:"zone_id"`                 //	int(11) unsigned	NO		0
	Zone_instance           int            `db:"zone_instance"`           //	int(11) unsigned	NO		0
	Y                       float32        `db:"y"`                       //	float	NO		0
	X                       float32        `db:"x"`                       //	float	NO		0
	Z                       float32        `db:"z"`                       //	float	NO		0
	Heading                 float32        `db:"heading"`                 //	float	NO		0
	Gender                  int            `db:"gender"`                  //	tinyint(11) unsigned	NO		0
	Race                    int            `db:"race"`                    //	smallint(11) unsigned	NO		0
	Class                   int            `db:"class"`                   //	tinyint(11) unsigned	NO		0
	Level                   int            `db:"level"`                   //	int(11) unsigned	NO		0
	Deity                   int            `db:"deity"`                   //	int(11) unsigned	NO		0
	Birthday                int            `db:"birthday"`                //	int(11) unsigned	NO		0
	Last_login              int64          `db:"last_login"`              //	int(11) unsigned	NO		0
	Time_played             int            `db:"time_played"`             //	int(11) unsigned	NO		0
	Level2                  int            `db:"level2"`                  //	tinyint(11) unsigned	NO		0
	Anon                    int            `db:"anon"`                    //	tinyint(11) unsigned	NO		0
	Gm                      int            `db:"gm"`                      //	tinyint(11) unsigned	NO		0
	Face                    int            `db:"face"`                    //	int(11) unsigned	NO		0
	Hair_color              int            `db:"hair_color"`              //	tinyint(11) unsigned	NO		0
	Hair_style              int            `db:"hair_style"`              //	tinyint(11) unsigned	NO		0
	Beard                   int            `db:"beard"`                   //	tinyint(11) unsigned	NO		0
	Beard_color             int            `db:"beard_color"`             //	tinyint(11) unsigned	NO		0
	Eye_color_1             int            `db:"eye_color_1"`             //	tinyint(11) unsigned	NO		0
	Eye_color_2             int            `db:"eye_color_2"`             //	tinyint(11) unsigned	NO		0
	Drakkin_heritage        int            `db:"drakkin_heritage"`        //	int(11) unsigned	NO		0
	Drakkin_tattoo          int            `db:"drakkin_tattoo"`          //	int(11) unsigned	NO		0
	Drakkin_details         int            `db:"drakkin_details"`         //	int(11) unsigned	NO		0
	Ability_time_seconds    int            `db:"ability_time_seconds"`    //	tinyint(11) unsigned	NO		0
	Ability_number          int            `db:"ability_number"`          //	tinyint(11) unsigned	NO		0
	Ability_time_minutes    int            `db:"ability_time_minutes"`    //	tinyint(11) unsigned	NO		0
	Ability_time_hours      int            `db:"ability_time_hours"`      //	tinyint(11) unsigned	NO		0
	Exp                     int            `db:"exp"`                     //	int(11) unsigned	NO		0
	Aa_points_spent         int            `db:"aa_points_spent"`         //	int(11) unsigned	NO		0
	Aa_exp                  int            `db:"aa_exp"`                  //	int(11) unsigned	NO		0
	Aa_points               int            `db:"aa_points"`               //	int(11) unsigned	NO		0
	Group_leadership_exp    int            `db:"group_leadership_exp"`    //	int(11) unsigned	NO		0
	Raid_leadership_exp     int            `db:"raid_leadership_exp"`     //	int(11) unsigned	NO		0
	Group_leadership_points int            `db:"group_leadership_points"` //	int(11) unsigned	NO		0
	Raid_leadership_points  int            `db:"raid_leadership_points"`  //	int(11) unsigned	NO		0
	Points                  int            `db:"points"`                  //	int(11) unsigned	NO		0
	Cur_hp                  int            `db:"cur_hp"`                  //	int(11) unsigned	NO		0
	Mana                    int            `db:"mana"`                    //	int(11) unsigned	NO		0
	Endurance               int            `db:"endurance"`               //	int(11) unsigned	NO		0
	Intoxication            int            `db:"intoxication"`            //	int(11) unsigned	NO		0
	Str                     int            `db:"str"`                     //	int(11) unsigned	NO		0
	Sta                     int            `db:"sta"`                     //	int(11) unsigned	NO		0
	Cha                     int            `db:"cha"`                     //	int(11) unsigned	NO		0
	Dex                     int            `db:"dex"`                     //	int(11) unsigned	NO		0
	Int                     int            `db:"int"`                     //	int(11) unsigned	NO		0
	Agi                     int            `db:"agi"`                     //	int(11) unsigned	NO		0
	Wis                     int            `db:"wis"`                     //	int(11) unsigned	NO		0
	Zone_change_count       int            `db:"zone_change_count"`       //	int(11) unsigned	NO		0
	Toxicity                int            `db:"toxicity"`                //	int(11) unsigned	NO		0
	Hunger_level            int            `db:"hunger_level"`            //	int(11) unsigned	NO		0
	Thirst_level            int            `db:"thirst_level"`            //	int(11) unsigned	NO		0
	Ability_up              int            `db:"ability_up"`              //	int(11) unsigned	NO		0
	Ldon_points_guk         int            `db:"ldon_points_guk"`         //	int(11) unsigned	NO		0
	Ldon_points_mir         int            `db:"ldon_points_mir"`         //	int(11) unsigned	NO		0
	Ldon_points_mmc         int            `db:"ldon_points_mmc"`         //	int(11) unsigned	NO		0
	Ldon_points_ruj         int            `db:"ldon_points_ruj"`         //	int(11) unsigned	NO		0
	Ldon_points_tak         int            `db:"ldon_points_tak"`         //	int(11) unsigned	NO		0
	Ldon_points_available   int            `db:"ldon_points_available"`   //	int(11) unsigned	NO		0
	Tribute_time_remaining  int            `db:"tribute_time_remaining"`  //	int(11) unsigned	NO		0
	Career_tribute_points   int            `db:"career_tribute_points"`   //	int(11) unsigned	NO		0
	Tribute_points          int            `db:"tribute_points"`          //	int(11) unsigned	NO		0
	Tribute_active          int            `db:"tribute_active"`          //	int(11) unsigned	NO		0
	Pvp_status              int            `db:"pvp_status"`              //	tinyint(11) unsigned	NO		0
	Pvp_kills               int            `db:"pvp_kills"`               //	int(11) unsigned	NO		0
	Pvp_deaths              int            `db:"pvp_deaths"`              //	int(11) unsigned	NO		0
	Pvp_current_points      int            `db:"pvp_current_points"`      //	int(11) unsigned	NO		0
	Pvp_career_points       int            `db:"pvp_career_points"`       //	int(11) unsigned	NO		0
	Pvp_best_kill_streak    int            `db:"pvp_best_kill_streak"`    //	int(11) unsigned	NO		0
	Pvp_worst_death_streak  int            `db:"pvp_worst_death_streak"`  //	int(11) unsigned	NO		0
	Pvp_current_kill_streak int            `db:"pvp_current_kill_streak"` //	int(11) unsigned	NO		0
	Pvp2                    int            `db:"pvp2"`                    //	int(11) unsigned	NO		0
	Pvp_type                int            `db:"pvp_type"`                //	int(11) unsigned	NO		0
	Show_helm               int            `db:"show_helm"`               //	int(11) unsigned	NO		0
	Group_auto_consent      int            `db:"group_auto_consent"`      //	tinyint(11) unsigned	NO		0
	Raid_auto_consent       int            `db:"raid_auto_consent"`       //	tinyint(11) unsigned	NO		0
	Guild_auto_consent      int            `db:"guild_auto_consent"`      //	tinyint(11) unsigned	NO		0
	Leadership_exp_on       int            `db:"leadership_exp_on"`       //	tinyint(11) unsigned	NO		0
	Resttimer               int            `db:"RestTimer"`               //	int(11) unsigned	NO		0
	Air_remaining           int            `db:"air_remaining"`           //	int(11) unsigned	NO		0
	Autosplit_enabled       int            `db:"autosplit_enabled"`       //	int(11) unsigned	NO		0
	Lfp                     int            `db:"lfp"`                     //	tinyint(1) unsigned	NO		0
	Lfg                     int            `db:"lfg"`                     //	tinyint(1) unsigned	NO		0
	Mailkey                 string         `db:"mailkey"`                 //	char(16)	NO
	Xtargets                int            `db:"xtargets"`                //	tinyint(3) unsigned	NO		5
	Firstlogon              int            `db:"firstlogon"`              //	tinyint(3)	NO		0
	E_aa_effects            int            `db:"e_aa_effects"`            //	int(11) unsigned	NO		0
	E_percent_to_aa         int            `db:"e_percent_to_aa"`         //	int(11) unsigned	NO		0
	E_expended_aa_spent     int            `db:"e_expended_aa_spent"`     //	int(11) unsigned	NO		0
	Aa_points_spent_old     int            `db:"aa_points_spent_old"`     //	int(11) unsigned	NO		0
	Aa_points_old           int            `db:"aa_points_old"`           //	int(11) unsigned	NO		0
	E_last_invsnapshot      int            `db:"e_last_invsnapshot"`      //	int(11) unsigned	NO		0
	Build_data              sql.NullString `db:"build_data"`              //	varchar(39)	YES		NULL
	Session                 sql.NullString `db:"session"`                 //	varchar(32)	YES		NULL
	Session_timeout         mysql.NullTime `db:"session_timeout"`         //	timestamp	NO		CURRENT_TIMESTAMP
}
