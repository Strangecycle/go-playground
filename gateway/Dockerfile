FROM golang
MAINTAINER Coven
WORKDIR /work/playground/gateway
COPY . .
EXPOSE 5000
RUN chmod 777 /work/playground/gateway/gateway
ENTRYPOINT ["/work/playground/gateway/gateway"]