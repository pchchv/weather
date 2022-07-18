# HTTP API service returning time and weather information
# Running the application
```
docker-compose up --build
```
# Running the application without Docker
```
go run .
```
## Running tests (app must be running)
```
go test
```
## HTTP Methods
```
/stats — Getting the time and weather
    example: http://localhost:8080/stats?city=Los%20Angeles
```
```
/time — Getting only the time
    example: http://localhost:8080/time?city=Siem Reap
```
```
/weather — Getting only the weather
    example: http://localhost:8080/time?city=Porto
```
### Params for ```.env``` file
```
APIKEY=000000000000000  // for service openweathermap.org
APIUSER=user            // for service geonames.org
```
