package account

import (
	"time"
)

type Account struct {
	Id             int       `db:"id"`
	Name           string    `db:"name"`
	CharName       string    `db:"charname"`
	SharedPlat     string    `db:"sharedplat"`
	Password       string    `db:"password"`
	Status         int       `db:"status"`
	LsAccountID    int       `db:"lsaccount_id"`
	GmSpeed        int       `db:"gmspeed"`
	Revoked        int       `db:"revoked"`
	Karma          int       `db:"karma"`
	MiniLoginIp    string    `db:"minilogin_ip"`
	HideMe         int       `db:"hideme"`
	RulesFlag      int       `db:"rulesflag"`
	SuspendedUntil time.Time `db:"suspendeduntil"`
	TimeCreation   int       `db:"time_creation"`
	Expansion      int       `db:"expansion"`
	BanReason      string    `db:"ban_reason"`
	SuspendReason  string    `db:"suspend_reason"`
}
