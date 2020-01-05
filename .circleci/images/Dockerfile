FROM golang:1.8

RUN mkdir -p /ProjectReferral/Get-me-in

COPY . src/github.com/ProjectReferral/Get-me-in

WORKDIR src/github.com/ProjectReferral/Get-me-in

RUN go get -v -t -d ./account-service/...
RUN go build -v ./account-service/cmd

# RUN go run ./account-service/cmd/main.go

CMD ["/ProjectReferral/Get-me-in/account-service/cmd ."]
