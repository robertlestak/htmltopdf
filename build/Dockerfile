FROM ubuntu:16.04

RUN mkdir -p /root/go/src/github.com/robertlestak/htmltopdf
WORKDIR /root/go/src/github.com/robertlestak/htmltopdf

RUN apt-get update && apt-get install -y curl vim xvfb \
    libfontconfig wkhtmltopdf git build-essential

RUN curl https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz -o go1.10.2.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz && rm go1.10.2.linux-amd64.tar.gz

ENV PATH="${PATH}:/usr/local/go/bin:/root/go/bin"

COPY . .

RUN go get -u golang.org/x/vgo

RUN CC=gcc vgo build cmd/htmltopdf/htmltopdf.go

CMD ["./htmltopdf"]
