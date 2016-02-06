package account

import ()

type AccountFlags struct {
	AccountID int    `db:"p_accid"`
	Flag      string `db:"p_flag"`
	Value     string `db:"p_value"`
}
