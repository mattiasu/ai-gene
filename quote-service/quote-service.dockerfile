FROM alpine:latest

RUN mkdir /app

COPY quoteServiceApp /app

CMD ["/app/quoteServiceApp"]