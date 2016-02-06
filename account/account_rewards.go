package account

import ()

type AccountRewards struct {
	AccountID int `db:"account_id"`
	RewardID  int `db:"reward_id"`
	Amount    int `db:"amount"`
}
