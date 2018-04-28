FROM resin/rpi-raspbian:latest

ENTRYPOINT []

ADD tempbot tempbot
RUN chmod +x tempbot

CMD ["./tempbot"]
