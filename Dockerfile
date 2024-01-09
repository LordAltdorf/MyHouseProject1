FROM golang:1.21-alpine
RUN go build -o awesomeProject4 ./main.go


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o awesomeProject4 ./main.go

CMD ["./awesome_Project4"]