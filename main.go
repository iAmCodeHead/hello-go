package main

import (
	"fmt"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
)

type server struct{}

func getRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world fetched successfully"}`))
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world created successfully"}`))
}

func putRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world updated successfully"}`))
}

func patchRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world patched successfully"}`))
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world deleted successfully"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
	if val, ok := pathParams["commentID"]; ok {
		commentID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	query := r.URL.Query()
	location := query.Get("location")

	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world fetched successfully"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "hello world created successfully"}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world updated successfully"}`))
	case "PATCH":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world patched successfully"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello world deleted successfully"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

// func main() {
// 	greetings := "Hello, world"
// 	fmt.Println("capital greetings to you: ", greetings)
// }

func main() {
	// http.HandleFunc("/home", home)
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/home", home)
	api.HandleFunc("/get", getRequest).Methods(http.MethodGet)
	api.HandleFunc("/post", postRequest).Methods(http.MethodPost)
	api.HandleFunc("/put", putRequest).Methods(http.MethodPut)
	api.HandleFunc("/delete", deleteRequest).Methods(http.MethodDelete)
	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServe(":8080", r))
}
