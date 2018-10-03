package domain

import (
	"database/sql"
	"fmt"

	//Postgres driver
	_ "github.com/lib/pq"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "DadosExtraidos"
)

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

	ct := fmt.Sprintf("CREATE TABLE IF NOT EXISTS HistoricoConsumo (" +
		"CPFConsumidor VARCHAR(255) NULL," +
		"Privado VARCHAR(255) NULL," +
		"Incompleto VARCHAR(255) NULL," +
		"DtUltimaCompra VARCHAR(255) NULL," +
		"TicketMedio VARCHAR(255) NULL," +
		"TicketUltimaCompra VARCHAR(255) NULL," +
		"CNPJLojaFrequente VARCHAR(255) NULL," +
		"CNPJLojaUltimaCompra VARCHAR(255) NULL," +
		"FlCPFConsumidorValido bool NULL," +
		"FlCNPJLojaFrequenteValido bool NULL," +
		"FlCNPJLojaUltimaCompraValido bool NULL)")

	ok, err := ExecutarQuery(ct)

	if ok == true {
		fmt.Println("Tabela HistoricoConsumo criada!")
	}
	if err != nil {
		fmt.Println(err)
	}
}

//ExecutarQuery executa as requisições ao banco de dados
func ExecutarQuery(query string) (bool, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return false, err
	}
	defer db.Close()

	db.Exec(query)
	return true, nil
}
