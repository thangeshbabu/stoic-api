FROM golang:1.20-bookworm as BASE
EXPOSE 8080

FROM BASE AS BUILD
WORKDIR /build

COPY ./main.go ./go.mod ./
RUN go build -o stoic-api

FROM BASE AS FINAL
WORKDIR /app

COPY --from=BUILD /build/stoic-api ./
COPY template /app/template
ENTRYPOINT ["./stoic-api"]




