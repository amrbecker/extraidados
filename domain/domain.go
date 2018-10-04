package domain

import (
	"database/sql"
	"fmt"
	"log"

	//Postgres driver
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "DadosExtraidos"
)

// HistoricoConsumo estrutura de dados
type HistoricoConsumo struct {
	CPFConsumidor,
	Privado,
	Incompleto,
	DtUltimaCompra,
	TicketMedio,
	TicketUltimaCompra,
	CNPJLojaFrequente,
	CNPJLojaUltimaCompra string
	FlCPFConsumidorValido,
	FlCNPJLojaFrequenteValido,
	FlCNPJLojaUltimaCompraValido bool
}

//TestarConexao com o conexão com o banco de dados
func TestarConexao() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado com sucesso!")
}

//CriarTabela insere a tabela no banco de dados se ainda não existir
func CriarTabela() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := `CREATE TABLE IF NOT EXISTS HistoricoConsumo (
		ID SERIAL PRIMARY KEY,
		CPFConsumidor VARCHAR(80) NULL,
		Privado VARCHAR(80) NULL,
		Incompleto VARCHAR(80) NULL,
		DtUltimaCompra VARCHAR(80) NULL,
		TicketMedio VARCHAR(80) NULL,
		TicketUltimaCompra VARCHAR(80) NULL,
		CNPJLojaFrequente VARCHAR(80) NULL,
		CNPJLojaUltimaCompra VARCHAR(80) NULL,
		FlCPFConsumidorValido bool,
		FlCNPJLojaFrequenteValido bool,
		FlCNPJLojaUltimaCompraValido bool)`

	db.Exec(query)

	if err == nil {
		fmt.Println("Tabela HistoricoConsumo criada!")
	}
}

//InserirDados recebe uma lista de históricos para inserir no banco de dados
func InserirDados(h []*HistoricoConsumo) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO HistoricoConsumo VALUES(?,?,?,?,?,?,?,?,?,?,?)")

	for i := 1; i < len(h); i++ {
		stmt.Exec(
			"h[i].CPFConsumidor",
			"h[i].Privado",
			"h[i].Incompleto",
			"h[i].DtUltimaCompra",
			"h[i].TicketMedio",
			"h[i].TicketUltimaCompra",
			"h[i].CNPJLojaFrequente",
			"h[i].CNPJLojaUltimaCompra",
			h[i].FlCPFConsumidorValido,
			h[i].FlCNPJLojaFrequenteValido,
			h[i].FlCNPJLojaUltimaCompraValido)
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
}
