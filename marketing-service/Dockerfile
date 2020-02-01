FROM golang:1.8

RUN rm -d src/github.com

RUN mkdir -p /ProjectReferral/Get-me-in

COPY . src/github.com/ProjectReferral/Get-me-in

WORKDIR src/github.com/ProjectReferral/Get-me-in

# RUN go get -v -t -d ./marketing-service/...
RUN go get -u ./marketing-service/...
RUN go build -v ./marketing-service/cmd

# RUN go run ./account-service/cmd/main.go

CMD ["/ProjectReferral/Get-me-in/marketing-service/cmd ."]
