package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	gowordstat "github.com/Sagleft/go-wordstat"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

	log.Println("done")
}

func run() error {
	client, err := gowordstat.NewClient(os.Getenv("YANDEX_API_TOKEN"))
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	data, err := client.GetRegions()
	if err != nil {
		return fmt.Errorf("get regions: %w", err)
	}

	PrintObject(data)
	return nil
}

func PrintObject(o any) {
	data, err := json.MarshalIndent(o, "", "	")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(data))
}
