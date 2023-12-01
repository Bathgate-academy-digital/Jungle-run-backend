FROM busybox:latest

COPY jungle-rush /jungle-rush
COPY bad_words.txt /bad_words.txt
COPY classes.txt /classes.txt
RUN chmod a+x /jungle-rush

EXPOSE 8080

CMD ./jungle-rush
