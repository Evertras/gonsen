package main

import "log"

func main() {
	s := newServer()

	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen and serve: %v", err)
	}
}
