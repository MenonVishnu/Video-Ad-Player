package database

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/MenonVishnu/Video-Ad-Player/backend/helpers"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// connection to Database
func init() {
	var err error

	DB, err = sql.Open("sqlite3", filepath.Join("data", "video_ad_player.db"))
	if err != nil {
		log.Fatal("Error Connecting to Database: ", err)
	}

	//creating log Table
	sqlStatement := `CREATE TABLE IF NOT EXISTS clickdata (clickid INTEGER PRIMARY KEY AUTOINCREMENT, adid INTEGER, timestamp TEXT, ip VARCHAR(20), videotimestamp DOUBLE);`

	_, err = DB.Exec(sqlStatement)

	if err != nil {
		log.Fatal("Error Creating clickdata Table: ", err)
	}
	log.Println("clickdata Table Available")

	//creating advertisement table
	sqlStatement = `CREATE TABLE IF NOT EXISTS advertisement (adid INTEGER PRIMARY KEY, imageurl TEXT, targeturl TEXT);`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		log.Fatal("Error Creating advertisement Table: ", err)
	}
	log.Println("advertisement Table Available!!")

	//to check if there is any data in advertisement table
	var tempData helpers.AdvData
	sqlStatement = `SELECT adid FROM advertisement;`
	err = DB.QueryRow(sqlStatement).Scan(&tempData)
	if err == sql.ErrNoRows {
		InsertDummyData("dummydata.json")
	} else {
		log.Println("Dummy Data Already Present")
	}

}

// inserting dummy data into the table using json file
func InsertDummyData(filename string) {

	cwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(cwd, "data", filename))
	if err != nil {
		log.Fatal("Could not find file: ", err)
	}
	defer file.Close()

	// Read file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error Reading JSON file: ", err)
	}

	// Parse JSON
	var dummyData []helpers.AdvData
	err = json.Unmarshal(bytes, &dummyData)
	if err != nil {
		log.Fatal("Error Parsing JSON data: ", err)
	}

	// Check if there's data to insert
	if len(dummyData) == 0 {
		log.Println("No data found in JSON file.")
		return
	}

	query := `INSERT INTO advertisement (adid, imageurl, targeturl) VALUES `
	values := []interface{}{}

	for _, data := range dummyData {
		query += "(?,?,?),"
		values = append(values, data.AdID, data.ImageUrl, data.TargetUrl)
	}
	query = query[:len(query)-1] + ";"

	_, err = DB.Exec(query, values...)
	if err != nil {
		log.Fatal("Error Inserting dummy data in Table: ", err)
	}

	log.Println("Dummy Data successfully added!!")
}

// get all the data from advertisement table
func GetAllAdv() ([]helpers.AdvData, error) {
	query := `SELECT adid, imageurl, targeturl FROM advertisement;`
	result, err := DB.Query(query)
	if err != nil {
		log.Fatal("Error executing query: ", err)
	}
	defer result.Close()

	var adv []helpers.AdvData
	for result.Next() {
		var data helpers.AdvData
		err = result.Scan(&data.AdID, &data.ImageUrl, &data.TargetUrl)
		if err != nil {
			return nil, err
		}
		adv = append(adv, data)
	}
	log.Println("Fetched Ads from Database.")
	return adv, nil
}

// add click data into the clickdata table
func AddClick(clickData helpers.ClickData) error {
	query := `INSERT INTO clickdata (adid, timestamp, ip, videotimestamp) VALUES (?,?,?,?);`
	_, err := DB.Exec(query, clickData.AdID, clickData.Timestamp, clickData.IP, clickData.VideoTimeStamp)
	if err != nil {
		return err
	}
	log.Println("Click Logged in Database.")
	return nil
}
