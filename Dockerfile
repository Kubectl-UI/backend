FROM golang:alpine AS builder

WORKDIR /phirmware/go/src/kubectl-gui

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o main .

FROM scratch

COPY --from=builder /phirmware/go/src/kubectl-gui/main /

ENTRYPOINT ["/main"]


