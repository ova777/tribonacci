## How to run service locally
### Run without build
```sh
go get github.com/go-martini/martini
go run server.go
```
### Build and run binary
```sh
go get github.com/go-martini/martini
go build -o main .
./main
```
Open [http://localhost:3000/tribonacci/10](http://localhost:3000/tribonacci/10) in your browser

## How to run service in docker container
```sh
docker build -t ova777:api .
docker run -p PORT:3000 ova777:api
```
where `PORT` is the free port on your docker ip address

Then open [http://IP:PORT/tribonacci/10](http://IP:PORT/tribonacci/10)
in your browser, where `IP` is the your docker ip address
