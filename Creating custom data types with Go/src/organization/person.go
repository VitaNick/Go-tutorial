package organization

import (
	"errors"
	"fmt"
	"strings"
)

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")

	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type Conflict interface {
	ID() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type europeUnionIdentifier struct {
	id      string
	country string
}

func NewEuropeanUnionIdentifier(id, country string) Citizen {
	return europeUnionIdentifier{
		id:      id,
		country: country,
	}
}

func (eui europeUnionIdentifier) ID() string {
	return eui.id
}

func (eui europeUnionIdentifier) Country() string {
	return fmt.Sprintf("EU: %s", eui.country)
}

type Name struct {
	First string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.First, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	First          string
	last           string
	twitterHandler TwitterHandler
	Citizen
	Conflict
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			First: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.First, p.last)
}

func (p *Person) ID() string {
	return p.Citizen.ID()
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler

	} else if !strings.HasPrefix(string(handler), "@") {

		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler

	return nil
}
