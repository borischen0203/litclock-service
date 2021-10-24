# litclock-service
This is a litclock-service. You can convert numeric time to human friendly text.

Numeric Time -> Human Friendly Text:
- 1:00 -> One o'clock
- 13:05 -> Five past one

## Features
- Convert numeric time to human text.


# How to use


## TBD
- No need to run main.go file, you can type the below command in the terminal directly.
- The app may sleep without using. Just wait for a few seconds to wake it up.

### pre:
- Install `curl` cli(https://formulae.brew.sh/formula/curl)

```bash
#Use curl command to call API and input any time in numericTime value
curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "https://litclock-service.herokuapp.com/api/litclock-service/v1/numeric-time"

```
### Demo
```bash
> curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "https://litclock-service.herokuapp.com/api/litclock-service/v1/numeric-time"
> {"textTime":"Eight past one"}
```

## Docker:
```bash
# Step1: docker pull
docker pull borischen0203/litclock-service

# Step2: docker run
docker run -it --rm -p 8080:8080 borischen0203/litclock

# Step3: Create a new terminal window and execute curl
curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "http://localhost:8080/api/litclock-service/v1/numeric-time"
```
### Docker run demo
```bash
>
>
```

## Local:

## Required
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
> curl -X POST -H "Content-Type: application/json" -d '{"numericTime" : "13:08"}' "http://localhost:8080/api/litclock-service/v1/numeric-time"
> {"textTime":"Eight past one"}
```

## Tech Stack
    - Golang
    - Gin framework
    - RESTful API
    - Swagger
    - Docker
    - Github action(CI)
    - Heroku (CD)

