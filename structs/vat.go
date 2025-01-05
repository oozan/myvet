package structs

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type VatPeriod struct {
	ID        int          `db:"Avain" json:"id"`
	VatType   string       `db:"Vat_type" json:"vatType"`
	Vat       float64      `db:"Vat" json:"vat"`
	BeginDate time.Time    `db:"Alkupvm" json:"beginDate"`
	EndDate   null.Time `db:"Loppupvm" json:"endDate"`
}
