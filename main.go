package main

import (
	arquivo "github.com/amrbecker/extraidados/arquivo"
	domain "github.com/amrbecker/extraidados/domain"
)

func main() {
	//Iniciando conexão com o banco de dados
	domain.TestarConexao()
	//Verifica se a tabela já existe e a cria
	domain.CriarTabela()
	//Processamento dos dados e persistência
	arquivo.ExtrairDados()
}
