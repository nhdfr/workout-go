package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct{}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWorkoutID := chi.URLParam(r, "id")
	log.Printf("HandleWorkoutByID called, raw param=%q, URL=%s", paramsWorkoutID, r.URL.Path)
	if paramsWorkoutID == "" {
		http.NotFound(w, r)
		return
	}
	workoutID, err := strconv.ParseInt(paramsWorkoutID, 10, 64)
	if err != nil {
		// fmt.Println(workoutID)
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "this is the workout id %d\n", workoutID)
}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "created a workout\n")
}
