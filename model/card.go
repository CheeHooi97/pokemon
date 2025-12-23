package model

import (
	"database/sql/driver"
	"time"
)

type Card struct {
	ID          string    `gorm:"primaryKey;type:varchar(50)" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	Supertype   string    `gorm:"type:varchar(50)" json:"supertype"`
	Subtypes    JSON      `gorm:"type:json" json:"subtypes"`
	HP          string    `gorm:"type:varchar(10)" json:"hp"`
	Types       JSON      `gorm:"type:json" json:"types"`
	EvolvesFrom string    `gorm:"type:varchar(255)" json:"evolvesFrom"`
	Images      JSON      `gorm:"type:json" json:"images"`
	Rarity      string    `gorm:"type:varchar(50)" json:"rarity"`
	Artist      string    `gorm:"type:varchar(255)" json:"artist"`
	SetID       string    `gorm:"type:varchar(50)" json:"setId"`
	Number      string    `gorm:"type:varchar(20)" json:"number"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		// handle string if driver returns string
		sStr, okStr := value.(string)
		if okStr {
			*j = append((*j)[0:0], sStr...)
			return nil
		}
	}
	*j = append((*j)[0:0], s...)
	return nil
}
