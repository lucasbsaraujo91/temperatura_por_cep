package entity

import "errors"

type Address struct {
	ZipCode      string
	Street       string
	Neighborhood string
	City         string
	State        string
}

func NewConsultZipCode(zipCode string, street string, neighborhood string, city string, state string) (*Address, error) {
	address := &Address{
		ZipCode:      zipCode,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}

	err := address.IsValid()
	if err != nil {
		return nil, err
	}

	return address, nil
}

func (a *Address) IsValid() error {
	if len(a.ZipCode) != 8 {
		return errors.New("invalid zipcode")
	}

	return nil
}
