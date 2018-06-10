FROM golang:alpine

LABEL maintainer="Min <route666@live.cn>"

# alpine do not include git so add it
RUN apk add --no-cache git

# Set apps working directory
WORKDIR /go/src/github.com/ilovelili/auth0example

ADD . /go/src/github.com/ilovelili/auth0example

# get go dependency
RUN echo "Installing Go dependencies ..."
RUN go get -d github.com/gorilla/mux
RUN go get -d github.com/gorilla/sessions
RUN go get -d github.com/joho/godotenv
RUN go get -d github.com/codegangsta/negroni
RUN go get -d golang.org/x/oauth2

CMD ["go", "run", "main.go", "server.go"]

EXPOSE 3000