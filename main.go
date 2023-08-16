package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Jordation/jsonl/internal"
	"github.com/Jordation/jsonl/provider"
	"github.com/Jordation/jsonl/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.Handle("/events", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		body := map[string]interface{}{}
		res := []*internal.PlayerCombatEvent{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logrus.Fatal(err)
			return
		}

		var pump, _ = provider.NewPump(utils.GetRelativePath("./utils/test.jsonl"))
		deltas, err := pump.GetDeltas(0, int(body["seq"].(float64)))
		if err != nil {
			logrus.Fatal(err)
		}

		pm := utils.GetPlayerMappings()
		player := internal.NewPlayerEventManager(pm, 2)
		ctx := context.Background()
		go player.Start(ctx)

		for _, evt := range deltas {
			player.Ingest(evt)
		}
		res = append(res, player.CombatEvents...)

		logrus.Info(body)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			logrus.Fatal(err)
			return
		}
	})).Methods("POST")
	// Add your routes as needed

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
