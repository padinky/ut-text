# ut-text

A homework for backend assessment in Ulven Tech

it is using public library below :
- gin-gonic, Go http web framework. See : https://github.com/gin-gonic/gin
- gubrak, Go functional utility library with syntactic sugar. It's like lodash, but for Go Programming language. See : https://github.com/novalagung/gubrak

## How to run :
Simply clone this repo, then run standard golang command below :
```sh
go mod tidy
go run main.go
```
it will listen to port ```:3000```

## Give it a shot :
```sh
curl -X POST http://localhost:3000/submit -H 'Content-Type: application/json' -d '{"text":"lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet lorem ipsum dolor sit amet hahahaha hehehe huhuhu hahaha ipsum ipsum ipsum ipsum"}'

It will return top 10 of most used word in array of json with structure :
```[{"text":"","count":0}]```

Thank you for the opportunity and consideration.