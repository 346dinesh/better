package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/346dinesh/better/internal"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	loginHandler(c.Writer, c.Request)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve form data
	email := r.FormValue("email")
	password := r.FormValue("password")
	org := r.FormValue("org")

	// Authenticate user in Supabase
	// Replace this with your own authentication logic
	id := authenticateUser(email, password, org)

	if id == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Redirect user to organization subdomain
	http.Redirect(w, r, internal.FrontEndServer+"/dashboard", http.StatusFound)
	//	http.Redirect(w, r, "https://ramaiahleena.fyndbetter.com/dashboard?menu=Manage", 200)
}

func authenticateUser(email, password, org string) *int64 {
	var db *sql.DB

	// Connect to the Supabase database
	db, err := sql.Open("postgres", "postgres://postgres:5HPHCU-$rcuQu2_@db.czgqvzsctxzzgbmsjzjy.supabase.co:6543/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query Supabase to authenticate user and retrieve organization
	// Replace this with your own Supabase query logic

	query := `SELECT id FROM user_credentials WHERE user_mail = $1 AND user_password = $2 AND user_org=$3`
	var id int64
	if db == nil {
		fmt.Println("db empty")
		return nil
	}
	err = db.QueryRow(query, email, password, org).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case when no rows were returned by the query
			fmt.Println("No matching user credentials found.")
		} else {
			// Handle other potential errors
			log.Fatal(err)
		}
		return nil
	}

	// Use the retrieved organization ID as needed
	fmt.Println("Organization ID:", id)

	return &id
}
