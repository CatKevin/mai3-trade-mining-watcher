package mining

import (
	"github.com/mcdexio/mai3-trade-mining-watcher/database/models"
	"github.com/shopspring/decimal"
)

// Snapshot defines struct to contain information of a user info
type Snapshot struct {
	Trader    string `gorm:"column:trader;primary_key;type:varchar(128);not null" json:"trader"`
	Epoch     int64  `gorm:"column:epoch;primary_key;type:bigint;not null" json:"epoch"`
	Timestamp int64  `gorm:"column:timestamp;primary_key;type:bigint;not null" json:"timestamp"`

	InitFee       decimal.Decimal `gorm:"column:init_fee;type:decimal(38,18);not null" json:"init_fee"`
	AccFee        decimal.Decimal `gorm:"column:acc_fee;type:decimal(38,18);not null" json:"acc_fee"`
	AccPosValue   decimal.Decimal `gorm:"column:acc_pos_value;type:decimal(38,18);not null" json:"acc_pos_value"`
	CurPosValue   decimal.Decimal `gorm:"column:cur_pos_value;type:decimal(38,18);not null" json:"cur_pos_value"`
	AccStakeScore decimal.Decimal `gorm:"column:acc_stake_score;type:decimal(38,18);not null" json:"acc_stake_score"`
	CurStakeScore decimal.Decimal `gorm:"column:cur_stake_score;type:decimal(38,18);not null" json:"cur_stake_score"`
	Score         decimal.Decimal `gorm:"column:score;type:decimal(38,18);not null" json:"score"`

	models.Base
}