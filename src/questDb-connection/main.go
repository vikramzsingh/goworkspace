package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

var conn *pgx.Conn
var err error

func main() {
	ctx := context.Background()
	conn, _ = pgx.Connect(ctx, "postgresql://admin:quest@localhost:8812/qdb")
	defer conn.Close(ctx)

	// text-based query
	_, err := conn.Exec(ctx, "CREATE TABLE IF NOT EXISTS trades (ts TIMESTAMP, date DATE, name STRING, value INT) timestamp(ts);")
	if err != nil {
		log.Fatalln(err)
	}

	// Prepared statement given the name 'ps1'
	_, err = conn.Prepare(ctx, "ps1", "INSERT INTO trades VALUES($1,$2,$3,$4)")
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 10; i++ {
		// Execute 'ps1' statement with a string and the loop iterator value
		_, err = conn.Exec(ctx, "ps1", time.Now(), time.Now().Round(time.Millisecond), "go prepared statement", i+1)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Read all rows from table
	rows, err := conn.Query(ctx, "SELECT * FROM trades")
	fmt.Println("Reading from trades table:")
	for rows.Next() {
		var name string
		var value int64
		var ts time.Time
		var date time.Time
		err = rows.Scan(&ts, &date, &name, &value)
		fmt.Println(ts, date, name, value)
	}

	err = conn.Close(ctx)
}

//const (
//	host     = "localhost"
//	port     = 8812
//	user     = "admin"
//	password = "quest"
//	dbname   = "qdb"
//)
//
//func main() {
//	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//
//	defer db.Close()
//
//	fmt.Println(connStr)
//	fmt.Println(db)
//	log.Println("Connected")
//}


//func main() {
//	host := "127.0.0.1:9009"
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
//	checkErr(err)
//
//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//	checkErr(err)
//	log.Println("Connected")
//	defer conn.Close()
//}
//
//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

