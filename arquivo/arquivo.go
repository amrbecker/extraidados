package arquivo

import (
	"bufio"
	"os"
	"strings"

	"github.com/amrbecker/extraidados/domain"
	"github.com/amrbecker/extraidados/util"
)

//ExtrairDados lê arquivo e salva no banco
func ExtrairDados() {

	// Lendo o arquivo
	file, err := os.Open("arquivo/teste.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(file)

	h := []*domain.HistoricoConsumo{}

	// Loop over all lines in the file.
	for scanner.Scan() {

		lines := scanner.Text()
		campos := strings.Fields(lines)

		ht := new(domain.HistoricoConsumo)
		ht.CPFConsumidor = campos[0]
		ht.Privado = campos[1]
		ht.Incompleto = campos[2]
		ht.DtUltimaCompra = campos[3]
		ht.TicketMedio = campos[4]
		ht.TicketUltimaCompra = campos[5]
		ht.CNPJLojaFrequente = campos[6]
		ht.CNPJLojaUltimaCompra = campos[7]
		ht.FlCPFConsumidorValido = util.CPF(campos[0])
		ht.FlCNPJLojaFrequenteValido = util.CNPJ(campos[6])
		ht.FlCNPJLojaUltimaCompraValido = util.CNPJ(campos[7])

		h = append(h, ht)

		domain.InserirDados(h)
	}
}
