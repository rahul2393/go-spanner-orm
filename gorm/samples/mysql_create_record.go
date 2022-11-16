package samples

import (
	"context"
	"fmt"
	"io"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLCreateRecord created a record in DB
func MySQLCreateRecord(ctx context.Context, w io.Writer, dsn string) error {
	// dsn := "user:password@tcp(127.0.0.1:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true})
	if err != nil {
		return err
	}

	currentTime := time.Now()
	singer := Singer{Name: "Jinzhu", Age: 18, Birthday: &currentTime}

	result := db.Create(&singer) // pass pointer of data to Create
	fmt.Fprintf(w, "record(s) updated %v.\n", singer.ID)
	fmt.Fprintf(w, "returns error %v.\n", result.Error)
	fmt.Fprintf(w, "returns inserted records count %v.\n", result.RowsAffected)
	return nil
}
