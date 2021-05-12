# FROM arm32v7/golang:1.12.13 AS builder
FROM golang:latest AS builder
RUN mkdir /go/src/ampgo
WORKDIR /go/src/ampgo

COPY ampgoserver.go .
COPY ampgosetup.go .
COPY artist.go .
COPY album.go .
COPY songs.go .
COPY mp3.go .
COPY go.mod .
COPY go.sum .
RUN export GOPATH=/go/src/ampgo
RUN go get -v /go/src/ampgo
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/ampgo

# FROM arm32v6/alpine:latest
FROM alpine:latest
# RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/ampgo/main .
RUN \
  mkdir ./data && \
  mkdir ./data/db && \
  mkdir ./static && \
  mkdir ./static/css && \
  mkdir ./static/templates && \
  chmod -R +rwx ./static

COPY static/animals.jpg ./static/

COPY static/css/loginstyles.css ./static/css/

COPY static/templates/home.html ./static/templates/
COPY static/templates/intro.html ./static/templates/
COPY static/templates/artist.html ./static/templates/
COPY static/templates/album.html ./static/templates/
COPY static/templates/song.html ./static/templates/
COPY static/templates/playlist.html ./static/templates/

RUN \
  mkdir ./fsData && \
  mkdir ./fsData/thumb && \
  chmod -R +rwx ./fsData && \
  mkdir ./logs && \
  chmod -R +rwx ./logs && \
  echo "Creating log file" > ./logs/ampgo_log.txt && \
  chmod -R +rwx ./logs/ampgo_log.txt
STOPSIGNAL SIGINT
CMD ["./main"]

