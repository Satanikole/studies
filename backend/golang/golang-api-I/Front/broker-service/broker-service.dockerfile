# base go image (copia todo o codigo em go para um conteiner)
FROM golang:1.18-alpine AS goCode

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api 

RUN chmod +x /app/brokerApp

#build a tiny docker image (copia o executável do container acima para um menor)
FROM alpine:latest

RUN mkdir /app

COPY --from=goCode /app/brokerApp /app

CMD [ "/app/brokerApp" ]