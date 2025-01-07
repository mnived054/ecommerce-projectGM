package main

import (
	"ecommerce-project/backend/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//Routes for handling user authentication
	r.HandleFunc("/api/signup", handlers.SignUpHandler).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")

	//Profile routes
	r.HandleFunc("/api/profile", handlers.GetProfileHandler).Methods("GET")
	r.HandleFunc("/api/profile/update", handlers.UpdateProfileHandler).Methods("POST")

	//Order routes
	r.HandleFunc("/api/orders.create", handlers.CreateOrderHandler).Methods("POST")
	r.HandleFunc("/api/orders", handlers.GetOrdersHandler).Methods("GET")

	//Cart routes
	r.HandleFunc("/api/cart", handlers.GetCartHandler).Methods("GET")
	r.HandleFunc("/api/cart/add", handlers.AddToCartHandler).Methods("POST")

	//Address routes
	r.HandleFunc("/api/address", handlers.GetAddressesHandler).Methods("GET")
	r.HandleFunc("/api/address/add", handlers.AddAddressHandler).Methods("POST")
	r.HandleFunc("/api/address/update", handlers.UpdateAddressHandler).Methods("POST")
	r.HandleFunc("/api/address/delete", handlers.DeleteAddressHandler).Methods("DELETE")

	//Start the server
	fmt.Println("Server is running at https://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}
