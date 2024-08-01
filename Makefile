# Nome do executável
EXEC = docker-compose-generator

# Diretório de origem
SRC_DIR = ./cmd/docker-compose-generator

# Arquivos de origem
SRC = $(SRC_DIR)/main.go

# Alvo padrão
all: build

# Compila o executável
build:
	go build -o $(EXEC) $(SRC)

# Executa o programa
run: build
	./${EXEC}
