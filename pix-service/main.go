package main

import (
	"github.com/izakdvlpr/codepix/cmd"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cmd.Execute()
}
