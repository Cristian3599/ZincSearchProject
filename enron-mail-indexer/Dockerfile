FROM golang:latest

WORKDIR /app

RUN apt-get update && apt-get install -y wget && \
    wget http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz

COPY . .

RUN go main.go

CMD ["./main"]
