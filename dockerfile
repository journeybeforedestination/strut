FROM golang:1.20 as build

WORKDIR /strut

COPY . .
RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM gcr.io/distroless/base-debian11 as release

COPY --from=build /app /app

EXPOSE 8080

ENTRYPOINT [ "/app" ]