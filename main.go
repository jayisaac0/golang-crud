package main

import (
	"github.com/jayisaac0/auth-service/src/cmd"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// main function entry point
func main() {
	cmd.HandleRequets()
}
