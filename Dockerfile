FROM golang:1.25.4-trixie AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./

RUN go mod download && go mod verify

ARG MAXPROCS=1
COPY . .

# Creates a staticaly-linked binary
RUN CGO_ENABLED=0 go build -v -o ./bin/ -gcflags=GOMAXPROCS=$MAXPROCS ./...

FROM scratch

COPY --from=builder /usr/src/app/bin/server ./

CMD [ "/server" ]
