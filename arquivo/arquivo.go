package arquivo

import (
	"bufio"
	"fmt"
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

	}

	for i := 1; i < len(h); i++ {
		q := fmt.Sprintf("INSERT INTO HistoricoConsumo VALUES('%s','%s','%s','%s','%s','%s','%s','%s',%t,%t,%t)",
			h[i].CPFConsumidor, h[i].Privado, h[i].Incompleto, h[i].DtUltimaCompra,
			h[i].TicketMedio, h[i].TicketUltimaCompra, h[i].CNPJLojaFrequente,
			h[i].CNPJLojaUltimaCompra, h[i].FlCPFConsumidorValido,
			h[i].FlCNPJLojaFrequenteValido, h[i].FlCNPJLojaUltimaCompraValido)

		ok, err := domain.ExecutarQuery(q)

		if ok == true {
			fmt.Println("Dados inseridos na tabela para o CPF: " + h[i].CPFConsumidor)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
