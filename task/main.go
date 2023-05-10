package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AkifhanIlgaz/gophercises/task/cmd"
	"github.com/AkifhanIlgaz/gophercises/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	cmd.Execute()
}

// Helper function to handle errors
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
