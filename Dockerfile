FROM ubuntu:latest
#ENV HELLO_IDENTIFIER 1
COPY main /main
CMD ["/main"]
