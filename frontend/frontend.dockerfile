FROM alpine:latest

RUN mkdir /app

COPY front-endApp /app

CMD ["/app/front-endApp"]