package database

import (
	"database/sql"
	"log"
)

var SupabaseDB *sql.DB

func ConnectDatabase() {
	var err error
	// Connect to the Supabase database
	SupabaseDB, err = sql.Open("postgres", "postgres://postgres:5HPHCU-$rcuQu2_@db.czgqvzsctxzzgbmsjzjy.supabase.co:6543/postgres")
	if err != nil {
		log.Fatal(err)
	}
	//defer SupabaseDB.Close()
}
