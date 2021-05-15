FROM golang
MAINTAINER Coven
WORKDIR /work/playground/file-service
COPY . .
EXPOSE 5001
RUN chmod 777 /work/playground/file-service/file-service
ENTRYPOINT ["/work/playground/file-service/file-service"]