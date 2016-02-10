package guild

import ()

type Guild struct {
	Id          int    `db:"id"`          //	int(11)	NO	PRI	NULL	auto_increment
	Name        int    `db:"name"`        //	varchar(32)	NO	UNI
	Leader      int    `db:"leader"`      //	int(11)	NO	UNI	0
	Minstatus   int    `db:"minstatus"`   //	smallint(5)	NO		0
	Motd        string `db:"motd"`        //	text	NO		NULL
	Tribute     int    `db:"tribute"`     //	int(10) unsigned	NO		0
	Motd_setter string `db:"motd_setter"` //	varchar(64)	NO
	Channel     string `db:"channel"`     //	varchar(128)	NO
	Url         string `db:"url"`         //	varchar(512)	NO
}
