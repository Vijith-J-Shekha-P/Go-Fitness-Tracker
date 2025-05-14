package main

import (
	"fmt"
	"time"
)

// Workout interface
type Workout interface {
	Duration() time.Duration
	CaloriesBurned() float64
	RecordStats()
	GetType() string
}

// CardioWorkout struct
type CardioWorkout struct {
	duration     time.Duration
	distance     float64
	avgHeartRate float64
}

func (c CardioWorkout) Duration() time.Duration {
	return c.duration
}

func (c CardioWorkout) CaloriesBurned() float64 {
	minutes := c.duration.Minutes()
	return minutes * 10 * (c.avgHeartRate / 100)
}

func (c CardioWorkout) RecordStats() {
	fmt.Printf("Cardio Workout - Duration: %.0f mins, Distance: %.2f km, Avg Heart Rate: %.0f bpm, Calories: %.2f\n",
		c.duration.Minutes(), c.distance, c.avgHeartRate, c.CaloriesBurned())
}

func (c CardioWorkout) GetType() string {
	return "Cardio"
}

// StrengthWorkout struct
type StrengthWorkout struct {
	duration time.Duration
	weight   int
	reps     int
	sets     int
}

func (s StrengthWorkout) Duration() time.Duration {
	return s.duration
}

func (s StrengthWorkout) CaloriesBurned() float64 {
	minutes := s.duration.Minutes()
	return minutes * 5 * (float64(s.weight) / 10)
}

func (s StrengthWorkout) RecordStats() {
	fmt.Printf("Strength Workout - Duration: %.0f mins, Weight: %d kg, Reps: %d, Sets: %d, Calories: %.2f\n",
		s.duration.Minutes(), s.weight, s.reps, s.sets, s.CaloriesBurned())
}

func (s StrengthWorkout) GetType() string {
	return "Strength"
}

func summarizeWorkouts(workouts []Workout) {
	for _, workout := range workouts {
		fmt.Println("Workout Type:", workout.GetType())
		fmt.Println("Duration:", workout.Duration().Minutes(), "minutes")
		workout.RecordStats()
		fmt.Println()
	}
}

func main() {
	var workouts []Workout

	cardio := CardioWorkout{
		duration:     30 * time.Minute,
		distance:     5.0,
		avgHeartRate: 140,
	}
	workouts = append(workouts, cardio)

	strength := StrengthWorkout{
		duration: 45 * time.Minute,
		weight:   50,
		reps:     10,
		sets:     3,
	}
	workouts = append(workouts, strength)

	fmt.Println("Workout Summary:")
	summarizeWorkouts(workouts)
}
