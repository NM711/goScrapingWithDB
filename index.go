package main

import (
	"context"
	"fmt"
	"testScraper/db"
	"testScraper/scrapers"
)

func main() {
  data := scrapers.FetchBooks()
  itemsArr := scrapers.ParseData(data)
  conn := db.InitPgsqlDB()
  fmt.Println(conn)
 // move this function to secondary db func
  for index, item := range itemsArr {
  // %d = int, %s = string
  res, err := conn.Exec(context.Background(), "INSERT INTO books (id, title) VALUES ($1, $2)", index, item)
  if err != nil {
      fmt.Println("ERROR:", err)
    }

    
  rowsAffected := res.RowsAffected()
  fmt.Println("Rows affected:", rowsAffected)
  }
  defer conn.Close(context.Background())
}
