package structs

import (
	"time"
)

/*
		k1.Laskunro,
    DATE(s1.Maksupvm) Maksupvm,
    k1.Kayntimaksu + k1.Sumtoimi + k1.Sumtarvike + k1.Sumtesti + k1.Sumhoito VATFULL,
    k1.Sumrehu VATREDUCED1,
    k1.Sumlaake VATREDUCED2,
    Sum(Case When s1.Kredit = 9 Then s1.Summa Else 0 End) Lahjakorttiosto,
    k1.Kayntimaksu + k1.Sumtoimi + k1.Sumtarvike + k1.Sumtesti + k1.Sumhoito + k1.Sumrehu + k1.Sumlaake + Sum(Case When s1.Kredit = 9 Then s1.Summa Else 0 End) BILLEDTOTAL,
    Sum(Case When s1.Debet = 2 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 2 Then s1.Summa Else 0 End) Kateinen,
    Sum(Case When s1.Debet = 3 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 3 Then s1.Summa Else 0 End) Pkortti,
    Sum(Case When s1.Debet = 4 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 4 Then s1.Summa Else 0 End) Lkortti,
    Sum(Case When s1.Debet = 5 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 5 Then s1.Summa Else 0 End) Tilisiirto,
    Sum(Case When s1.Debet = 6 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 6 Then s1.Summa Else 0 End) Reskontra,
    Sum(Case When s1.Debet = 7 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 7 Then s1.Summa Else 0 End) Tapiola,
    Sum(Case When s1.Debet = 8 Then s1.Summa Else 0 End)-Sum(Case When s1.Kredit = 8 Then s1.Summa Else 0 End) ResursBank,
    Sum(Case When s1.Debet = 9 Then s1.Summa Else 0 End) Lahjakorttimaksu,
    Sum(Case When (s1.Debet in (2,3,4,5,6,7,8) and s1.Kredit <> 7) Then s1.Summa Else 0 End)-Sum(Case When (s1.Kredit in (2,3,4,5,6,7,8) and s1.Debet <> 7) Then s1.Summa Else 0 End) Maksettu_yht
*/
type DailyIncomeReportItem struct {
	BillNumber       int       `db:"Laskunro" json:"billNumber"`
	PaymentDate      time.Time `db:"Maksupvm" json:"paymentDate"`
	VATFull          float64   `db:"VATFULL" json:"vatFull"`
	VATReduced1      float64   `db:"VATREDUCED1" json:"vatReduced1"`
	VATReduced2      float64   `db:"VATREDUCED2" json:"vatReduced2"`
	GiftCardPurchase float64   `db:"Lahjakorttiosto" json:"giftCardPurchase"`
	Cash             float64   `db:"Kateinen" json:"cash"`
	BilledTotal      float64   `db:"BILLEDTOTAL" json:"billedTotal"`
	DebitCard        float64   `db:"Pkortti" json:"debitCard"`
	CreditCard       float64   `db:"Lkortti" json:"creditCard"`
	Mash             float64   `db:"Mash" json:"mash"`
	Transfer         float64   `db:"Tilisiirto" json:"transfer"`
	Ledger           float64   `db:"Reskontra" json:"ledger"`
	Tapiola          float64   `db:"Tapiola" json:"tapiola"`
	ResursBank       float64   `db:"ResursBank" json:"resursBank"`
	GiftCardPayment  float64   `db:"Lahjakorttimaksu" json:"giftCardPayment"`
	PaidTotal        float64   `db:"Maksettu_yht" json:"paidTotal"`
}

/*
		DATE(sub.Maksupvm) Maksupvm,
    Sum(sub.VATFULL) VATFULL,
    Sum(sub.VATREDUCED1) VATREDUCED1,
    Sum(sub.VATREDUCED2) VATREDUCED2,
    Sum(sub.Lahjakorttiosto) Lahjakorttiosto,
    Sum(sub.BILLEDTOTAL) BILLEDTOTAL,
    Sum(sub.Kateinen) Kateinen,
    Sum(sub.Pkortti) Pkortti,
    Sum(sub.Lkortti) Lkortti,
    Sum(sub.Tilisiirto) Tilisiirto,
    Sum(sub.Reskontra) Reskontra,
    Sum(sub.Tapiola) Tapiola,
    Sum(sub.ResursBank) ResursBank,
    Sum(sub.Lahjakorttimaksu) Lahjakorttimaksu,
    Sum(sub.Maksettu_yht) Maksettu_yht
*/
type MonthlyIncomeReportItem struct {
	PaymentDate      time.Time `db:"Maksupvm" json:"paymentDate"`
	VATFull          float64   `db:"VATFULL" json:"vatFull"`
	VATReduced2      float64   `db:"VATREDUCED2" json:"vatReduced2"`
	VATReduced1      float64   `db:"VATREDUCED1" json:"vatReduced1"`
	PaidTotal        float64   `db:"Maksettu_yht" json:"paidTotal"`
	GiftCardPayment  float64   `db:"Lahjakorttimaksu" json:"giftCardPayment"`
	ResursBank       float64   `db:"ResursBank" json:"resursBank"`
	Tapiola          float64   `db:"Tapiola" json:"tapiola"`
	Ledger           float64   `db:"Reskontra" json:"ledger"`
	Transfer         float64   `db:"Tilisiirto" json:"transfer"`
	Cash             float64   `db:"Kateinen" json:"cash"`
	BilledTotal      float64   `db:"BILLEDTOTAL" json:"billedTotal"`
	GiftCardPurchase float64   `db:"Lahjakorttiosto" json:"giftCardPurchase"`
	CreditCard       float64   `db:"Lkortti" json:"creditCard"`
	DebitCard        float64   `db:"Pkortti" json:"debitCard"`
	Mash             float64   `db:"Mash" json:"mash"`
}

/*
DATE(s1.Maksupvm) Maksupvm,
    DATE(s1.Kirjauspvm) Kirjauspvm,
    s1.Laskunro Laskunro,
    SUM(CASE WHEN s1.Debet = 2 THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit = 2 THEN s1.Summa ELSE 0 END) Kateinen,
    SUM(CASE WHEN s1.Debet = 3 THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit = 3 THEN s1.Summa ELSE 0 END) Pkortti,
    SUM(CASE WHEN s1.Debet = 4 THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit = 4 THEN s1.Summa ELSE 0 END) Lkortti,
    SUM(CASE WHEN s1.Debet = 5 THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit = 5 THEN s1.Summa ELSE 0 END) Tilisiirto,
    SUM(CASE WHEN s1.Debet = 9 THEN s1.Summa ELSE 0 END) Lahjakorttimaksu,
    SUM(CASE WHEN s1.Debet = 10 THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit = 10 THEN s1.Summa ELSE 0 END) Luottotappio,
	SUM(CASE WHEN s1.Debet in (2,3,4,5,9,10) THEN s1.Summa ELSE 0 END)-Sum(CASE WHEN s1.Kredit in (2,3,4,5,9,10) THEN s1.Summa ELSE 0 END) Yhteensa
*/
type PaidReceivables struct {
	PaymentDate     time.Time `db:"Maksupvm" json:"paymentDate"`
	EntryDate       time.Time `db:"Kirjauspvm" json:"entryDate"`
	BillNumber      int       `db:"Laskunro" json:"billNumber"`
	Cash            float64   `db:"Kateinen" json:"cash"`
	DebitCard       float64   `db:"Pkortti" json:"debitCard"`
	CreditCard      float64   `db:"Lkortti" json:"creditCard"`
	Mash            float64   `db:"Mash" json:"mash"`
	Transfer        float64   `db:"Tilisiirto" json:"transfer"`
	GiftCardPayment float64   `db:"Lahjakorttimaksu" json:"giftCardPayment"`
	CreditLoss      float64   `db:"Luottotappio" json:"creditLoss"`
	Total           float64   `db:"Yhteensa" json:"total"`
}

/*
   MONTH(ARVOPVM) MONTH,
   ROUND(SUM(KAYNTIMAKSU)/1.24,2) Klinikkamaksut,
   ROUND(SUM(SUMTOIMI)/1.24,2) Toimenpiteet,
   ROUND(SUM(SUMTESTI)/1.24,2) Testit,
   ROUND(SUM(SUMHOITO)/1.24,2) Hoitotyö,
   ROUND(SUM(SUMTARVIKE)/1.24,2) Tarvikkeet,
   ROUND(SUM(SUMREHU)/1.14,2) Rehut,
   ROUND(SUM(SUMLAAKE)/1.1,2) Lääkkeet,
   ROUND(SUM(KAYNTIMAKSU)/1.24,2)+ROUND(SUM(SUMLAAKE)/1.1,2)+ROUND(SUM(SUMREHU)/1.14,2)+ROUND(SUM(SUMTOIMI)/1.24,2)+ROUND(SUM(SUMHOITO)/1.24,2)+ROUND(SUM(SUMTESTI)/1.24,2)+ROUND(SUM(SUMTARVIKE)/1.24,2) Yhteensä
*/
type AnnualSalesLeft struct {
	Month            int     `db:"MONTH"`
	ClinicPayment    float64 `db:"Klinikkamaksut"`
	Procedures       float64 `db:"Toimenpiteet"`
	Tests            float64 `db:"Testit"`
	Nursing          float64 `db:"Hoitotyö"`
	Supplies         float64 `db:"Tarvikkeet"`
	Feeds            float64 `db:"Rehut"`
	Medicines        float64 `db:"Lääkkeet"`
	Total            float64 `db:"Yhteensa"`
	ClinicPaymentVAT float64 `db:"Klinikkamaksut_ALV"`
	ProceduresVAT    float64 `db:"Toimenpiteet_ALV"`
	TestsVAT         float64 `db:"Testit_ALV"`
	NursingVAT       float64 `db:"Hoitotyö_ALV"`
	SuppliesVAT      float64 `db:"Tarvikkeet_ALV"`
	FeedsVAT         float64 `db:"Rehut_ALV"`
	MedicinesVAT     float64 `db:"Lääkkeet_ALV"`
	TotalVAT         float64 `db:"Yhteensa_ALV"`
}

type DailySales struct {
	DateOfVisit  time.Time `db:"Kayntipaiva" json:"dateOfVisit"`
	TimeOfVisit  string    `db:"Kayntiaika" json:"timeOfVisit"`
	Veterinarian string    `db:"Elainlaakari" json:"veterinarian"`
	Invoice      float64   `db:"Lasku" json:"invoice"`
	Customer     string    `db:"Asiakas" json:"customer"`
	ClinicFee    float64   `db:"Klinikkamaksu" json:"clinicFee"`
	ClinicFeeVAT float64   `db:"Klinikkamaksu_ALV" json:"clinicFeeVat"`
	Measures     float64   `db:"Toimenpiteet" json:"measures"`
	MeasuresVAT  float64   `db:"Toimenpiteet_ALV" json:"measuresVat"`
	Tests        float64   `db:"Testit" json:"tests"`
	TestsVAT     float64   `db:"Testit_ALV" json:"testsVat"`
	Nursing      float64   `db:"Hoitotyo" json:"nursing"`
	NursingVAT   float64   `db:"Hoitotyo_ALV" json:"nursingVat"`
	Supplies     float64   `db:"Tarvikkeet" json:"supplies"`
	SuppliesVAT  float64   `db:"Tarvikkeet_ALV" json:"suppliesVat"`
	Feeds        float64   `db:"Rehut" json:"feeds"`
	FeedsVAT     float64   `db:"Rehut_ALV" json:"feedsVat"`
	Drugs        float64   `db:"Laakkeet" json:"drugs"`
	DrugsVAT     float64   `db:"Laakkeet_ALV" json:"drugsVat"`
	Total        float64   `db:"Yhteensa" json:"total"`
	TotalVAT     float64   `db:"Yhteensa_ALV" json:"totalVat"`
}

type AnnualSalesRight struct {
	Month               int     `db:"Kuukausi"`
	Visits              int     `db:"Käynnit"`
	AveragePrice        float64 `db:"Myynti_per_käynti"`
	AveragePriceVAT     float64 `db:"Myynti_per_käynti_ALV"`
	VisitsPerWeekday    float64 `db:"Käyntejä_per_arkipäivä"`
	VisitsPerNonWeekday float64 `db:"Käyntejä_per_ei_arkipäivä"`
	Workdays            int     `db:"Työpäiviä"`
	Weekdays            int     `db:"Arkipäiviä"`
	NonWeekdays         int     `db:"Ei_arkipäiviä"`
}

type AnnualSalesItem struct {
	Month         int     `json:"month"`
	ClinicPayment float64 `json:"clinicPayment"`
	Procedures    float64 `json:"procedures"`
	Tests         float64 `json:"tests"`
	Nursing       float64 `json:"nursing"`
	Supplies      float64 `json:"supplies"`
	Feeds         float64 `json:"feeds"`
	Medicines     float64 `json:"medicines"`
	Total         float64 `json:"total"`

	ClinicPaymentVAT float64 `json:"clinicPaymentVat"`
	ProceduresVAT    float64 `json:"proceduresVat"`
	TestsVAT         float64 `json:"testsVat"`
	NursingVAT       float64 `json:"nursingVat"`
	SuppliesVAT      float64 `json:"suppliesVat"`
	FeedsVAT         float64 `json:"feedsVat"`
	MedicinesVAT     float64 `json:"medicinesVat"`
	TotalVAT         float64 `json:"totalVat"`
	Visits           int     `json:"visits"`
	AveragePrice     float64 `json:"averagePrice"`
	AveragePriceVAT  float64 `json:"averagePriceVat"`

	Workdays         int     `json:"workdays"`
	Weekdays         int     `json:"weekdays"`
	Holidays         int     `json:"holidays"`
	VisitsPerWeekday float64 `json:"visitsPerWeekday"`
	VisitsPerHoliday float64 `json:"visitsPerHoliday"`
	XRayStatements   int     `json:"xrayStatements"`
}

type RangeTotalItems struct {
	PaymentCash            float64 `db:"Kateinen" json:"paymentcash"`
	PaymentDebitCard       float64 `db:"Pkortti" json:"paymentDebitCard"`
	PaymentCreditCard      float64 `db:"Lkortti" json:"paymentCreditCard"`
	PaymentMash            float64 `db:"Mash" json:"paymentMash"`
	PaymentTransfer        float64 `db:"Tilisiirto" json:"paymentTransfer"`
	PaymentGiftCardPayment float64 `db:"Lahjakorttimaksu" json:"paymentGiftCardPayment"`
	PaymentCreditLoss      float64 `db:"Luottotappio" json:"paymentCreditLoss"`
	PaymentTotal           float64 `db:"Yhteensa" json:"paymentTotal"`
	Cash                   float64 `db:"Total_Kateinen" json:"cash"`
	DebitCard              float64 `db:"Total_Pkortti" json:"debitCard"`
	CreditCard             float64 `db:"Total_Lkortti" json:"creditCard"`
	Mash                   float64 `db:"Total_Mash" json:"mash"`
	Transfer               float64 `db:"Total_Tilisiirto" json:"transfer"`
	GiftCardPayment        float64 `db:"Total_Lahjakorttimaksu" json:"giftCardPayment"`
	CreditLoss             float64 `db:"Total_Luottotappio" json:"creditLoss"`
	Total                  float64 `db:"Total_Yhteensa" json:"total"`
}

type MedicineItem struct {
	DateOfVisit  time.Time `db:"Kayntipaiva" json:"dateOfVisit"`
	Invoice      float64   `db:"Lasku" json:"invoice"`
	CustomerName string    `db:"Asiakas" json:"customerName"`
	AnimalName   string    `db:"Elaimen_nimi" json:"animalName"`
	Medicine     string    `db:"Laakkeen_nimi" json:"medicine"`
	Amount       float64   `db:"Maara" json:"amount"`
	Unit         string    `db:"Yksikko" json:"unit"`
}

type AnnualSalesXray struct {
	Month int `db:"Kuukausi"`
	Count int `db:"Lkm"`
}

type VATResults struct {
	VatGeneral  float64 `db:"VAT_GENERAL"`
	VatMedicine float64 `db:"VAT_MEDICINE"`
	VatFeeds    float64 `db:"VAT_FEEDS"`
}
