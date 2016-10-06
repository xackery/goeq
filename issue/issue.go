package issue

import (
	//"database/sql"
	"time"
)

type Issue struct {
	Id                    int       `db:"id"`               //	int(11) unsigned	NO	PRI	NULL	auto_increment
	Github_issue_id       int       `db:"github_issue_id"`  //	int(10) unsigned	NO		0
	Is_in_progress        int       `db:"is_in_progress"`   //	tinyint(1) unsigned	NO		0
	Is_fixed              int       `db:"is_fixed"`         //	tinyint(1) unsigned	NO		0
	Is_deleted            int       `db:"is_deleted"`       //	tinyint(1) unsigned	NO		0
	My_name               string    `db:"my_name"`          //	varchar(64)	YES		NULL
	My_class              string    `db:"my_class"`         //	varchar(64)	YES		NULL
	My_account_id         int       `db:"my_account_id"`    //	int(10) unsigned	YES		NULL
	My_character_id       int       `db:"my_character_id"`  //	int(10) unsigned	YES		NULL
	My_zone_id            int       `db:"my_zone_id"`       //	int(10) unsigned	YES		NULL
	My_zone_name          string    `db:"my_zone_name"`     //	varchar(64)	YES		NULL
	My_x                  float64   `db:"my_x"`             //	float	YES		NULL
	My_y                  float64   `db:"my_y"`             //	float	YES		NULL
	My_z                  float64   `db:"my_z"`             //	float	YES		NULL
	Message               string    `db:"message"`          //	varchar(512)	YES		NULL
	Create_date           time.Time `db:"create_date"`      //	timestamp	NO		CURRENT_TIMESTAMP
	Last_modified         time.Time `db:"last_modified"`    //	timestamp	YES		CURRENT_TIMESTAMP	on update CURRENT_TIMESTAMP
	Tar_name              string    `db:"tar_name"`         //	varchar(64)	YES		NULL
	Tar_is_npc            int       `db:"tar_is_npc"`       //	tinyint(1)	YES		NULL
	Tar_is_client         int       `db:"tar_is_client"`    //	tinyint(1)	YES		NULL
	Tar_account_id        int       `db:"tar_account_id"`   //	int(10)	YES		NULL
	Tar_character_id      int       `db:"tar_character_id"` //	int(10)	YES		NULL
	Tar_npc_type_id       int       `db:"tar_npc_type_id"`  //	int(10)	YES		NULL
	Tar_npc_spawngroup_id int       `db:"tar_npc_spawngroup_id"`
	Item_id               int       `db:"item_id"`
	Item_name             string    `db:"item_name"`
	Client                string    `db:"client"`
}
