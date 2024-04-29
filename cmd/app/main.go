package main

import "to-do-checklist/internal/app"

func main() {
	application := app.NewApp("8080")
	application.Start()
}
