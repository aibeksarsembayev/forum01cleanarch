FROM golang:1.19

WORKDIR /app

COPY . .

RUN go build -o forum ./cmd

ENV PORT 3000

EXPOSE $PORT

VOLUME [ "/app/data" ]

CMD ["./forum"]