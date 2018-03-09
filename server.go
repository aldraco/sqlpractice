package main

import (
  "database/sql"
  "fmt"
  "strconv"

  _ "github.com/mattn/go-sqlite3"
  "github.com/andlabs/ui"
)

func main() {
  database, _ := sql.Open("sqlite3", "./learning.db")
  statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
  statement.Exec()
  //statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
  //statement.Exec("Ashley", "Drake")
  rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
  var id int
  var firstname string
  var lastname string
  for rows.Next() {
    rows.Scan(&id, &firstname, &lastname)
    fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
  }
  doWindows(database)

}

func runQuery(database *sql.DB, queryString string) {
  rows, _ := database.Query("SELECT id, firstname FROM people where firstname == (?)", queryString)
  //rows, _ := statement.Exec(queryString)
  var id int
  var firstname string
  for rows.Next() {
    rows.Scan(&id, &firstname)
    fmt.Println(strconv.Itoa(id) + ": " + firstname)
  }
}

func doWindows(database *sql.DB) {
  err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 800, 500, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
      runQuery(database, input.Text())
			greeting.SetText("Hello, " + input.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}




// function to download the database
// kaggle has some but you need to sign in (auth?)
// net/http
//io, os
// out, err := os.Create("dbfile.db")
// defer out.Close()
// resp, err := http.Get('URL OF DATABSE')
// defer resp.Body.Close()
// n, err := io.Copy(out, resp.Body)


// function to set up the database
// depends on what foramt it arrives in, how to standardize


// function to clean up
// should delete any existing databasefiles
// os.Remove()
