package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// replace 'user', 'password', 'tcp(your-instance-connection-name)', 'dbname' with your actual values
	dsn := "root:411528Apjy2k18?!@tcp(34.101.43.143)/hackfest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return db, nil
}
