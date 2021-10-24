# FROM golang:1.16-alpine
# WORKDIR /app
# ADD . /app

# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# RUN cd /app && go build -o main

# CMD [ "/app/main" ]

# build stage
FROM golang:1.16-alpine AS build
ADD . /src
RUN cd /src && go build -o litclock-service

# final stage
FROM alpine as runtime

# # add TimeZone tzdata
# RUN apk add --no-cache tzdata
# ENV TZ=Europe/London

WORKDIR /app
COPY --from=build /src/litclock-service /app/litclock-service
COPY .env /app/
# RUN cd /app && chmod +x run.sh

CMD [ "/app/litclock-service" ]







