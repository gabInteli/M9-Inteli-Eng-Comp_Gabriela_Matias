# Primeira etapa: construir o servidor API
ARG GO_VERSION=1.21.6
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build-api
WORKDIR /app

# Copie e baixe as dependências do módulo
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte do servidor API e construa
COPY api ./api
RUN CGO_ENABLED=0 go build -o /bin/server ./api

# Segunda etapa: construir o publisher (pub)
FROM golang:${GO_VERSION} AS build-pub
WORKDIR /app

# Copie e baixe as dependências do módulo
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte do publisher
COPY pub .

# Construa o publisher
RUN go build -o /bin/pub

# Terceira etapa: construir o subscriber (sub)
FROM golang:${GO_VERSION} AS build-sub
WORKDIR /app

# Copie e baixe as dependências do módulo
COPY go.mod go.sum ./
RUN go mod download

# Copie o código-fonte do subscriber
COPY sub .

# Construa o subscriber
RUN go build -o /bin/sub

# Quarta etapa: construir a imagem final
FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
COPY --from=build-api /bin/server /bin/
COPY --from=build-pub /bin/pub /bin/
COPY --from=build-sub /bin/sub /bin/

# Configure o ambiente e exponha a porta
EXPOSE 1883

# Defina o ponto de entrada para o servidor API
ENTRYPOINT ["/bin/server"]
