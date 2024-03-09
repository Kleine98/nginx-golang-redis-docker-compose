package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Test Redis connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong) // Output: PONG

	// Define handler function for setting value to a list
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		// Parse JSON request body
		var requestBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Extract key and value from request body
		key := requestBody["key"].(string)
		value := requestBody["value"]

		// Push the value to the list in Redis
		_, err := client.LPush(ctx, key, value).Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Value successfully added to the list in Redis"))
	})

	// Define handler function for getting all values by key
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		// Parse query parameters
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			http.Error(w, "Missing key parameter", http.StatusBadRequest)
			return
		}
		key := keys[0]

		// Get all values from the list in Redis
		vals, err := client.LRange(ctx, key, 0, -1).Result()
		if err != nil {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}

		// Convert values to JSON and send response
		jsonResponse, err := json.Marshal(vals)
		if err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Start HTTP server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
