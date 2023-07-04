package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "iamdpk", "lalitpur", "44600", "1999", "1	"},
		{"2", "sjdpk", "kathmandu", "44400", "1999", "1	"},
	}
	return CustomerRepositoryStub{customer: customers}
}
