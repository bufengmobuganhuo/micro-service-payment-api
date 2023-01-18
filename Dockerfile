FROM alpine
ADD micro-api /micro-api
ENTRYPOINT [ "/micro-api" ]
