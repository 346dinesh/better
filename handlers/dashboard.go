package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/346dinesh/better/database"
	"github.com/gin-gonic/gin"
)

type DashBoardResp struct {
	BillAmount       int64 `json:"bill_amount"`
	TotalQueries     int   `json:"total_queries"`
	CompletedQueries int   `json:"completed_queries"`
}

func DashBaordHandler(c *gin.Context) {
	w := c.Writer

	// Retrieve the user ID from the "Auth" header
	userID := c.Request.Header.Get("Auth")

	resp := GetAllDashBaordRequires(userID)
	if resp == nil {
		http.Error(c.Writer, "no response recieved", http.StatusOK)
	}
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Println("error")
	}
}

type dbRes struct {
	BillAmount sql.NullInt64
	Stage      int
}

func GetAllDashBaordRequires(userId string) *DashBoardResp {
	db := database.SupabaseDB

	// Query Supabase to authenticate user and retrieve organization
	// Replace this with your own Supabase query logic
	if database.SupabaseDB == nil {
		fmt.Println("db empty")
		return nil
	}
	query := `
	SELECT bill_amount,stage
	from user_credentials JOIN patient_records ON user_credentials.hospital_id=patient_records.incoming_hospital_id AND user_credentials.user_id=$1;
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle case when no rows were returned by the query
			fmt.Println("No matching user credentials found.")
		} else {
			// Handle other potential errors
			log.Fatal(err)
			fmt.Println("error from rows")
		}
		return nil
	}
	defer rows.Close()
	var a []dbRes
	for rows.Next() {

		var res dbRes
		if err = rows.Scan(&res.BillAmount, &res.Stage); err != nil {
			log.Fatal("Error scanning row:", err)
		}
		a = append(a, res)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal("Error processing rows:", err)
	}

	totalBill := 0
	totalQueries := 0
	completedQueries := 0
	for _, values := range a {
		if values.BillAmount.Valid {
			totalBill += int(values.BillAmount.Int64)
		}
		if values.Stage >= 3 {
			completedQueries += 1
		}
		totalQueries += 1
	}
	return &DashBoardResp{
		TotalQueries:     totalQueries,
		CompletedQueries: completedQueries,
		BillAmount:       int64(totalBill),
	}
}
