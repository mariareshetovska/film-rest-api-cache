# Golang API caching in Redis and in App Memory 

## General Info

This project demonstrates the caching responses  in Radis and  in the memory of the Web application.
### Work flow
1. The client requests data.
2. The server checks whether there is data in the program's memory.
3. There is memory - it is given. If not, check the **Redis** store.
4. If it is in **Redis** - it is provided. If not, makes a query to the database.
5. Returns data to the client. Updates **Redis** and cache.
6. Inserts the response time of receiving data from the memory cache of the Web app, Redis or the database of the endpoint into the table in the database response_time_log, in which the last response time for a certain request in microseconds will be logged
### Technologies
* Golang
* gorilla/mux
* lib/pq
* joho/godotenv
* Redis
* ElephantSQL
* patrickmn/go-cache

## Setup
Clone repository 

```bash
git clone https://github.com/mariareshetovska/film-rest-api-cache.git
```
Make sure you already have Redis running on port 6379. Build and run the project.
```bash
cd film-rest-api-cache
go build
./film-rest-api
```
## API Reference
These are the endpoints available from the app
### GET films/{title}
Returns a film by title. Example: 
**http://localhost:3000/films/grossewonderful**

```json
{
    "id": 3,
    "title": "Grosse Wonderful",
    "description": "A Epic Drama of a Cat And a Explorer who must Redeem a Moose in Australia",
    "release_year": "2006-01-01T00:00:00Z"
}
```
Returns the response time of receiving data from the memory cache of the Web app, Redis or the database

## GET /logs
```json
[
    {
        "id": 10,
        "request": "/film/academydinosaur",
        "time_db": 1251703,
        "time_redis": 177,
        "time_memory": 170
    },
    {
        "id": 13,
        "request": "/film/grossewonderful",
        "time_db": 1172539,
        "time_redis": 236,
        "time_memory": 113
    }
]