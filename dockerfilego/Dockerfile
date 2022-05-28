FROM alpine AS build
WORKDIR /go
RUN apk update && apk add --update go gcc g++
COPY ./main1.go ./main1.go
RUN go build main1.go

FROM alpine
WORKDIR /go
COPY --from=build /go/main1 /go
EXPOSE 80
CMD ["./main1"]

