FROM ubuntu:22.04

RUN apt-get update -qq

COPY . /app
RUN make /app
CMD python /app/app.py
