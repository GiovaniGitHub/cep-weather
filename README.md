# Go Expert - LAB Desafio 

## Descrição
O sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin).
## Conteúdo

1. [Como Rodar o Projeto](#como-rodar-o-projeto)
2. [Testes Automatizados](#testes-automatizados)
3. [Docker](#docker)
4. [Deploy no Google Cloud Run](#deploy-no-google-cloud-run)

## Como Rodar o Projeto
### Ambiente de Desenvolvimento

1. Certifique-se de ter o Golang 1.19 instalado em sua máquina.
2. Clone o repositório: `git clone https://github.com/GiovaniGitHub/cep-weather.git`
3. Navegue até o diretório do projeto: `cd cep-weather`
4. Crie um .env a partir do .env.template e altere o campo
**Exemplo**
```bash
WEB_SERVER_PORT=8080
ENVIRONMENT=development
URL_BASE=http://localhost
```

### Rodar Sem Docker
 - Requisitos basicos:
   - Golang v1.19

```bash
    make run # Roda o projeto
```

```bash
    make test # Executa os testes
```

```bash
    make all # Executa os testes e o projeto
```

### Rodar Com Docker
 - Requisitos basicos:
   - Docker
- Altere o campo **CONTAINER_NAME** no arquivo makefile 

```bash
    make build-docker # Cria a imagem docker do projeto
```

```bash
    make run-docker # Roda o projeto
```

### Ambiente de Produção

1. Em producao a aplicacao esta rodando no Google Cloud Run.
2. Segue um teste possivel

```bash
    curl -H "Content-Type: application/json" https://cep-weather-prqp4ppyua-uc.a.run.app/cep/70070080
```
Onde a saída possível é:
```json
{"temp_C":"28","temp_F":"82.40","temp_K":"301.00"}
```