package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"

	"demo-gorm-gen/internal/pkg/gu"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/infrastructure/database/gormgen",
		ModelPkgPath:  "internal/infrastructure/database/gormgen/model",
		Mode:          gen.WithQueryInterface | gen.WithDefaultQuery,
		FieldSignable: true,
	})

	db := gu.Must(gorm.Open(rawsql.New(rawsql.Config{
		DriverName: "mysql",
		FilePath: []string{
			"internal/infrastructure/database/sql",
		},
	})))

	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	// Generate the code
	g.Execute()
}
