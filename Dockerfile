# FROM golang:1.19
# Useing alpine to reduce image size by about 50% ~200Mb
FROM golang:1.19-alpine

WORKDIR /usr/src/app

ENV PORT 4000
ENV ENV production
ENV SUBGRAPH_SECRET some-secret-value

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN pwd
RUN ls -la
# RUN go build -v -o /usr/local/bin/app ./...
RUN go build -v -o /usr/local/bin/app .

EXPOSE 4000

CMD ["app"]