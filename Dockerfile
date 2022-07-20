FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/sudharshan3/GO-lms
RUN cd /build && git clone https://github.com/sudharshan3/GO-lms.git

RUN cd /build/GO-lms && go build

EXPOSE 4000

ENTRYPOINT [ "/build/GO-lms/main" ]