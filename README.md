<img src="https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png" alt="yoda-gopher" width=350>

[![CICD](https://github.com/borischen0203/litclock-service/actions/workflows/CICD.yml/badge.svg)](https://github.com/borischen0203/litclock-service/actions/workflows/CICD.yml)


# litclock-service
This is a litclock-service. You can convert numeric time to human friendly text.

Numeric Time -> Human Friendly Text:
- 1:00 -> One o'clock
- 13:05 -> Five past one

## Features
- Convert numeric time to human text.

# How to use

## Run directly:
- The service already deployed on Heroku.
- No need to run main.go file, you can type the below command in the terminal directly.
- `The app may sleep without using. Just wait for a few seconds to wake it up.`

Required
- Install `curl` cli(https://formulae.brew.sh/formula/curl)

Use below curl command to call API and input any time in numericTime value
```bash
curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "https://litclock-service.herokuapp.com/api/litclock-service/v1/numeric-time"
```

Demo
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "https://litclock-service.herokuapp.com/api/litclock-service/v1/numeric-time"
$ {"textTime":"Eight past one"}
```

## Run in Postman:
Use below API path and body in Postman
```bash
[POST] https://litclock-service.herokuapp.com/api/litclock-service/v1/numeric-time
```
Body
```bash
{
    "numericTime" : "13:08"
}
```

## Run in Docker:
Required
- Install docker

### Run process
Step1: Pull docker image (borischen0203/litclock-service)
```bash
docker pull borischen0203/litclock-service
```
Step2: Run docker image as below command
```bash
docker run -it --rm -p 8080:8080 borischen0203/litclock-service
```
Step3: Create a new terminal window and execute curl
```bash
curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "http://localhost:8080/api/litclock-service/v1/numeric-time"
```

## Run in Local:

Required
- Install go(version >= 1.6)
- Install `make` cli(https://formulae.brew.sh/formula/make)
- Install `curl` cli(https://formulae.brew.sh/formula/curl)

```bash
# clone a repo
git clone https://github.com/borischen0203/litclock-service.git

# Use `make` to execute makefile run test and build
make all

# Run the local server
make run

# Use `curl` to call API
curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "http://localhost:8080/api/litclock-service/v1/numeric-time"
```

### Local run demo:
```bash
$ curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "http://localhost:8080/api/litclock-service/v1/numeric-time"
$ {"textTime":"Eight past one"}
```

## Tech Stack
- Golang
- Gin framework
- RESTful API
- Swagger
- Docker
- Github action(CI)
- Heroku (CD)

## Todo:
- [ ]
