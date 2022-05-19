FROM alpine AS build
RUN apk update
WORKDIR /go
COPY ./main.go ./main.go
RUN apk add --update go gcc g++ && go build main.go

FROM alpine
WORKDIR /go
COPY --from=build /go/main /go
EXPOSE 80
CMD ["./main"]

