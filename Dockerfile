FROM golang:1.14.1-alpine

WORKDIR /build
COPY . .
RUN go build -o bin/nino

#FROM alpine:latest
#COPY --from=builder /build/bin/nino /app
#WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["bin/nino"]