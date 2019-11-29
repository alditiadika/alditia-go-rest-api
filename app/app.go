package app

import (
	"context"
	"fmt"
	"github.com/alditiadika/alditia-go-rest-api/app/handle"
	"github.com/alditiadika/alditia-go-rest-api/config"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

// App model
type App struct {
	Router *mux.Router
	DB     *mongo.Client
}

//Initialize method
func (a *App) Initialize() {
	a.connectDB()
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	//user route
	a.Get("/user", a.handleRequest(handle.GetUser))
	a.Post("/user", a.handleRequest(handle.Insertuser))
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

//Get Method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

//Post method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}
func (a *App) handleRequest(handler func(w http.ResponseWriter, r *http.Request, Clnt *mongo.Client)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, a.DB)
	}
}

func (a *App) connectDB() {
	conf := config.GetConf()
	clientOption := options.Client().ApplyURI(conf.URL)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		fmt.Println("error")
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("error, can't PING to database")
		log.Fatal(err)
	}
	a.DB = client
	fmt.Println("Connected to MongoDB!")
}
