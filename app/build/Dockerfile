FROM golang:latest

ARG http_port=8090

COPY build/_output/bin /bin

# for main web interface:
EXPOSE ${http_port}

CMD ["/bin/http-server"]
