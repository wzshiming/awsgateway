package awsgateway_test

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wzshiming/awsgateway"
)

func Example() {
	http.HandleFunc("/", hello)
	log.Fatal(awsgateway.ListenAndServe(":3000", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from Go")
}
