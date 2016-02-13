package grid

import ()

type GridEntries struct {
	Gridid  int     `db:"gridid"`  //	int(10)	NO	PRI	0
	Zoneid  int     `db:"zoneid"`  //	int(10)	NO	PRI	0
	Number  int     `db:"number"`  //	int(10)	NO	PRI	0
	X       float64 `db:"x"`       //	float	NO		0
	Y       float64 `db:"y"`       //	float	NO		0
	Z       float64 `db:"z"`       //	float	NO		0
	Heading float64 `db:"heading"` //	float	NO		0
	Pause   int     `db:"pause"`   //	int(10)	NO		0
}
