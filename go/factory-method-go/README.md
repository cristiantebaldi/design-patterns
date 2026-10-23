# Factory Method em Go

O projeto reproduz os exemplos de pagamento em TypeScript usando Go e mostra as duas abordagens:

- **Sem Factory Method**: o código cliente cria cada processador e o entrega ao checkout.
- **Com Factory Method**: o checkout recebe uma fábrica e pede a ela a criação do processador.

As fábricas escondem os detalhes de configuração de cada integração. Os valores de credenciais usados são apenas exemplos fictícios; não os substitua por credenciais reais no código.

## Executar

Na pasta `factory-method-go`:

```sh
go run .
```

Para executar os testes:

```sh
go test ./...
```

## Estrutura

`payment` define as interfaces `Processor` e `ProcessorFactory`, além dos checkouts. Cada pacote em `providers` implementa o processador e sua fábrica.
