package db

import (
	"context"
	"fmt"
	"os"
  "github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

func InitPgsqlDB() *pgx.Conn {
 godotenv.Load()
 conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
 
 if (err != nil) {
   fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
   os.Exit(1)
 }

 fmt.Println(conn)

 return conn
}
