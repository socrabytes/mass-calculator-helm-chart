package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"math"
	"log"
)

// Mass struct holds the density of a material.
type Mass struct {
	Density float64
}

// MassVolume declares methods for density and volume calculations
type MassVolume interface {
	density() float64        
	volume(dimension float64) float64 
}

// Sphere represents a sphere made of a material.
type Sphere struct {
	Mass 
}

// density method returns the density of the sphere's material.
func (s Sphere) density() float64 {
	return s.Density
}

// volume method calculates the volume of the sphere based on its diameter.
func (s Sphere) volume(dimension float64) float64 {
	radius := dimension / 2.0 
	return (4.0 / 3.0) * math.Pi * math.Pow(radius, 3) 
}

// Cube represents a cube made of a material.
type Cube struct {
	Mass 
}

// density method returns the density of the cube's material.
func (c Cube) density() float64 {
	return c.Density
}

// volume method calculates the volume of the cube based on its side length.
func (c Cube) volume(dimension float64) float64 {
	return math.Pow(dimension, 3) 
}

// Handler function processes HTTP requests and calculates the weight based on dimension and material density.
func Handler(massVolume MassVolume) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s", r.URL.String())
		if dimension, err := strconv.ParseFloat(r.URL.Query().Get("dimension"), 64); err == nil {
			weight := massVolume.density() * massVolume.volume(dimension)
			w.Write([]byte(fmt.Sprintf("%.2f", math.Round(weight * 100) / 100)))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}
}

// healthz function for liveness probe
func healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Healthz endpoint hit")
	w.WriteHeader(http.StatusOK)
}

// readyz function for readiness probe
func readyz(w http.ResponseWriter, r *http.Request) {
	log.Println("Readyz endpoint hit")
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Get the port number from command-line arguments.
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Error parsing port: %v", err) 
	}

	// Initialize Sphere and Cube with their respective densities.
	aluminiumSphere := Sphere{Mass{Density: 2.710}} // Density of aluminum is 2.710 g/cm^3
	ironCube := Cube{Mass{Density: 7.874}}         // Density of iron is 7.874 g/cm^3

	
	// Set up HTTP handlers for calculating weights of the sphere and cube.
	http.HandleFunc("/aluminium/sphere", Handler(aluminiumSphere))
	http.HandleFunc("/iron/cube", Handler(ironCube))
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/readyz", readyz)

	log.Printf("Starting server on port %d", port)
	// Start the HTTP server on the specified port.
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
