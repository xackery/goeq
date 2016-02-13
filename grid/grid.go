package grid

import ()

type Grid struct {
	Id     int `db:"id"`     //	int(10)	NO	PRI	0
	Zoneid int `db:"zoneid"` //	int(10)	NO	PRI	0
	Type   int `db:"type"`   //	int(10)	NO		0
	Type2  int `db:"type2"`  //	int(10)	NO		0
}
