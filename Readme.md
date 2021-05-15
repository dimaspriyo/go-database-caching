# Go Caching Database

## Prerequisite

 1. Port 27017, 8080, 8081
 2. docker-compose 
 
 ## Getting Started
1. `sudo docker-compose up`
2. Hit `http://localhost:8080`

## Description

First Time Access
```
{
"Start": "2021-05-15T16:29:15.788924713+07:00",
"End": "2021-05-15T16:29:16.212689881+07:00",
"TotalMicrosecond": 423765,
"Cache": false,
"Data": [....]
}
```

Second Time Access
```
{
"Start": "2021-05-15T16:29:46.634980139+07:00",
"End": "2021-05-15T16:29:46.64812242+07:00",
"TotalMicrosecond": 13142,
"Cache": true,
"Data": [...]
}
```

Notice the difference in `TotalMicrosecond` 


