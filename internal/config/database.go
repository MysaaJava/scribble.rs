package config

import (
	"strings"
	"strconv"
	"slices"
	"math/rand"
	"time"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/blockloop/scan/v2"
)

var db *sql.DB


func InitializeDatabase(config *Config) {
	connstr := "user='"+config.Database.User + "' password='" + config.Database.Password + "' host='" + config.Database.Host+ "' port='" + config.Database.Port + "' dbname='" + config.Database.Database + "' sslmode='" + config.Database.SslMode+ "'"
	log.Print("Connecting to database with connection string: "+connstr)
	database, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	
	// Trying to ping the database
	pingErr := database.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connection to the database successfully established !")
	db = database
}

func GetAllWords(groupIDs []int) []string {
	
	if len(groupIDs) == 0{
		return nil
	}
	
	// We create the base array of our recursive query
	var groupIDsStr []string
	for _, i := range groupIDs {
		groupIDsStr = append(groupIDsStr, "("+strconv.Itoa(i)+")")
	}
	baseTable := strings.Join(groupIDsStr, ",")
	var statement = `
		WITH RECURSIVE cte AS (
			(SELECT (column1::bigint) AS id FROM (VALUES %s) AS t)
			UNION
			(SELECT wg.id FROM wgroups wg
			JOIN cte ON wg.parentid = cte.id)
		) SELECT str FROM words w JOIN cte ON w.groupid = cte.id;
	`
	fullStatement := strings.Replace(statement, "%s", baseTable, 1)
	log.Println(fullStatement)
	res, err := db.Query(fullStatement)
	if err != nil {
		log.Fatal(err)
	}
	
	var words []string
	errscan := scan.Rows(&words, res)
	if errscan != nil {
		log.Fatal(err)
	}
	
	log.Printf("Got %d words like for example %s",len(words),words[45])
	return words
}

func RandNUniqueOfSlice(count int, arr []string) []string {
	rand.Seed(time.Now().UnixNano())
	out := make([]string,count)
	for i := 0; i < count; i++ {
		var candidate string
		for k := 0; k < count * count * 100 && (slices.Contains(out,candidate)); k++ {
			candidate = arr[rand.Intn(len(arr))]
		}
		out[i] = candidate
	}
	return out
}


