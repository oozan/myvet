package structs

import (
	"database/sql"
	"time"

	"gopkg.in/guregu/null.v4"
)

// JSONErr is an abstraction of an error message as a JSON-object.
type JSONErr struct {
	Error string `json:"error"`
}

//WorkShift is a DAO and DTO for the klinikkaohjelma_kehitys.ttyovuoro table
type WorkShift struct {
	WorkShiftID      int       `json:"workShiftId" db:"Id"`
	EndTime          time.Time `json:"endTime" db:"LopetusAjankohta"`
	StartTime        time.Time `json:"startTime" db:"AloitusAjankohta"`
	Type             null.Int  `json:"type" db:"Tyyppi"`
	EmployeeID       int       `json:"employeeId" db:"TyontekijaId"`
	WorkShiftModelID null.Int  `json:"workShiftModelId" db:"TyovuoroMalliId"`
}

//Species is a DAO and DTO for the klinikkaohjelma_kehitys.tlaji table
type Species struct {
	SpeciesID int         `json:"speciesId" db:"Lajinro"`
	Name      null.String `json:"name" db:"Nimi"`
	Deleted   null.String `json:"deleted" db:"Deleted"`
	Jarj      null.Int    `json:"jarj" db:"Jarj"`
	IDEXXID   null.String `json:"IDEXXId" db:"IDEXX_ID"`
}

//Customer struct contains all fields for DB and JSON DTO.
type Customer struct {
	CustomerID   int         `json:"customerId" db:"Asnro"`
	FirstName    null.String `json:"firstName" db:"Etunimi"`
	Address      null.String `json:"address" db:"Osoite"`
	PostalCode   null.String `json:"postalCode" db:"Postinumero"`
	PostalCity   null.String `json:"postalCity" db:"Postitoimipaikka"`
	PhoneNumber1 null.String `json:"phoneNumber1" db:"Puhnro1"`
	PhoneNumber2 null.String `json:"phoneNumber2" db:"Puhnro2"`
	Note         null.String `json:"note" db:"Huom"`
	Type         null.String `json:"type" db:"Tyyppi"`
	Company      null.String `json:"company" db:"Yrnimi"`
	Language     null.String `json:"language" db:"Kieli"`
	About        null.String `json:"about" db:"Tietoja"`
	ProductName  null.String `json:"productName" db:"Tulosnimi"`
	Deleted      null.String `json:"deleted" db:"Deleted"`
	LastName     null.String `json:"lastName" db:"Sukunimi"`
	Email        null.String `json:"email" db:"Sposti"`
	SMS          null.String `json:"sms" db:"SMS"`
	Distance     null.Int    `json:"distance" db:"Etaisyys"`
	Lat          null.Float  `json:"lat" db:"Lat"`
	Lng          null.Float  `json:"lng" db:"Lng"`
	Debt         null.Float  `json:"debt" db:"Asvelka"`
}

type CustomerAnimalsDTO struct {
	CustomerID int         `db:"customerId"`
	AnimalIDs  null.String `db:"animalIds"`
}
type CustomerAnimalIDsDTO struct {
	CustomerID int   `db:"customerId" json:"customerId"`
	AnimalIDs  []int `db:"animalIds" json:"animalIds"`
}

//Animal struct contains all fields for DB and JSON DTO.
type Animal struct {
	AnimalID          int         `json:"animalId" db:"Elainnro"`
	CustomerID        int         `json:"customerId" db:"Asnro"`
	Name              null.String `json:"name" db:"Nimi"`
	SpeciesID         null.Int    `json:"speciesId" db:"Laji"`
	BreedID           null.Int    `json:"breedId" db:"Rotu"`
	BirthDate         null.Time   `json:"birthDate" db:"Spaiva"`
	Gender            null.Int    `json:"gender" db:"Spuoli"`
	Tattoo            null.String `json:"tattoo" db:"Tatuointi"`
	Microchip         null.String `json:"microchip" db:"Mikrosiru"`
	Sterilized        null.Int    `json:"sterilized" db:"Steriloitu"`
	SterilizationDate null.Time   `json:"sterilizationDate" db:"Sterilaika"`
	DateOfDeath       null.Time   `json:"dateOfDeath" db:"Kuolpaiva"`
	Note              null.String `json:"huom" db:"Huom"`
	BirthYear         null.String `json:"birthYear" db:"Svuosi"`
	RegisterNumber    null.String `json:"registerNumber" db:"Reknro"`
	Seurok            null.Time   `json:"seurok" db:"Seurok"`
	Seuroktul         null.Int    `json:"seuroktul" db:"Seuroktul"`
	FormalName        null.String `json:"formalName" db:"Virnimi"`
	Deleted           null.String `json:"deleted" db:"Deleted"`
	Weight            null.Float  `json:"weight" db:"Paino"`
	Kutsu             null.Int    `json:"kutsu" db:"Kutsu"`
	RontgenNro        null.String `json:"rontgennro" db:"RontgenNro"`
	Insurance         null.String `json:"insurance" db:"Vakuutus"`
	About             null.String `json:"about" db:"Yleista"`
	PortaaliRotu      null.String `json:"portaalirotu" db:"PortaaliRotu"`
	PortaaliLaji      null.String `json:"portaaliLaji" db:"PortaaliLaji"`
	Type              null.Int    `json:"tyyppi" db:"Tyyppi"`
	// WeightHistory     []WeightHistoryEntry `json:"weightHistory"`
	WeightHistory []WeightHistoryEntry `json:"weightHistory"`
}

// AnimalAppointment is a DAO and DTO for the klinikkaohjelma_kehitys.telainkaynti table
type AnimalAppointment struct {
	AAID          int          `json:"aaId" db:"Elkanro"`
	AnimalID      int          `json:"animalId" db:"Elainnro"`
	AppointmentID int          `json:"appointmentId" db:"Kayntinro"`
	Anamnesis     null.String  `json:"anamnesis" db:"Anamneesi"`
	Advisor       null.String  `json:"advisor" db:"Hohje"`
	Status        null.String  `json:"status" db:"Status"`
	Species       null.Int     `json:"species" db:"Laji"`
	FinancialInfo null.String  `json:"financialInfo" db:"Tulostieto"`
	Deleted       null.String  `json:"deleted" db:"Deleted"`
	IdexxCount    null.Int     `json:"idexxCount" db:"idexx_count"`
	Medicines     []AAMedicine `json:"medicines"`
}

// AAMedicine is a DAO and DTO for the klinikkaohjelma_kehitys.tkaynti_laake table
// it's to be joined to AnimalAppointment
type AAMedicine struct {
	AAMedicineID           int         `json:"aaMedicineId" db:"Avain"`
	AnimalAppointmentID    null.Int    `json:"animalAppointmentId" db:"Elkanro"`
	MedicineID             null.Int    `json:"medicineId" db:"Tlaake_laakenro"`
	Price                  null.Float  `json:"price" db:"Hinta"`
	Unit                   null.Int    `json:"unit" db:"Annosyks"`
	Amount                 null.Float  `json:"amount" db:"Maara"`
	Name                   null.String `json:"name" db:"Nimi"`
	AppointmentProcedureID null.Int    `json:"appointmentProcedureId" db:"Kayntitoiminro"`
	Deleted                null.String `json:"deleted" db:"Deleted"`
	AppointmentTestID      null.Int    `json:"appointmentTestId" db:"Kayntitestinro"`
	Assign                 null.Int    `json:"assign" db:"Luovutettu"`
}
type Medicine struct {
	MedicineID      int         `json:"medicineId" db:"Laakenro"`
	MedicineGroupID int         `json:"medicineGroupId" db:"Laakeryhnro"`
	Name            null.String `json:"name" db:"Nimi"`
	UnitPrice       null.Float  `json:"unitPrice" db:"Ykshinta"`
	Unit            null.Int    `json:"unit" db:"Annosyks"`
	PackageSize     null.String `json:"packageSize" db:"Pakkauskoko"`
	Dosage          null.String `json:"dosage" db:"Annosteluohje"`
	Strength        null.String `json:"strength" db:"Vahvuus"`
	Deleted         null.Int    `json:"deleted" db:"Deleted"`
	Index           null.Int    `json:"index" db:"Jarj"`
	Barcode         null.String `json:"barcode" db:"Viivakoodi"`
	Name2           null.String `json:"name2" db:"Nimi2"`
	Permit          null.Int    `json:"permit" db:"Erityislupa"`
	PurchasePrice   null.Float  `json:"purchasePrice" db:"Ostohinta"`
	Sale            null.Int    `json:"sale" db:"Myyntiosa"`
	Wholesale       null.Int    `json:"wholesale" db:"Tukkuri"`
	Oriola          null.String `json:"oriola" db:"Oriola"`
	Prevett         null.String `json:"prevett" db:"Prevett"`
	VNR             null.String `json:"vnr" db:"VNR"`
	MSI             null.String `json:"msi" db:"MSI"`
	Prevett2        null.String `json:"prevett2" db:"Prevett2"`
}

type Unit struct {
	UnitID  int         `json:"unitId" db:"Avain"`
	Name    null.String `json:"name" db:"Nimi"`
	Deleted null.Int    `json:"deleted" db:"Deleted"`
	Index   null.Int    `json:"index" db:"Jarj"`
}

// Appointment struct contains all fields for DB and JSON DTO - klinikkaohjelma_kehitys.tkaynti table
type Appointment struct {
	AppointmentID          int                 `json:"appointmentId" db:"Kayntinro"`
	Date                   null.Time           `json:"date" db:"Kayntipvm"`
	CustomerID             int                 `json:"customerId" db:"Asnro"`
	Purpose                null.String         `json:"purpose" db:"Tulosyy"`
	Type                   null.String         `json:"type" db:"Tyyppi"`
	Deleted                null.String         `json:"deleted" db:"Deleted"`
	EmployeeID             null.Int            `json:"employeeId" db:"Elainlaakari"`
	Sumtoimi               null.Float          `json:"sumtoimi" db:"Sumtoimi"`
	Sumtesti               null.Float          `json:"sumtesti" db:"Sumtesti"`
	Sumtarvike             null.Float          `json:"sumtarvike" db:"Sumtarvike"`
	Sumrehu                null.Float          `json:"sumrehu" db:"Sumrehu"`
	Sumlaake               null.Float          `json:"sumlaake" db:"Sumlaake"`
	VisitFee               null.Float          `json:"visitFee" db:"Kayntimaksu"`
	Hohje                  null.String         `json:"hohje" db:"Hohje"`
	State                  null.Int            `json:"state" db:"Tila"`
	InvoiceDate            null.Time           `json:"invoiceDate" db:"Laskupvm"`
	DueDate                null.Time           `json:"dueDate" db:"Erapvm"`
	Debt                   null.Float          `json:"debt" db:"Sumvelka"`
	Resurs                 null.Float          `json:"resurs" db:"Sumresurs"`
	Ledger                 null.Float          `json:"ledger" db:"Sumreskontra"`
	LedgerReason           null.Int            `json:"ledgerReason" db:"Reskontrasyy"`
	LedgerReasonInfo       null.String         `json:"ledgerReasonInfo" db:"Reskontrasyy_selite"`
	PriceDate              null.Time           `json:"priceDate" db:"Arvopvm"`
	Reference              null.String         `json:"reference" db:"Viite"`
	BillNumber             null.Int            `json:"billNumber" db:"Laskunro"`
	Sumhoito               null.Float          `json:"sumhoito" db:"Sumhoito"`
	Sent                   null.String         `json:"sent" db:"Lahetetty"`
	Tapiola                null.Float          `json:"tapiola" db:"Sumtapiola"`
	File                   null.String         `json:"file" db:"Tiedosto"`
	MailLah                null.String         `json:"mailLah" db:"Mail_lah"`
	MailLah2               null.String         `json:"mailLah2" db:"Mail_lah2"`
	StartTime              null.String         `json:"startTime" db:"Hoitoaika"`
	EndTime                null.String         `json:"endTime" db:"Loppuaika"`
	OperationEndTime       null.String         `json:"operationEndTime" db:"Toimi_loppu"`
	OperationStartTime     null.String         `json:"operationStartTime" db:"Toimi_alku"`
	AdditionalInformation  null.String         `json:"additionalInformation" db:"MuutaHuomioitavaa"`
	AnimalInformation      null.String         `json:"animalInformation" db:"ElainTiedot"`
	VarausMontaElainta     null.String         `json:"varausMontaElainta" db:"VarausMontaElainta"`
	PurposeCategory        null.Int            `json:"purposeCategory" db:"TulosyyKategoriaId"`
	VarausElainLaji        null.Int            `json:"varausElainLaji" db:"VarausElainLaji"`
	VarausElainIka         null.Int            `json:"varausElainIka" db:"VarausElainIka"`
	VarausElainNimi        null.String         `json:"varausElainNimi" db:"VarausElainNimi"`
	VarausElainRotu        null.Int            `json:"varausElainRotu" db:"VarausElainRotu"`
	TapiolaComment         null.String         `json:"tapiolaComment" db:"Tapiola_comment"`
	TapiolaSivu            null.Int            `json:"tapiolaSivu" db:"TapiolaSivu"`
	TapiolaCommunication   null.String         `json:"tapiolaCommunication" db:"Tapiola_communication"`
	GiftCard               null.Int            `json:"giftCard" db:"Lahjakortti"`
	GiftCardValue          null.Float          `json:"giftCardValue" db:"Lahjakorttiarvo"`
	UsedGiftCard           null.Int            `json:"usedGiftCard" db:"Lahjakorttikaytetty"`
	DischargeTime          null.Time           `json:"dischargeTime" db:"Kotiutusaika"`
	DischargeInfo          null.String         `json:"dischargeInfo" db:"Kotiutusohje"`
	CustomerChatRoomNumber null.Int            `json:"customerChatRoomNumber" db:"AsiakasChatHuonenro"`
	Creator                null.Int            `json:"creator" db:"Luoja"`
	All                    null.Int            `json:"all" db:"Kaikki"`
	AnimalAppointments     []AnimalAppointment `json:"animalAppointments"`
	//	Reservation            ReservationDTO      `json:"reservation"`
}

type SearchAppointmentDTO struct {
	AppointmentID          int                 `json:"appointmentId" db:"Kayntinro"`
	Date                   null.Time           `json:"date" db:"Kayntipvm"`
	CustomerID             int                 `json:"customerId" db:"Asnro"`
	Purpose                null.String         `json:"purpose" db:"Tulosyy"`
	Type                   null.String         `json:"type" db:"Tyyppi"`
	EmployeeID             null.Int            `json:"employeeId" db:"Elainlaakari"`
	Sumtoimi               null.Float          `json:"sumtoimi" db:"Sumtoimi"`
	Sumtesti               null.Float          `json:"sumtesti" db:"Sumtesti"`
	Sumtarvike             null.Float          `json:"sumtarvike" db:"Sumtarvike"`
	Sumrehu                null.Float          `json:"sumrehu" db:"Sumrehu"`
	Sumlaake               null.Float          `json:"sumlaake" db:"Sumlaake"`
	VisitFee               null.Float          `json:"visitFee" db:"Kayntimaksu"`
	Hohje                  null.String         `json:"hohje" db:"Hohje"`
	State                  null.Int            `json:"state" db:"Tila"`
	InvoiceDate            null.Time           `json:"invoiceDate" db:"Laskupvm"`
	DueDate                null.Time           `json:"dueDate" db:"Erapvm"`
	Ledger                 null.Float          `json:"ledger" db:"Sumreskontra"`
	LedgerReason           null.Int            `json:"ledgerReason" db:"Reskontrasyy"`
	LedgerReasonInfo       null.String         `json:"ledgerReasonInfo" db:"Reskontrasyy_selite"`
	PriceDate              null.Time           `json:"priceDate" db:"Arvopvm"`
	Reference              null.String         `json:"reference" db:"Viite"`
	BillNumber             null.Int            `json:"billNumber" db:"Laskunro"`
	Sumhoito               null.Float          `json:"sumhoito" db:"Sumhoito"`
	Sent                   null.String         `json:"sent" db:"Lahetetty"`
	Tapiola                null.Float          `json:"tapiola" db:"Sumtapiola"`
	File                   null.String         `json:"file" db:"Tiedosto"`
	MailLah                null.String         `json:"mailLah" db:"Mail_lah"`
	MailLah2               null.String         `json:"mailLah2" db:"Mail_lah2"`
	StartTime              null.String         `json:"startTime" db:"Hoitoaika"`
	EndTime                null.String         `json:"endTime" db:"Loppuaika"`
	OperationEndTime       null.String         `json:"operationEndTime" db:"Toimi_loppu"`
	OperationStartTime     null.String         `json:"operationStartTime" db:"Toimi_alku"`
	AdditionalInformation  null.String         `json:"additionalInformation" db:"MuutaHuomioitavaa"`
	AnimalInformation      null.String         `json:"animalInformation" db:"ElainTiedot"`
	VarausMontaElainta     null.String         `json:"varausMontaElainta" db:"VarausMontaElainta"`
	PurposeCategory        null.Int            `json:"purposeCategory" db:"TulosyyKategoriaId"`
	VarausElainLaji        null.Int            `json:"varausElainLaji" db:"VarausElainLaji"`
	VarausElainIka         null.Int            `json:"varausElainIka" db:"VarausElainIka"`
	VarausElainNimi        null.String         `json:"varausElainNimi" db:"VarausElainNimi"`
	VarausElainRotu        null.Int            `json:"varausElainRotu" db:"VarausElainRotu"`
	TapiolaComment         null.String         `json:"tapiolaComment" db:"Tapiola_comment"`
	TapiolaSivu            null.Int            `json:"tapiolaSivu" db:"TapiolaSivu"`
	TapiolaCommunication   null.String         `json:"tapiolaCommunication" db:"Tapiola_communication"`
	GiftCard               null.Int            `json:"giftCard" db:"Lahjakortti"`
	GiftCardValue          null.Float          `json:"giftCardValue" db:"Lahjakorttiarvo"`
	UsedGiftCard           null.Int            `json:"usedGiftCard" db:"Lahjakorttikaytetty"`
	DischargeTime          null.Time           `json:"dischargeTime" db:"Kotiutusaika"`
	DischargeInfo          null.String         `json:"dischargeInfo" db:"Kotiutusohje"`
	CustomerChatRoomNumber null.Int            `json:"customerChatRoomNumber" db:"AsiakasChatHuonenro"`
	Creator                null.Int            `json:"creator" db:"Luoja"`
	All                    BitBool             `json:"all" db:"Kaikki"`
	AnimalAppointments     []AnimalAppointment `json:"animalAppointments,omitempty"`
	//	Reservation            ReservationDTO      `json:"reservation"`
}
type SearchAppointmentDAO struct {
	AppointmentID          int         `json:"appointmentId" db:"Kayntinro"`
	Date                   null.Time   `json:"date" db:"Kayntipvm"`
	CustomerID             int         `json:"customerId" db:"Asnro"`
	Purpose                null.String `json:"purpose" db:"Tulosyy"`
	Type                   null.String `json:"type" db:"Tyyppi"`
	Deleted                null.String `json:"deleted" db:"Deleted"`
	EmployeeID             null.Int    `json:"employeeId" db:"Elainlaakari"`
	Sumtoimi               null.Float  `json:"sumtoimi" db:"Sumtoimi"`
	Sumtesti               null.Float  `json:"sumtesti" db:"Sumtesti"`
	Sumtarvike             null.Float  `json:"sumtarvike" db:"Sumtarvike"`
	Sumrehu                null.Float  `json:"sumrehu" db:"Sumrehu"`
	Sumlaake               null.Float  `json:"sumlaake" db:"Sumlaake"`
	VisitFee               null.Float  `json:"visitFee" db:"Kayntimaksu"`
	Hohje                  null.String `json:"hohje" db:"Hohje"`
	State                  null.Int    `json:"state" db:"Tila"`
	InvoiceDate            null.Time   `json:"invoiceDate" db:"Laskupvm"`
	DueDate                null.Time   `json:"dueDate" db:"Erapvm"`
	Ledger                 null.Float  `json:"ledger" db:"Sumreskontra"`
	LedgerReason           null.Int    `json:"ledgerReason" db:"Reskontrasyy"`
	LedgerReasonInfo       null.String `json:"ledgerReasonInfo" db:"Reskontrasyy_selite"`
	PriceDate              null.Time   `json:"priceDate" db:"Arvopvm"`
	Reference              null.String `json:"reference" db:"Viite"`
	BillNumber             null.Int    `json:"billNumber" db:"Laskunro"`
	Sumhoito               null.Float  `json:"sumhoito" db:"Sumhoito"`
	Sent                   null.String `json:"sent" db:"Lahetetty"`
	Tapiola                null.Float  `json:"tapiola" db:"Sumtapiola"`
	File                   null.String `json:"file" db:"Tiedosto"`
	MailLah                null.String `json:"mailLah" db:"Mail_lah"`
	MailLah2               null.String `json:"mailLah2" db:"Mail_lah2"`
	StartTime              null.String `json:"startTime" db:"Hoitoaika"`
	EndTime                null.String `json:"endTime" db:"Loppuaika"`
	OperationEndTime       null.String `json:"operationEndTime" db:"Toimi_loppu"`
	OperationStartTime     null.String `json:"operationStartTime" db:"Toimi_alku"`
	AdditionalInformation  null.String `json:"additionalInformation" db:"MuutaHuomioitavaa"`
	AnimalInformation      null.String `json:"animalInformation" db:"ElainTiedot"`
	VarausMontaElainta     null.String `json:"varausMontaElainta" db:"VarausMontaElainta"`
	PurposeCategory        null.Int    `json:"purposeCategory" db:"TulosyyKategoriaId"`
	VarausElainLaji        null.Int    `json:"varausElainLaji" db:"VarausElainLaji"`
	VarausElainIka         null.Int    `json:"varausElainIka" db:"VarausElainIka"`
	VarausElainNimi        null.String `json:"varausElainNimi" db:"VarausElainNimi"`
	VarausElainRotu        null.Int    `json:"varausElainRotu" db:"VarausElainRotu"`
	TapiolaComment         null.String `json:"tapiolaComment" db:"Tapiola_comment"`
	TapiolaSivu            null.Int    `json:"tapiolaSivu" db:"TapiolaSivu"`
	TapiolaCommunication   null.String `json:"tapiolaCommunication" db:"Tapiola_communication"`
	GiftCard               null.Int    `json:"giftCard" db:"Lahjakortti"`
	GiftCardValue          null.Float  `json:"giftCardValue" db:"Lahjakorttiarvo"`
	UsedGiftCard           null.Int    `json:"usedGiftCard" db:"Lahjakorttikaytetty"`
	DischargeTime          null.Time   `json:"dischargeTime" db:"Kotiutusaika"`
	DischargeInfo          null.String `json:"dischargeInfo" db:"Kotiutusohje"`
	CustomerChatRoomNumber null.Int    `json:"customerChatRoomNumber" db:"AsiakasChatHuonenro"`
	Creator                null.Int    `json:"creator" db:"Luoja"`
	AppointmentAnimalIDs   null.String `db:"KayntiElaimet"`
	// Reservation Reservation `json:"reservation"`
}

type CustomersAnimalsSearchResult struct {
}

//Payment is a DAO and DTO for the klinikkaohjelma_kehitys.tsuoritus table
type Payment struct {
	PaymentID   int         `json:"paymentId" db:"Avain"`
	BillNumber  int         `json:"billNumber" db:"Laskunro"`
	PaymentDate time.Time   `json:"paymentDate" db:"Maksupvm"`
	EntryDate   time.Time   `json:"entryDate" db:"Kirjauspvm"`
	Cash        null.Float  `json:"cash" db:"Kateinen"`
	DebitCard   null.Float  `json:"debitCard" db:"Pkortti"`
	CreditCard  null.Float  `json:"creditCard" db:"Lkortti"`
	Transfer    null.Float  `json:"transfer" db:"Tilisiirto"`
	EventType   null.String `json:"eventType" db:"Tapahtumalaji"`
	Deleted     null.String `json:"deleted" db:"Deleted"`
	Ledger      null.Float  `json:"ledger" db:"Reskontra"`
	CreditLoss  null.Float  `json:"creditLoss" db:"Luottotappio"`
	Tapiola     null.Float  `json:"tapiola" db:"Tapiola"`
}

/*
| Avain Laskunro  Maksutapahtuma Lahjakorttinro Maksupvm Kirjauspvm Kredit Debet Summa Deleted | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
|        | int(10) unsigned | NO   |     | 0       |                |
|  | int(3)           | NO   |     | NULL    |                |
|  | int(5)           | YES  |     | NULL    |                |
|        | datetime         | NO   |     | NULL    |                |
|      | datetime         | NO   |     | NULL    |                |
|          | int(11)          | YES  |     | NULL    |                |
|           | int(11)          | YES  |     | NULL    |                |
|           | decimal(15,2)    | YES  |     | NULL    |                |
|         | char(1)          | YES  |     | NULL    |                |
+----------------+------------------+------+-----+---------+----------------+
*/
type Payment2 struct {
	PaymentID   int        `db:"Avain" json:"paymentId"`
	BillNumber  int        `db:"Laskunro" json:"billNumber"`
	Transaction int        `db:"Maksutapahtuma" json:"transaction"`
	GiftCardID  null.Int   `db:"Lahjakorttinro" json:"giftCardId"`
	PaymentDate time.Time  `db:"Maksupvm" json:"paymentDate"`
	EntryDate   time.Time  `db:"Kirjauspvm" json:"entryDate"`
	Kredit      null.Int   `db:"Kredit" json:"kredit"`
	Debet       null.Int   `db:"Debet" json:"debet"`
	Sum         null.Float `db:"Summa" json:"sum"`
	Deleted     bool       `db:"Deleted" json:"deleted"`
}

type TapiolaReply struct {
	Payment       Payment2 `json:"payment"`
	Comment       string   `json:"comment"`
	AppointmentId int      `json:"appointmentId"`
}

/*
| Avain       | int(11)     | NO   | PRI | NULL    | auto_increment |
| Tilinumero  | int(11)     | YES  |     | NULL    |                |
| Nimi        | varchar(80) | YES  |     | NULL    |                |
| Kreditlabel | varchar(80) | YES  |     | NULL    |                |
| Debetlabel  | varchar(80) | YES  |     | NULL    |                |
| Deleted     | int(11)     | YES  |     | NULL    |                |
+-------------+-------------+------+-----+---------+----------------+
*/
type AccountMap struct {
	AccountMapID int    `db:"Avain" json:"accountMapId"`
	AccountID    int    `db:"Tilinumero" json:"accountId"`
	Name         string `db:"Nimi" json:"name"`
	KreditLabel  string `db:"Kreditlabel" json:"kreditlabel"`
	DebetLabel   string `db:"Debetlabel" json:"debetlabel"`
	Deleted      bool   `db:"Deleted" json:"deleted"`
}

//Receivable struct contains all fields for DB and JSON DTO.
type Receivable struct {
	Laskunumero      int     `json:"billNumber" db:"Laskunro"`
	Kayntipaiva      string  `json:"date" db:"Kayntipaiva"`
	Viite            string  `json:"refNumber" db:"Viite"`
	Ledger           float64 `json:"ledger" db:"Reskontra"`
	Tapiola          float64 `json:"tapiola" db:"Tapiola"`
	VAT              float32 `json:"VAT" db:"VAT"`
	VATReduced1      float32 `json:"VATReduced1" db:"VATReduced1"`
	VATReduced2      float32 `json:"VATReduced2" db:"VATReduced2"`
	Sukunimi         string  `json:"lastName" db:"Sukunimi"`
	Etunimi          string  `json:"firstName" db:"Etunimi"`
	Osoite           string  `json:"address" db:"Osoite"`
	Postinumero      string  `json:"postalCode" db:"Postinumero"`
	Postitoimipaikka string  `json:"postalCity" db:"Postitoimipaikka"`
	Reskontrasyy     string  `json:"ledgerReason" db:"Reskontrasyy"`
}

//User struct contains all fields for DB and JSON DTO
//for the klinikkaohjelma_kehitys.tportaalikayttaja table.
type User struct {
	UserName    string      `json:"userName" db:"Kayttajatunnus"`
	Password    string      `json:"password" db:"Salasana"`
	Deleted     int         `json:"deleted" db:"Deleted"`
	Enabled     int         `json:"enabled" db:"Kaytossa"`
	Expires     null.Int    `json:"expires" db:"Vanhentuu"`
	Lang        null.String `json:"lang" db:"Kieli"`
	Locked      int         `json:"locked" db:"Lukittu"`
	Permissions null.Int    `json:"permissions" db:"Oikeudet"`
	Role        null.Int    `json:"role" db:"Rooli"`
	Pin         null.String `json:"pin" db:"Pin"`
	Picnum      null.Int    `json:"picNum" db:"Kuvanro"`
}

// UserDTO is a stripped down version of User with only attributes safe and useful on the frontend.
type UserDTO struct {
	UserName    string      `json:"userName"`
	Enabled     int         `json:"enabled"`
	Expires     null.Int    `json:"expires"`
	Lang        null.String `json:"lang"`
	Locked      int         `json:"locked"`
	Permissions null.Int    `json:"permissions"`
	Role        null.Int    `json:"role"`
	Picnum      null.Int    `json:"picNum"`
}

//MessageType is a DAO and DTO for the klinikkaohjelma_kehitys.tviestityyppi table
type MessageType struct {
	Tunniste     int         `json:"id" db:"Tunniste"`
	Prioriteetit null.String `json:"defaultPriorities" db:"Prioriteetit"`
	Deleted      null.Bool   `json:"deleted" db:"Deleted"`
	Kuvaus       null.String `json:"description" db:"Kuvaus"`
	Nimi         null.String `json:"name" db:"Nimi"`
	Tyyppi       null.Int    `json:"type" db:"Tyyppi"`
}

//MessageTemplateLocalized is a DAO and DTO for the klinikkaohjelma_kehitys.tviestimalli_kaannos table
type MessageTemplateLocalized struct {
	Tunniste            int         `json:"id" db:"Tunniste"`
	Kieli               string      `json:"locale" db:"Kieli"`
	Sisalto             null.String `json:"content" db:"Sisalto"`
	Nimi                null.String `json:"name" db:"Nimi"`
	SmsSisalto          null.String `json:"smsContent" db:"SmsSisalto"`
	Otsikko             null.String `json:"subject" db:"Otsikko"`
	ViestimalliTunniste null.Int    `json:"template" db:"ViestimalliTunniste"`
}

//MessageTemplate is a DAO and DTO for the klinikkaohjelma_kehitys.tviestimalli table
type MessageTemplate struct {
	Tunniste     int       `json:"id" db:"Tunniste"`
	Deleted      null.Bool `json:"deleted" db:"Deleted"`
	Lahettaja    null.Bool `json:"from" db:"Lahettaja"`
	RequiredType null.Int  `json:"requiredType" db:"requiredType"`
	Viestityyppi null.Int  `json:"type" db:"Viestityyppi"`
}

//MessageChannel is a DAO and DTO for the klinikkaohjelma_kehitys.tviestikanava table
type MessageChannel struct {
	Tunniste int         `json:"id" db:"Tunniste"`
	Deleted  null.Bool   `json:"deleted" db:"Deleted"`
	Nimi     null.String `json:"name" db:"Nimi"`
}

//Employee is a DAO and DTO for the klinikkaohjelma_kehitys.ttyontekija table
type Employee struct {
	EmployeeID int `json:"employeeId" db:"Avain"`
	//Name            sql.NullString `json:"name" db:"Nimi"`
	Name            null.String
	Type            null.String `json:"type" db:"Tyyppi"`
	Username        null.String `json:"username" db:"Kayttajatunnus"`
	Password        null.String `json:"password" db:"Salasana"`
	VetNumber       null.String `json:"vetNumber" db:"Ellnro"`
	Active          null.String `json:"active" db:"Active"`
	Index           null.String `json:"index" db:"Jarj"`
	ReservationLock int         `json:"reservationLock" db:"AjanvarausLukko"`
}

type EmployeeDTO struct {
	EmployeeID      int         `json:"employeeId" db:"Avain"`
	Name            null.String `json:"name" db:"Nimi"`
	Type            null.String `json:"type" db:"Tyyppi"`
	Username        null.String `json:"username" db:"Kayttajatunnus"`
	VetNumber       null.String `json:"vetNumber" db:"Ellnro"`
	Active          null.String `json:"active" db:"Active"`
	Index           null.String `json:"index" db:"Jarj"`
	ReservationLock int         `json:"reservationLock" db:"AjanvarausLukko"`
}

//ContactDetailChannel is a DAO and DTO for the klinikkaohjelma_kehitys.tyhteystieto_kanava table
type ContactDetailChannel struct {
	Tunniste      int      `json:"id" db:"Tunniste"`
	Active        int      `json:"active" db:"Aktiivinen"`
	Deleted       int      `json:"deleted" db:"Deleted"`
	ContactDetail null.Int `json:"contactDetail" db:"Yhteystieto"`
	Channel       null.Int `json:"messageChannel" db:"Kanava"`
}

//ContactDetail is a DAO and DTO for the klinikkaohjelma_kehitys.tyhteystieto table
type ContactDetail struct {
	ContactDetailID       int                    `json:"contactDetailId" db:"Tunniste"`
	Deleted               int                    `json:"deleted" db:"Deleted"`
	Detail                null.String            `json:"detail" db:"Yhteystieto"`
	Index                 null.Int               `json:"index" db:"Jarj"`
	Notes                 null.String            `json:"notes" db:"Lisatiedot"`
	Type                  null.Int               `json:"type" db:"Tyyppi"`
	CustomerID            null.Int               `json:"customerId" db:"Asnro"`
	ContactDetailChannels []ContactDetailChannel `json:"contactDetailChannels"`
}

//Breed is a DAO and DTO for the klinikkaohjelma_kehitys.trotu table
type Breed struct {
	BreedID   int         `json:"breedId" db:"Rotunro"`
	SpeciesID int         `json:"speciesId" db:"Lajinro"`
	Name      null.String `json:"name" db:"Nimi"`
	Deleted   null.String `json:"deleted" db:"Deleted"`
	Index     null.Int    `json:"index" db:"Jarj"`
}

// PurposeDTO is a replacement for the complicated system underneath.
type PurposeDTO struct {
	PurposeCategoryID int    `json:"purposeCategoryId" db:"Tunniste"`
	Name              string `json:"name" db:"Otsikko"`
	Duration          int    `json:"duration" db:"AikaArvioMinuutit"`
}

// PurposeGroup is a DAO and DTO for the klinikkaohjelma_kehitys.tkaynti_syy_ryhma table
type PurposeGroup struct {
	PurposeGroupID int       `json:"purposeGroupId" db:"Tunniste"`
	Deleted        null.Bool `json:"deleted" db:"Deleted"`
	Jarj           null.Int  `json:"index" db:"Jarj"`
	Parent         null.Int  `json:"parent" db:"Vanhempi"`
}

// PurposeGroupLocalized is a DAO and DTO for the klinikkaohjelma_kehitys.tkaynti_syy_ryhma_kaannos table
type PurposeGroupLocalized struct {
	PurposeGroupLocalizedID int    `json:"purposeGroupLocalizedId" db:"Tunniste"`
	Locale                  string `json:"locale" db:"Kieli"`
	Name                    string `json:"name  " db:"Nimi"`
	PurposeGroupID          int    `json:"group  " db:"KayntiSyyRyhmaTunniste"`
}

// PurposeCategory is a DAO and DTO for the klinikkaohjelma_kehitys.tkaynti_syy_kategoria table
type PurposeCategory struct {
	PurposeCategoryID      int       `json:"purposeCategoryId" db:"Tunniste"`
	Deleted                null.Bool `json:"deleted" db:"Deleted"`
	EstimateMinutes        null.Int  `json:"timeEstimateMinutes" db:"AikaArvioMinuutit"`
	Index                  null.Int  `json:"index" db:"Jarj"`
	PurposeGroupID         null.Int  `json:"purposeGroupId" db:"Ryhma"`
	RequiresAdditionalInfo null.Bool `json:"additionalInfoRequired" db:"VaatiiLisatiedot"`
	SlotIntervarMinutes    null.Int  `json:"slotIntervalMinutes" db:"VarausvaliMinuutit"`
	DefaultSpecies         null.Int  `json:"defaultSpecies" db:"Oletuslaji"`
	EmailTemplate          null.Int  `json:"emailTemplate" db:"Sahkopostipohja"`
	Varattavissa           null.Bool `json:"reservable" db:"Ajanvaraus"`
	MessageTemplate        null.Int  `json:"messageTemplate" db:"Viestipohja"`
}

// PurposeCategoryLocalized is a DAO and DTO for the klinikkaohjelma_kehitys.tkaynti_syy_kategoria_kaannos table
type PurposeCategoryLocalized struct {
	PurposeCategoryLocalizedID int    `json:"purposeCategoryLocalizedId" db:"Tunniste"`
	Locale                     string `json:"locale" db:"Kieli"`
	Kuvaus                     string `json:"description" db:"Kuvaus"`
	Topic                      string `json:"topic" db:"Otsikko"`
	PurposeCategoryID          string `json:"category" db:"KayntiSyyKategoriaTunniste"`
}

// NonMatchingSum is DAO and DTO for a specialized SQL query looking for erroneous data.
type NonMatchingSum struct {
	BillNumber int     `json:"billNumber" db:"Laskunro"`
	Billed     float64 `json:"billed" db:"Laskutettu"`
	Paid       float64 `json:"paid" db:"Maksettu"`
}

//TODO: change type to some enum type?
type LogEvent struct {
	LogID     int         `json:"logID" db:"logID"`
	CreatedAt time.Time   `json:"createdAt" db:"created_at"`
	LogType   string      `json:"type" db:"log_type"`
	Source    null.String `json:"source" db:"source"`
	Target    null.String `json:"target" db:"target"`
	LogData   null.String `json:"payload" db:"log_payload"`
}

type TemporaryPage struct {
	TemporaryPageID int    `db:"Avain" json:"temporaryPageId"`
	Target          string `db:"Kohde" json:"target"`
	AccessKey       string `db:"Tunnus" json:"accessKey"`
	Type            string `db:"Tyyppi" json:"type"`
	// ACTIVE(1), USED(2), DELETED(3), ACTION_REQUIRED(4), UNKNOWN(99);
	State      int          `db:"Tila" json:"state"`
	Expiration sql.NullTime `db:"Eraantyminen" json:"expiration"`
}

// DTO representing the data that is POSTed when adding a tapiola comment to an appointment
type TapiolaCommunicationDTO struct {
	AppointmentID int    `json:"appointmentId"`
	Answer        string `json:"answer"`
}

// DAO/DTO represeting a weigh history entry in tpainohistoria table
type WeightHistoryEntry struct {
	ID              int         `db:"Avain" json:"weightHistoryId"`
	AnimalID        int         `db:"Elainnro" json:"animalId"`
	Deleted         null.String `db:"Deleted" json:"Deleted"`
	Weight          float64     `db:"Paino" json:"weight"`
	MeasurementDate null.Time   `db:"Mittauspvm" json:"measurementDate"`
}
