package main

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "os"
)

var conn *pgx.Conn

func connectDB() {
    var err error
    conn, err = pgx.Connect(context.Background(), "<connection_string>")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    fmt.Println("Connected to the database.")
}

func closeDB() {
    conn.Close(context.Background())
}

// Example function to query the database
func getClientByID(clientID int) {
    var clientName string
    err := conn.QueryRow(context.Background(), "SELECT client_name FROM clients WHERE client_id=$1", clientID).Scan(&clientName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
        return
    }
    fmt.Println("Client Name:", clientName)
}

// Example function to insert a new client
func insertClient(id int, name string, balance float64) {
    _, err := conn.Exec(context.Background(), "INSERT INTO clients (client_id, client_name, account_balance) VALUES ($1, $2, $3)", id, name, balance)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
        return
    }
    fmt.Println("Client inserted successfully.")
}

func main() {
    connectDB()
    defer closeDB()

    // Example usage
    getClientByID(1)
    insertClient(4, "Jane Doe", 1000.50)
}