FROM golang:1.17.5 AS build

WORKDIR /src/

COPY ./cmd /src/cmd
COPY ./web /src/web
COPY ./main.go /src/main.go
COPY ./internal /src/internal
COPY ./docs /src/docs
COPY ./scripts /src/scripts
COPY ./go.mod /src/go.mod
COPY ./go.sum /src/go.sum

# for build local
#RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 go build -o /src/mdm

FROM scratch

WORKDIR /app

COPY --from=build /src/mdm /bin/mdm
COPY config.yml /app/config.yml

EXPOSE 80
EXPOSE 3000

ENTRYPOINT ["mdm"]
CMD ["serve", "-c", "/app/config.yml"]
