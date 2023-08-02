FROM golang:latest

WORKDIR /usr/src/awesomeProject

RUN go version
ENV GOPAHT=/

COPY ./ ./

RUN go mod download
RUN go build -o bin/bin cmd/main.go


EXPOSE 8080

CMD [ "./bin/bin" ]