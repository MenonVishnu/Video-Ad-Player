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

//global sql DB variable 
var DB *sql.DB

// On Startup
func init() {
	var err error

	//Database connection
	DB, err = sql.Open("sqlite3", filepath.Join("data", "video_ad_player.db"))
	if err != nil {
		log.Fatal("Error Connecting to Database: ", err)
	}

	//creating clickdata Table
	sqlStatement := `CREATE TABLE IF NOT EXISTS clickdata (clickid INTEGER PRIMARY KEY AUTOINCREMENT, adid INTEGER, timestamp TEXT, ip VARCHAR(20), videotimestamp DOUBLE);`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		log.Fatal("Error Creating clickdata Table: ", err)
	}
	log.Println("clickdata Table Available.")

	//creating advertisement table
	sqlStatement = `CREATE TABLE IF NOT EXISTS advertisement (adid INTEGER PRIMARY KEY, imageurl TEXT, targeturl TEXT);`
	_, err = DB.Exec(sqlStatement)
	if err != nil {
		log.Fatal("Error Creating advertisement Table: ", err)
	}
	log.Println("advertisement Table Available.")

	//to check if there is any data in advertisement table
	var tempAdv helpers.AdvData
	sqlStatement = `SELECT adid FROM advertisement;`
	err = DB.QueryRow(sqlStatement).Scan(&tempAdv)
	//if no rows present in result
	if err == sql.ErrNoRows { 
		InsertDummyData("dummydata.json")
	} else {
		log.Println("Dummy Data Already Present")
	}

}

// Inserting dummy data into the table using json file
func InsertDummyData(filename string) {
	//Get current working directory
	cwd, _ := os.Getwd()
	file, err := os.Open(filepath.Join(cwd, "data", filename))
	if err != nil {
		log.Println("Could not find file: ", err)
		return
	}
	defer file.Close()

	// Reading file content
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("Error Reading JSON file: ", err)
		return
	}

	// Parse JSON
	var dummyData []helpers.AdvData
	err = json.Unmarshal(bytes, &dummyData)
	if err != nil {
		log.Println("Error Parsing JSON data: ", err)
		return
	}

	// Handles no data in dummydata
	if len(dummyData) == 0 {
		log.Println("No data found in JSON file.")
		return
	}

	//Inserts dummy data in advertisement
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
