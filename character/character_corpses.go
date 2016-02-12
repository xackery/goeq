package character

import (
	"database/sql"
	"time"
)

type CharacterCorpses struct {
	Id               sql.NullInt64 `db:"id"`               //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Charid           int           `db:"charid"`           //	int(11) unsigned	NO		0
	Charname         string        `db:"charname"`         //	varchar(64)	NO
	Zone_id          int           `db:"zone_id"`          //	smallint(5)	NO	MUL	0
	Instance_id      int           `db:"instance_id"`      //	smallint(5) unsigned	NO	MUL	0
	X                float64       `db:"x"`                //	float	NO		0
	Y                float64       `db:"y"`                //	float	NO		0
	Z                float64       `db:"z"`                //	float	NO		0
	Heading          float64       `db:"heading"`          //	float	NO		0
	Time_of_death    time.Time     `db:"time_of_death"`    //	datetime	NO		0000-00-00 00:00:00
	Is_rezzed        int           `db:"is_rezzed"`        //	tinyint(3) unsigned	YES		0
	Is_buried        int           `db:"is_buried"`        //	tinyint(3)	NO		0
	Was_at_graveyard int           `db:"was_at_graveyard"` //	tinyint(3)	NO		0
	Is_locked        int           `db:"is_locked"`        //	tinyint(11)	YES		0
	Exp              int           `db:"exp"`              //	int(11) unsigned	YES		0
	Size             int           `db:"size"`             //	int(11) unsigned	YES		0
	Level            int           `db:"level"`            //	int(11) unsigned	YES		0
	Race             int           `db:"race"`             //	int(11) unsigned	YES		0
	Gender           int           `db:"gender"`           //	int(11) unsigned	YES		0
	Class            int           `db:"class"`            //	int(11) unsigned	YES		0
	Deity            int           `db:"deity"`            //	int(11) unsigned	YES		0
	Texture          int           `db:"texture"`          //	int(11) unsigned	YES		0
	Helm_texture     int           `db:"helm_texture"`     //	int(11) unsigned	YES		0
	Copper           int           `db:"copper"`           //	int(11) unsigned	YES		0
	Silver           int           `db:"silver"`           //	int(11) unsigned	YES		0
	Gold             int           `db:"gold"`             //	int(11) unsigned	YES		0
	Platinum         int           `db:"platinum"`         //	int(11) unsigned	YES		0
	Hair_color       int           `db:"hair_color"`       //	int(11) unsigned	YES		0
	Beard_color      int           `db:"beard_color"`      //	int(11) unsigned	YES		0
	Eye_color_1      int           `db:"eye_color_1"`      //	int(11) unsigned	YES		0
	Eye_color_2      int           `db:"eye_color_2"`      //	int(11) unsigned	YES		0
	Hair_style       int           `db:"hair_style"`       //	int(11) unsigned	YES		0
	Face             int           `db:"face"`             //	int(11) unsigned	YES		0
	Beard            int           `db:"beard"`            //	int(11) unsigned	YES		0
	Drakkin_heritage int           `db:"drakkin_heritage"` //	int(11) unsigned	YES		0
	Drakkin_tattoo   int           `db:"drakkin_tattoo"`   //	int(11) unsigned	YES		0
	Drakkin_details  int           `db:"drakkin_details"`  //	int(11) unsigned	YES		0
	Wc_1             int           `db:"wc_1"`             //	int(11) unsigned	YES		0
	Wc_2             int           `db:"wc_2"`             //	int(11) unsigned	YES		0
	Wc_3             int           `db:"wc_3"`             //	int(11) unsigned	YES		0
	Wc_4             int           `db:"wc_4"`             //	int(11) unsigned	YES		0
	Wc_5             int           `db:"wc_5"`             //	int(11) unsigned	YES		0
	Wc_6             int           `db:"wc_6"`             //	int(11) unsigned	YES		0
	Wc_7             int           `db:"wc_7"`             //	int(11) unsigned	YES		0
	Wc_8             int           `db:"wc_8"`             //	int(11) unsigned	YES		0
	Wc_9             int           `db:"wc_9"`             //	int(11) unsigned	YES		0
}
