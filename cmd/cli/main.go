package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/joshuaAllday/perkbox/pkg/placeholder"
	"github.com/joshuaAllday/perkbox/pkg/utils"
	"github.com/pkg/errors"
)

var (
	path string = "./input/data.csv"
	url  string = "https://jsonplaceholder.typicode.com/todos"
)

type Cli struct {
	ctx     context.Context
	service *placeholder.Placeholder
}

var cli *Cli

func init() {
	cli = &Cli{
		ctx: context.Background(),
	}
}

func main() {

	pls, err := placeholder.NewToDos(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	cli.service = pls
	log.Println("Loading ids from file...")

	csv, err := utils.ReadCsvFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}

	todos, err := cli.service.GetTodos(cli.ctx, csv)
	if err != nil {
		log.Fatal(err.Error())
	}

	btes, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to json encode the struct"))
	}

	log.Println("Output: \n", string(btes))
	log.Println("Finished")
}
