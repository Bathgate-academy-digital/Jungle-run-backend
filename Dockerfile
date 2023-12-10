FROM busybox:latest

COPY jungle-rush /jungle-rush
COPY data/bad_words.txt /data/bad_words.txt
COPY data/classes.txt /data/classes.txt
RUN chmod a+x /jungle-rush

EXPOSE 8080

CMD ./jungle-rush
