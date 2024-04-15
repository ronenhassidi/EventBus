FROM debian:bullseye-slim

RUN apt-get update

RUN apt-get install ca-certificates -y

ADD kafka /usr/local/bin

RUN chmod 755 /usr/local/bin/kafka

RUN mkdir -p /etc/kafka

ENTRYPOINT ["/bin/bash", "-c"]

CMD ["/usr/local/bin/kafka"]
