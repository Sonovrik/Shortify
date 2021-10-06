FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN go mod download

#RUN go get github.com/githubnemo/CompileDaemon
#ENTRYPOINT CompileDaemon -command="./shortify"


EXPOSE ${SHORTIFY_CONTAINER_PORT}

CMD ["sh", "start.sh"]