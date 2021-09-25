FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /goapp

ENV PORT=8080

EXPOSE ${PORT}

#CMD [ "/goapp", "-port", "$PORT"]
CMD ["sh", "-c", "/goapp -port $PORT"]