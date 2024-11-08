# Go Host Lookup

Este projeto em Go realiza consultas de IP e servidores de nomes (NS) para um host especificado pelo usuário. Ele utiliza a biblioteca padrão `net` do Go para realizar as consultas e a biblioteca `cli` para capturar os argumentos.

## Funcionalidades

1. **buscarIps**: Consulta e exibe todos os endereços IP associados a um host especificado.
2. **buscarServer**: Consulta e exibe todos os servidores de nomes (NS) associados a um host especificado.

## Instalação

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório.
3. go mod tidy

```bash
git clone https://github.com/StyvenNoronha/seu-repositorio.git
cd seu-repositorio
```

## Instale as dependências do projeto.
```bash
go mod
```
## Uso

- Busca por Ip
```bash
linha-de-comando.exe ip --host amazon.com.br
```

- Busca por Servidor
```bash
linha-de-comando.exe server --host amazon.com.br
```


