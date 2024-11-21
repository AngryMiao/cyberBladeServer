FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY  _build/customer /app/main
COPY  certs /app/certs
COPY  80.json /app/80.json

CMD ["./main", "-conf", "config/customer.yaml"]
