package main

import (
	"context"
	"github.com/yuzujoe/newrelic-zap-logs-in-context/logger"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	r := mux.NewRouter()

	logger.NewLogger()

	app, err := makeNewRelicApplication()
	if err != nil {
		logger.Lgr.Error("error init newrelic apm agent", nil)
		os.Exit(1)
	}

	logger.Lgr.Info("success newrelic init", nil)

	r.Use(nrgorilla.Middleware(app))

	r.HandleFunc("/example/{name}", ExampleHandler)

	http.ListenAndServe(":8000", r)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(context.Background())
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)

	logger.Lgr.Info("name is: "+vars["name"], txn)
	logger.Lgr.Warn("name is: "+vars["name"], txn)
	logger.Lgr.Error("name is: "+vars["name"], txn)
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
