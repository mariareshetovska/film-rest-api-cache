package main

import (
	"encoding/json"
	"film-rest-api/cach"
	"film-rest-api/db"
	"film-rest-api/utils"
	"log"

	"net/http"

	"time"

	_ "github.com/lib/pq"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

var cacheRedis cach.CacheRedis
var cacheApp cach.CacheApp

func main() {
	InitRedisCache()
	InitAppCache()

	r := mux.NewRouter()
	r.HandleFunc("/films/{title}", GetFilmByTitle).Methods("GET")
	r.HandleFunc("/logs", GetResponseLogs).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
	}
	log.Println("Server is running on port 3000")
	log.Fatal(srv.ListenAndServe())
}

func GetFilmByTitle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	params := mux.Vars(r)
	title := params["title"]
	request := "/film/" + title

	// verify and get memory cache
	b, err := cacheApp.GetApp(title)
	if err != nil {
		log.Fatal(err)
	}
	var result *db.Film

	// if memory cache exist
	if b != nil {
		err := json.Unmarshal(b, &result)
		if err != nil {
			log.Fatal(err)
		}
		utils.ToJson(w, result)
		timeMemory := time.Since(start).Microseconds()

		log.Println("timeMemory", timeMemory)
		db.InsertMemoryDBLog(request, timeMemory)
		return
		// if memory cache doesn't exist
	} else {
		//get cache from redis if exist
		result = cacheRedis.GetRedis(title)
		if result == nil {
			// get cache from database
			result, err := db.GetFilmByTitle(title)
			if err != nil {
				utils.ErrorResponse(w, err, http.StatusBadRequest)
				return
			}
			utils.ToJson(w, result)

			// insert response time log to TimeDB
			timeDB := time.Since(start).Microseconds()
			log.Println("timeDB", timeDB)
			db.InsertTimeDBLog(request, timeDB)
			// caching data into redis
			cacheRedis.SetRedis(title, result, 30)

		} else {
			// caching response into App memory
			cacheApp.SetApp(title, result, 15*time.Second)
			utils.ToJson(w, result)

			// insert response time log of redis into table
			timeRedis := time.Since(start).Microseconds()
			log.Println("timeRedis", timeRedis)
			db.InsertRedisDBLog(request, timeRedis)
		}
	}

}

func GetResponseLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := db.GetAllLogs()
	if err != nil {
		log.Fatal(err)
		return
	}
	utils.ToJson(w, logs)
}

func InitRedisCache() {
	cacheRedis = &cach.RedisCache{
		Client: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
		}),
	}

}

func InitAppCache() {
	cacheApp = &cach.AppCache{
		Client: cache.New(5*time.Minute, 10*time.Minute),
	}
}
