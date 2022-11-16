package samples

// [START spanner_migrate_model]

import (
	"context"
	"database/sql"
	"io"
	"time"

	"gorm.io/gorm"

	_ "github.com/googleapis/go-sql-spanner"

	spannergorm "github.com/rahul2393/go-spanner-orm/gorm"
)

// CREATE TABLE singers (
//
//	id INT64,
//	created_at TIMESTAMP,
//	updated_at TIMESTAMP,
//	deleted_at TIMESTAMP,
//	name STRING(MAX),
//	email STRING(MAX),
//	age INT64,
//	birthday TIMESTAMP,
//	member_number STRING(MAX),
//	activated_at TIMESTAMP,
//
// ) PRIMARY KEY(id);
//
// CREATE INDEX idx_singers_deleted_at ON singers(deleted_at);

type Singer struct {
	gorm.Model
	Name         string
	Email        string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
	Age          uint8  `gorm:"<-:update"`          // allow read and update
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	Songs        []string `gorm:"-"` // ignore this field when write and read with struct
}

// MigrateModel validates the GORM struct declaration with Spanner and run the migrations
func MigrateModel(ctx context.Context, w io.Writer, dsn string) error {
	// dsn := "projects/my-project/instances/my-instance/databases/my-database"
	db, err := gorm.Open(spannergorm.New(spannergorm.Config{
		DriverName: "spanner",
		DSN:        dsn,
	}), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return err
	}
	// Automatically create the "songs" table based on the `Account`
	// model.
	return db.AutoMigrate(&Singer{})
}

// [END spanner_migrate_model]
