package entities

import (
	"encoding/xml"
	"time"
)

type AccountResults struct {
	XMLName  xml.Name  `xml:"EntityResults"`
	Accounts []Account `xml:"Entity"`
}

type Account struct {
	XMLName     xml.Name `xml:"Entity"`
	ID          int      `xml:"id"`
	AccountName string
}

type ResourceResults struct {
	XMLName   xml.Name   `xml:"EntityResults"`
	Resources []Resource `xml:"Entity"`
}

type Resource struct {
	XMLName    xml.Name `xml:"Entity"`
	ID         int      `xml:"id"`
	ResourceID int
	FirstName  string
	LastName   string
}

type RoleResults struct {
	XMLName xml.Name `xml:"EntityResults"`
	Roles   []Role   `xml:"Entity"`
}

type Role struct {
	XMLName xml.Name `xml:"Entity"`
	ID      int      `xml:"id"`
	RoleID  int
	Name    string
}

type TicketResults struct {
	XMLName xml.Name `xml:"EntityResults"`
	Tickets []Ticket `xml:"Entity"`
}

type Ticket struct {
	XMLName            xml.Name `xml:"Entity"`
	ID                 int      `xml:"id"`
	AccountID          int
	AssignedResourceID int
	TicketNumber       string
	Title              string
	Description        string
	Status             int
}

type TimeEntryResults struct {
	XMLName     xml.Name    `xml:"EntityResults"`
	TimeEntries []TimeEntry `xml:"Entity"`
}

type TimeEntry struct {
	XMLName       xml.Name `xml:"Entity"`
	ID            int      `xml:"id"`
	HoursWorked   float32
	ResourceID    int
	RoleID        int
	ResourceName  string
	RoleName      string
	TicketID      int
	Ticket        *Ticket    `xml:"-"`
	Account       *Account   `xml:"-"`
	StartDateTime customTime `xml:"StartDateTime"`
}

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	const atForm = "2006-01-02T15:04:05"
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(atForm, v)
	if err != nil {
		return nil
	}
	*c = customTime{parse}
	return nil
}
