FROM centos

WORKDIR /go_web

COPY . /go_web

EXPOSE 5000

ENTRYPOINT ["./gin_web"]