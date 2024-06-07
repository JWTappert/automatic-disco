package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type Event struct {
	StartDate time.Time
	Duration  int
	Title     string
	Class     string
	Style     string
}

func randate() time.Time {
	now := time.Now()
	min := time.Date(now.Year(), now.Month(), now.Day()-3, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(now.Year(), now.Month(), now.Day()+3, 23, 59, 59, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

func main() {
	chi := chi.NewRouter()

	chi.Get("/", func(w http.ResponseWriter, r *http.Request) {
		dummyEvents := make([]Event, 0)

		for i := 0; i < 20; i++ {
			start := randate()
			duration := 0
			if rand.Int63n(10)%2 == 0 {
				duration = 30
			} else {
				duration = 60
			}

			calCol := start.Day() % 7
			hour := (start.Hour() * 2 * 6) + 2
			leng := duration * 12

			newEvent := Event{
				StartDate: start,
				Duration:  duration,
				Title:     "Patient Sim " + strconv.Itoa(i),
				Class:     "relative mt-px flex sm:col-start-" + strconv.Itoa(hour),
				Style:     "grid-row: " + strconv.Itoa(calCol) + " / span " + strconv.Itoa(leng),
			}
			dummyEvents = append(dummyEvents, newEvent)
		}

		component := index(dummyEvents)

		component.Render(r.Context(), w)
	})

	http.ListenAndServe(":5000", chi)
}