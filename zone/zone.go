package zone

import (
	"database/sql"
)

type Zone struct {
	Short_name          sql.NullString `db:"short_name"`          //	varchar(32)	YES	MUL	NULL
	Id                  sql.NullInt64  `db:"id"`                  //	int(10)	NO	PRI	NULL	auto_increment
	File_name           sql.NullString `db:"file_name"`           //	varchar(16)	YES		NULL
	Long_name           sql.NullString `db:"long_name"`           //	text	NO		NULL
	Map_file_name       sql.NullString `db:"map_file_name"`       //	varchar(100)	YES		NULL
	Safe_x              float64        `db:"safe_x"`              //	float	NO		0
	Safe_y              float64        `db:"safe_y"`              //	float	NO		0
	Safe_z              float64        `db:"safe_z"`              //	float	NO		0
	Graveyard_id        float64        `db:"graveyard_id"`        //	float	NO		0
	Min_level           int            `db:"min_level"`           //	tinyint(3) unsigned	NO		0
	Min_status          int            `db:"min_status"`          //	tinyint(3) unsigned	NO		0
	Zoneidnumber        int            `db:"zoneidnumber"`        //	int(4)	NO	MUL	0
	Version             int            `db:"version"`             //	tinyint(3) unsigned	NO		0
	Timezone            int            `db:"timezone"`            //	int(5)	NO		0
	Maxclients          int            `db:"maxclients"`          //	int(5)	NO		0
	Ruleset             int            `db:"ruleset"`             //	int(10) unsigned	NO		0
	Note                sql.NullString `db:"note"`                //	varchar(80)	YES		NULL
	Underworld          int            `db:"underworld"`          //	float	NO		0
	Minclip             float64        `db:"minclip"`             //	float	NO		450
	Maxclip             float64        `db:"maxclip"`             //	float	NO		450
	Fog_minclip         float64        `db:"fog_minclip"`         //	float	NO		450
	Fog_maxclip         float64        `db:"fog_maxclip"`         //	float	NO		450
	Fog_blue            int            `db:"fog_blue"`            //	tinyint(3) unsigned	NO		0
	Fog_red             int            `db:"fog_red"`             //	tinyint(3) unsigned	NO		0
	Fog_green           int            `db:"fog_green"`           //	tinyint(3) unsigned	NO		0
	Sky                 int            `db:"sky"`                 //	tinyint(3) unsigned	NO		1
	Ztype               int            `db:"ztype"`               //	tinyint(3) unsigned	NO		1
	Zone_exp_multiplier float64        `db:"zone_exp_multiplier"` //	decimal(6,2)	NO		0.00
	Walkspeed           float64        `db:"walkspeed"`           //	float	NO		0.4
	Time_type           int            `db:"time_type"`           //	tinyint(3) unsigned	NO		2
	Fog_red1            int            `db:"fog_red1"`            //	tinyint(3) unsigned	NO		0
	Fog_green1          int            `db:"fog_green1"`          //	tinyint(3) unsigned	NO		0
	Fog_blue1           int            `db:"fog_blue1"`           //	tinyint(3) unsigned	NO		0
	Fog_minclip1        float64        `db:"fog_minclip1"`        //	float	NO		450
	Fog_maxclip1        float64        `db:"fog_maxclip1"`        //	float	NO		450
	Fog_red2            int            `db:"fog_red2"`            //	tinyint(3) unsigned	NO		0
	Fog_green2          int            `db:"fog_green2"`          //	tinyint(3) unsigned	NO		0
	Fog_blue2           int            `db:"fog_blue2"`           //	tinyint(3) unsigned	NO		0
	Fog_minclip2        float64        `db:"fog_minclip2"`        //	float	NO		450
	Fog_maxclip2        float64        `db:"fog_maxclip2"`        //	float	NO		450
	Fog_red3            int            `db:"fog_red3"`            //	tinyint(3) unsigned	NO		0
	Fog_green3          int            `db:"fog_green3"`          //	tinyint(3) unsigned	NO		0
	Fog_blue3           int            `db:"fog_blue3"`           //	tinyint(3) unsigned	NO		0
	Fog_minclip3        float64        `db:"fog_minclip3"`        //	float	NO		450
	Fog_maxclip3        float64        `db:"fog_maxclip3"`        //	float	NO		450
	Fog_red4            int            `db:"fog_red4"`            //	tinyint(3) unsigned	NO		0
	Fog_green4          int            `db:"fog_green4"`          //	tinyint(3) unsigned	NO		0
	Fog_blue4           int            `db:"fog_blue4"`           //	tinyint(3) unsigned	NO		0
	Fog_minclip4        float64        `db:"fog_minclip4"`        //	float	NO		450
	Fog_maxclip4        float64        `db:"fog_maxclip4"`        //	float	NO		450
	Fog_density         float64        `db:"fog_density"`         //	float	NO		0
	Flag_needed         string         `db:"flag_needed"`         //	varchar(128)	NO
	Canbind             int            `db:"canbind"`             //	tinyint(4)	NO		1
	Cancombat           int            `db:"cancombat"`           //	tinyint(4)	NO		1
	Canlevitate         int            `db:"canlevitate"`         //	tinyint(4)	NO		1
	Castoutdoor         int            `db:"castoutdoor"`         //	tinyint(4)	NO		1
	Hotzone             int            `db:"hotzone"`             //	tinyint(3) unsigned	NO		0
	Insttype            int            `db:"insttype"`            //	tinyint(1) unsigned zerofill	NO		0
	Shutdowndelay       int            `db:"shutdowndelay"`       //	bigint(16) unsigned	NO		5000
	Peqzone             int            `db:"peqzone"`             //	tinyint(4)	NO		1
	Expansion           int            `db:"expansion"`           //	tinyint(3)	NO		0
	Suspendbuffs        int            `db:"suspendbuffs"`        //	tinyint(1) unsigned	NO		0
	Rain_chance1        int            `db:"rain_chance1"`        //	int(4)	NO		0
	Rain_chance2        int            `db:"rain_chance2"`        //	int(4)	NO		0
	Rain_chance3        int            `db:"rain_chance3"`        //	int(4)	NO		0
	Rain_chance4        int            `db:"rain_chance4"`        //	int(4)	NO		0
	Rain_duration1      int            `db:"rain_duration1"`      //	int(4)	NO		0
	Rain_duration2      int            `db:"rain_duration2"`      //	int(4)	NO		0
	Rain_duration3      int            `db:"rain_duration3"`      //	int(4)	NO		0
	Rain_duration4      int            `db:"rain_duration4"`      //	int(4)	NO		0
	Snow_chance1        int            `db:"snow_chance1"`        //	int(4)	NO		0
	Snow_chance2        int            `db:"snow_chance2"`        //	int(4)	NO		0
	Snow_chance3        int            `db:"snow_chance3"`        //	int(4)	NO		0
	Snow_chance4        int            `db:"snow_chance4"`        //	int(4)	NO		0
	Snow_duration1      int            `db:"snow_duration1"`      //	int(4)	NO		0
	Snow_duration2      int            `db:"snow_duration2"`      //	int(4)	NO		0
	Snow_duration3      int            `db:"snow_duration3"`      //	int(4)	NO		0
	Snow_duration4      int            `db:"snow_duration4"`      //	int(4)	NO		0
	Gravity             float64        `db:"gravity"`             //	float	NO		0.4
	Type                int            `db:"type"`                //	int(3)	NO		0
	Skylock             int            `db:"skylock"`             //	tinyint(4)	NO		0
	Levels              int            `db:"levels"`              //  int(4) unsigned NOT NULL DEFAULT '0' COMMENT 'Level range to hunt here',
	Description         string         `db:"description"`         // varchar(128) DEFAULT NULL,
	Sort                int            `db:"sort"`                // tinyint(1) unsigned NOT NULL DEFAULT '0',
}
