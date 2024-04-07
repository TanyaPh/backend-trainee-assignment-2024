FROM golang:alpine AS builder
WORKDIR /

COPY app/go.mod app/go.sum ./
RUN go mod download

COPY /app ./
RUN go build -o /bin/app ./cmd/app/main.go


FROM alpine AS runner
COPY --from=builder /bin/app /
CMD [ "/app" ]
