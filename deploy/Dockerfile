FROM alpine:3.14.0
FROM python:3

RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir grpcio && \
    pip install --no-cache-dir grpcio-tools && \
    pip install --no-cache-dir requests

COPY start.sh ./app/
COPY configs ./app/
ADD bin ./app/bin

WORKDIR app

ENTRYPOINT ./start.sh