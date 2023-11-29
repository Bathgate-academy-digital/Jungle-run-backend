FROM ubuntu:latest

COPY jungle-rush /jungle-rush
RUN chmod a+x /jungle-rush

EXPOSE 8080

CMD ./jungle-rush
