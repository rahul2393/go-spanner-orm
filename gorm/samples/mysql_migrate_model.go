package samples

// [START spanner_gorm_structs]

import (
	"context"
	"io"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// CREATE TABLE `singers` (
//  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
//  `created_at` datetime(3) DEFAULT NULL,
//  `updated_at` datetime(3) DEFAULT NULL,
//  `deleted_at` datetime(3) DEFAULT NULL,
//  `name` longtext,
//  `email` longtext,
//  `age` tinyint unsigned DEFAULT NULL,
//  `birthday` datetime(3) DEFAULT NULL,
//  `member_number` longtext,
//  `activated_at` datetime(3) DEFAULT NULL,
//  PRIMARY KEY (`id`),
//  KEY `idx_singers_deleted_at` (`deleted_at`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

// MYSQLMigrateModel validates the GORM struct declaration with Spanner and run the migrations
func MYSQLMigrateModel(ctx context.Context, w io.Writer, dsn string) error {
	// dsn := "user:password@tcp(127.0.0.1:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true})
	if err != nil {
		return err
	}
	// Automatically create the "songs" table based on the `Account`
	// model.
	return db.AutoMigrate(&Singer{})
}

// [END spanner_gorm_structs]
