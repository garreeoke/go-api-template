FROM golang:alpine

RUN mkdir /sample-go-api
ADD . /sample-go-api/
WORKDIR /sample-go-api
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o sample-go-api .
RUN adduser -S -D -H -h /sample-go-api sample-go-api-user
USER sample-go-api-user
CMD ["./sample-go-api"]