package hello

import (
	"log"
	"net/http"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/log"

	// @see https://github.com/ziutek/mymysql
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	// _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
	"fmt"
	"github.com/ziutek/mymysql/mysql"
)

func init() {
	http.HandleFunc("/", handler)

	// Leave the 5th parameter (database name) blank because we drop and recreate the database in install.sql
	db := mysql.New("tcp", "", "127.0.0.1:3306", "user", "password","acgtweet")
	// db, err := sql.Open("mysql", "user@cloudsql(project-id:instance-name)/dbname")

	err := db.Connect()


	//buff, err := ioutil.ReadFile("install.sql")
	//if err != nil {
	//	panic(err)
	//}
	//s := string(buff)

	//result, err := db.Start(s);
	//log.Printf("err: %v", err);
	//log.Printf("Result: %v", result);
	//// Now we've installed, set the database name or following commands wont work.
	//db.Use("acgtweet");


	// TODO Change to prepared statements and get a datatype back: https://github.com/ziutek/mymysql#example-2---prepared-statements
	res, err :=  db.Start("SELECT * FROM users;")
	log.Printf("Getting users..")
	log.Printf("Error: %v", err)


	fields := res.Fields()

	fields

	// Print all rows
	for {
		row, err := res.GetRow()
		checkError(err)

		if row == nil {
			// No more rows
			break
		}

		// Print all cols
		for _, col := range row {
			if col == nil {
				fmt.Print("<NULL>")
			} else {
				log.Printf("Column: %v=%v", string(col.([]byte)))
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func checkError(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")

	// To log.
	//c := appengine.NewContext(r)
	//c.Infof("Requested URL: %v", r.URL)
}
