package db

import (
	"log"
)

type Film struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear string `json:"release_year"`
}

func GetFilmByID(id int64) (*Film, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from films where id=$1"
	rs, err := con.Query(sql, id)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var film Film
	for rs.Next() {
		err := rs.Scan(&film.ID, &film.Title, &film.Description, &film.ReleaseYear)
		if err != nil {
			return nil, err
		}
	}
	return &film, nil
}

func GetFilmByTitle(title string) (*Film, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from films where lower(replace(title, ' ', '')) LIKE $1"
	rs, err := con.Query(sql, title)
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	var film Film
	for rs.Next() {
		err := rs.Scan(&film.ID, &film.Title, &film.Description, &film.ReleaseYear)
		if err != nil {
			return nil, err
		}
	}
	log.Println("get cache from db")
	return &film, nil
}

// insert or update last time Microseconds

func InsertTimeDBLog(request string, timeDB int64) {
	con := Connect()
	defer con.Close()
	sql_insert := "insert into response_time_log (request, time_db) values ($1, $2) ON CONFLICT (request) DO UPDATE SET time_db = $3"
	_, err := con.Query(sql_insert, request, timeDB, timeDB)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertRedisDBLog(request string, timeRedis int64) {
	con := Connect()
	defer con.Close()
	sql_insert := "insert into response_time_log (request, time_redis) values ($1, $2) ON CONFLICT (request) DO UPDATE SET time_redis = $3"
	_, err := con.Query(sql_insert, request, timeRedis, timeRedis)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertMemoryDBLog(request string, timeMemory int64) {
	con := Connect()
	defer con.Close()
	sql_insert := "insert into response_time_log (request, time_memory) values ($1, $2) ON CONFLICT (request) DO UPDATE SET time_memory = $3"
	_, err := con.Query(sql_insert, request, timeMemory, timeMemory)
	if err != nil {
		log.Fatal(err)
	}
}
