package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	r := mux.NewRouter()

	app, err := makeNewRelicApplication()
	if err != nil {
		log.Fatalf("error init newrelic apm agent reason: %v", err)
		os.Exit(1)
	}

	r.Use(nrgorilla.Middleware(app))

	r.HandleFunc("/example/{name}", ExampleHandler)

	http.ListenAndServe(":8000", r)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	log.Printf("name is: %v", vars["name"])
}

func makeNewRelicApplication() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("newrelic-zap-logs-in-context"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogEnabled(false),
	)
	if err != nil {
		return nil, err
	}

	if err = app.WaitForConnection(5 * time.Second); err != nil {
		return nil, err
	}

	return app, nil
}
