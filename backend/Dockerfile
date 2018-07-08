FROM golang:1.10.3-alpine3.7

ARG arg_firebase_json_name

RUN apk update && \
    apk add git && \
    go get -u github.com/gorilla/mux && \
    go get -u github.com/gorilla/handlers && \
    go get -u firebase.google.com/go && \
    go get -u google.golang.org/api/option

ENV CREDENTIALS=/firebase/credentials/${arg_firebase_json_name}
RUN mkdir -p $CREDENTIALS
ADD ./app/credentials/${arg_firebase_json_name} /firebase/credentials

EXPOSE 8000

WORKDIR /app

