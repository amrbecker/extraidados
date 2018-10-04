#  Serviço para extração de dados de um arquivo

Serviço para extração de dados de um arquivo txt ou csv e persistência em banco de dados relacional.
### Requisitos do serviço:
  - Ler os dados de um arquivo txt/csv
  - Separar os dados em atributos do histórico de consumo
  - Validar se os CPFs e CNPJs são válidos
  - Persistir as informações no banco de dados relacional

### Mapeamento de negócio
O arquivo recebido possui dados históricos de consumidores e suas compras. 
Os dados extraídos do arquivo são: 

  - CPF do consumidor
  - CNPJ da Loja mais frequente
  - CNPJ da Loja da ultima compra
  - Valor ticket da última compra
  - Valor ticket médio
  - Data da última compra
  - Informação privada
  - Informação incompleta

Após a extração dos dados, é verificado a validade dos CPFs e CNPJs.

Os dados processados são atribuídos à entidade HistoricoConsumo, e estas informações são inseridas na tabela HistoricoConsumo, do banco de dados DadosExtraidos. 

### Dicionário de Dados
##### Entidade: HistoricoConsumo
| Atributo | Tipo | Descrição |
| ------ | ------ | ------ |
| CPFConsumidor | VARCHAR | CPF do consumidor |
| Privado | VARCHAR | Informação privada |
| Incompleto | VARCHAR | Informação incompleta |
| DtUltimaCompra | VARCHAR | Data da última compra |
| TicketMedio | VARCHAR | Valor ticket médio |
| TicketUltimaCompra | VARCHAR | Valor ticket da última compra |
| CNPJLojaFrequente | VARCHAR | CNPJ da loja mais frequente |
| CNPJLojaUltimaCompra | VARCHAR | CNPJ da loja da ultima compra |
| FlCPFConsumidorValido | Bool | Validade do CPF |
| FlCNPJLojaFrequenteValido | Bool | Validade do CNPJ |
| FlCNPJLojaUltimaCompraValido | Bool | Validade do CNPJ |

### Tecnologias utilizadas

O serviço utiliza os seguintes projetos de código aberto para funcionar corretamente:

* [Golang](https://golang.org) - O Go é uma linguagem de programação de código aberto que facilita a criação de software simples, confiável e eficiente.
* [PostgreSQL](https://www.postgresql.org) - Sistema gerenciador de banco de dados objeto relacional, desenvolvido como projeto de código aberto.
* [Docker](https://docs.docker.com/) - Plataforma Open Source escrito em Go, que facilita a criação e administração de ambientes isolados.
* [Docker Compose](https://docs.docker.com/compose) - Ferramenta para a criação e execução de múltiplos containers de aplicação. 


### Instalação

O projeto requer [Docker](https://docs.docker.com/) v1.13+ e a ferramenta Docker Compose.

Para criar e inicializar os contêineres, execute o comando.

```sh
$ docker-compose up
```

Licença
----

MIT

