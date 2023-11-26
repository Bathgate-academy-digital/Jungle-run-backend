FROM ubuntu

RUN apt-get update && apt-get install curl -y
RUN curl -OL https://github.com/Bathgate-academy-digital/Jungle-run-backend/releases/download/0.1.0-alpha/jungle-run
RUN chmod a+x jungle-run

EXPOSE 8080

CMD ./jungle-run
