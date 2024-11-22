package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// CustomDate wraps time.Time for custom JSON and database handling
type CustomDate struct {
	time.Time
}

// UnmarshalJSON parses JSON date in `yyyy-MM-dd` format
func (c *CustomDate) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	t, err := time.Parse("2006-01-02", str) // Match `yyyy-MM-dd` format
	if err != nil {
		return err
	}
	c.Time = t
	return nil
}

// MarshalJSON outputs JSON date in `yyyy-MM-dd` format
func (c CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format("2006-01-02"))), nil
}

// Value implements the `driver.Valuer` interface for database storage
func (c CustomDate) Value() (driver.Value, error) {
	return c.Time.Format("2006-01-02"), nil
}

// Scan implements the `sql.Scanner` interface for database retrieval
func (c *CustomDate) Scan(value interface{}) error {
	if value == nil {
		*c = CustomDate{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		c.Time = v
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		c.Time = t
	default:
		return fmt.Errorf("cannot scan type %T into CustomDate", value)
	}
	return nil
}

type Loan struct {
	LoanID          uint       `gorm:"primaryKey;autoIncrement" json:"loan_id"`
	BorrowerID      uint       `gorm:"not null" json:"borrower_id"`
	PrincipalAmount float64    `gorm:"type:decimal(15,2);not null" json:"principal_amount"`
	InterestRate    float64    `gorm:"type:decimal(5,2);not null" json:"interest_rate"`
	StartDate       CustomDate `gorm:"type:date;not null" json:"start_date"`
	EndDate         CustomDate `gorm:"type:date;not null" json:"end_date"`
	Status          string     `gorm:"type:varchar(50);not null" json:"status"`
}
