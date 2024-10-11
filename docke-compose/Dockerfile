# Etapa de build
FROM golang:1.23 AS build
WORKDIR /app

# Copiar arquivos do projeto
COPY . .

# Definir o diretório onde o main.go está localizado
WORKDIR /app/temperatura_por_cep/cmd

# Compilar o binário
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/cloudrun

# Etapa final
FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT [ "./cloudrun" ]