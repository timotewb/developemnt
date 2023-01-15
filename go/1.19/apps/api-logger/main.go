package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "192.168.144.131"
	port     = 5432
	user     = "API"
	password = "api"
	dbname   = "API"
)

func main() {
	// define connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// validate connection arguments
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// make connection
	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB connection successful")
	}

	// create logging table if not exists (monthly)
	t := time.Now()
	year := strconv.Itoa(t.Year())               // type int
	month := fmt.Sprintf("%02d", int(t.Month())) // type time.Month
	code := `CREATE TABLE IF NOT EXISTS logging.api_` + year + month + `(
		host_name varchar NOT NULL,
		job_name varchar NOT NULL,
		pid integer NOT NULL,
		start_dt timestamp without time zone NOT NULL,
		end_dt timestamp without time zone NOT NULL
	);`
	_, err = db.Exec(code)
	if err != nil {
		panic(err)
	}

	// insert record
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s\n", hostname)
	pid := os.Getpid()
	fmt.Printf("PID: %d\n", pid)
	currentTime := time.Now()
	fmt.Println("StartDateTime : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Jobname: job")
	code = fmt.Sprintf(`INSERT INTO logging.api_202301 (host_name, job_name, pid, start_dt, end_dt) VALUES('%s', '%s', %s, '%s', '%s');`, hostname, "jobby", strconv.Itoa(pid), currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))
	_, err = db.Exec(code)
	if err != nil {
		panic(err)
	}

}
