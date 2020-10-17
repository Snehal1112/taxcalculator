FROM golang:latest AS builder
LABEL maintainer="Snehal Dangroshiya <snehaldangroshiya@gmail.com>"

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o taxcal
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o taxcal .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache --upgrade bash

WORKDIR /root/
COPY --from=builder /app/taxcal .
COPY --from=builder /app/run.sh .

RUN cat run.sh
RUN ls -l

EXPOSE 8773
CMD ["./run.sh"]