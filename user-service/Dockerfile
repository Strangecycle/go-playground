FROM golang
MAINTAINER Coven
WORKDIR /work/playground/user-service
COPY . .
EXPOSE 5001
RUN chmod 777 /work/playground/user-service/user-service
ENTRYPOINT ["/work/playground/user-service/user-service"]