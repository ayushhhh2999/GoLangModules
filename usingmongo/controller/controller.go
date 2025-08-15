package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ayushhhh2999/mymodules/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Replace <password> with your actual passwords
	connectionString = "mongodb+srv://ayushsingh2999:<password>@cluster0.epuwalb.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	dbName           = "netflix"
	colName          = "watchlist"
)

var collection *mongo.Collection

// Connect to MongoDB (init runs when this package is imported)
func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("MongoDB ping error: %v", err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)
	if collection == nil {
		log.Fatal("Mongo collection creation failed")
	} else {
		fmt.Println("Collection set successfully")
	}
}

// insertone inserts a new movie document.
func insertone(movie models.Netflix) error {
	_, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}
	fmt.Println("Movie inserted successfully")
	return nil
}

// getallmovies fetches all movies from collection.
func getallmovies() ([]models.Netflix, error) {
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("find failed: %w", err)
	}
	defer cur.Close(context.TODO())

	var movies []models.Netflix
	for cur.Next(context.TODO()) {
		var movie models.Netflix
		if err := cur.Decode(&movie); err != nil {
			return nil, fmt.Errorf("decode failed: %w", err)
		}
		movies = append(movies, movie)
	}
	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}
	return movies, nil
}

// getonemovie fetches a single movie by ID.
func getonemovie(id string) (*models.Netflix, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": objID}
	var movie models.Netflix
	err = collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

// deletemovie deletes a single movie by ID.
func deletemovie(id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": objID}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return result.DeletedCount > 0, nil
}

// deleteallmovies deletes all movies.
func deleteallmovies() (int64, error) {
	result, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

// markwatched marks a movie as watched by ID.
func markwatched(id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, fmt.Errorf("invalid id: %w", err)
	}
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

// Getallthemovies handler returns all movies as JSON array.
func Getallthemovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	movies, err := getallmovies()
	if err != nil {
		http.Error(w, "Failed to fetch movies", http.StatusInternalServerError)
		return
	}
	if len(movies) == 0 {
		http.Error(w, "No movies found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
	}
}

// Getonethemovies handler returns one movie by ID as JSON.
func Getonethemovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "ID parameter missing", http.StatusBadRequest)
		return
	}

	movie, err := getonemovie(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Movie not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch movie", http.StatusInternalServerError)
		}
		return
	}

	if err := json.NewEncoder(w).Encode(movie); err != nil {
		http.Error(w, "Failed to encode movie", http.StatusInternalServerError)
	}
}

// Deleteallthemovies handler deletes all movies.
func Deleteallthemovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	deletedCount, err := deleteallmovies()
	if err != nil {
		http.Error(w, "Failed to delete movies", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%d movie(s) deleted successfully\n", deletedCount)
}

// Deletetheone handler deletes one movie by ID.
func Deletetheone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	deleted, err := deletemovie(id)
	if err != nil {
		http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
		return
	}
	if !deleted {
		http.Error(w, "No movie found with given ID", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Movie with ID %s deleted successfully\n", id)
}

// Markthewatched handler marks one movie as watched.
func Markthewatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	updated, err := markwatched(id)
	if err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}
	if !updated {
		http.Error(w, "No movie found with given ID", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Movie with ID %s marked as watched successfully\n", id)
}

// Addthemovie handler adds a new movie.
func Addthemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie models.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if movie.Movie == "" {
		http.Error(w, "Movie title is required", http.StatusBadRequest)
		return
	}
	if err := insertone(movie); err != nil {
		http.Error(w, "Failed to add movie", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Movie '%s' added successfully\n", movie.Movie)
}
