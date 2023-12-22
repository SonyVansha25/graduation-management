package graduation

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Student struct {
	ID             int64  `gorm:"primaryKey"`
	Name           string `gorm:"NOT NULL;type:VARCHAR(100);default:NULL" json:"name"`
	FinalGrade     int    `gorm:"NOT NULL;type:INTEGER;default:NULL;check:final_grade >= 1 AND final_grade <= 100" json:"finalGrade"`
	Passed         bool   `gorm:"NOT NULL" json:"passed"`
	SchoolCohortOf int    `gorm:"NOT NULL;type:INTEGER;default:NULL" json:"schoolCohortOf"`
	CreatedAt      string `gorm:"NOT NULL;type:VARCHAR(80)" json:"createdAt"`
	UpdatedAt      string `gorm:"type:VARCHAR(80)" json:"updatedAt"`
}

func setupDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s TimeZone=Asia/Jakarta", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(new(Student)); err != nil {
		return nil, err
	}

	return db, nil
}
