FROM ubuntu
ADD ./launcher ./
RUN ["ls", "-l"]
ENTRYPOINT ["./launcher","run" ]