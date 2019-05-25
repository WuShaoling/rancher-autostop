FROM ubuntu:16.04

ADD auto-stop /usr/bin/

RUN chmod +x /usr/bin/auto-stop

CMD ["auto-stop"]