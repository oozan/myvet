package structs

import "gopkg.in/guregu/null.v4"

// DAO/DTO structs for the route search/customersAnimals/{terms}
type CustomersAnimalsSearchDAO struct {
	FirstName null.String `db:"etunimi"`
	LastName  null.String `db:"sukunimi"`
	Animals   null.String `db:"elaimet"` // "," delimited list obtained by using GROUP_CONCAT
}

type CustomersAnimalsSearchDTO struct {
	CustomerID int      `json:"customerId"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	Animals    []string `json:"animals"`
	Debt       float64  `json:"debt"`
}
