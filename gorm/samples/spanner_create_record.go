package samples

// [START spanner_create_record]

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"io"
	"time"

	_ "github.com/googleapis/go-sql-spanner"

	spannergorm "github.com/rahul2393/go-spanner-orm/gorm"
)

// CreateRecord created a record in DB
// TODO: this fails in Spanner with error= converting argument $3 type: spanner: code = "InvalidArgument", desc = "unsupported value type: {0001-01-01 00:00:00 +0000 UTC false}"
func CreateRecord(ctx context.Context, w io.Writer, dsn string) error {
	// dsn := "projects/my-project/instances/my-instance/databases/my-database"
	db, err := gorm.Open(spannergorm.New(spannergorm.Config{
		DriverName: "spanner",
		DSN:        dsn,
	}), &gorm.Config{PrepareStmt: true})
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

// [END spanner_create_record]
