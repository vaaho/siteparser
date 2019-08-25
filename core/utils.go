package core

import "log"

const CsvSeparator = ";"

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
