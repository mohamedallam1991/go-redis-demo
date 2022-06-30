package checking

import "log"

func Checking(err error, m ...string) {
	if err != nil {
		log.Fatal(err, m)
		// log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
}
