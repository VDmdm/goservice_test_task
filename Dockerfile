FROM golang:1.19rc2-alpine3.16 as build_stage
WORKDIR /build
COPY . /build
RUN go build -o app

FROM alpine:latest
COPY --from=build_stage /build/app .
ENTRYPOINT ["./app"]
