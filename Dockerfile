FROM ubuntu:22.04

COPY ./ /run
WORKDIR /run
EXPOSE 8080
RUN chmod u+x ./serverguardian
CMD ["/run/serverguardian"]