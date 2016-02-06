package account

import (
	"time"
)

type AccountIp struct {
	AccountID int       `db:"accid"`
	Ip        string    `db:"ip"`
	Count     int       `db:"count"`
	LastUsed  time.Time `db:"lastused"`
}
