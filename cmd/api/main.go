package main

import "log"

func main() {

	config := config{
		addr: ":8191",
	}

	app := &application{
		config: config,
	}

	log.Fatal(app.run())

}
