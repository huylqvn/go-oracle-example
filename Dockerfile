FROM golang:1.16 AS builder

RUN apt-get update
RUN apt-get install -y pkg-config && \
    apt-get install -y libaio1 && \
    apt-get install unzip

ENV CLIENT_FILENAME instantclient_12_1.zip
COPY /oracle/${CLIENT_FILENAME} .
COPY /oracle/oci8.pc /usr/lib/pkgconfig/oci8.pc

ENV LD_LIBRARY_PATH /usr/lib:/usr/local/lib:/usr/instantclient_12_1

# to build the application with mattn/go-oci8, it is necessary to extract all files, including the SDK.
RUN unzip ${CLIENT_FILENAME} -d /usr &&  \
    ln -s /usr/instantclient_12_1/libclntsh.so.12.1 /usr/instantclient_12_1/libclntsh.so && \
    ln -s /usr/instantclient_12_1/libclntshcore.so.12.1 /usr/instantclient_12_1/libclntshcore.so && \
    ln -s /usr/instantclient_12_1/libocci.so.12.1 /usr/instantclient_12_1/libocci.so

RUN go get -u github.com/mattn/go-oci8

ENV GO111MODULE=on
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN go build -v -o server cmd/server/main.go

EXPOSE 8080
CMD ["/app/server"]
