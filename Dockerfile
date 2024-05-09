FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o survey-app ./cmd/main.go

CMD ["./survey-app"]