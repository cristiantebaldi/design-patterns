# Builder em Go

Este projeto adapta os exemplos TypeScript para Go, mostrando duas abordagens:

- **Sem Director**: o cliente configura o builder etapa por etapa e chama `Build`.
- **Com Director**: o `CheckoutDirector` aplica receitas prontas reutilizando um builder.

Os builders de cartão e PIX validam seus campos obrigatórios ao construir o pedido. `Build` retorna um erro em vez de lançar uma exceção. O exemplo inclui descontos percentuais calculados no total; os dados de pagamento são apenas fictícios.

## Executar

Na pasta `builder-go`:

```sh
go run .
```

Para validar os pacotes:

```sh
go test ./...
```

## Estrutura

- `order`: tipos do pedido e dados de pagamento.
- `builder`: builders concretos e validação/construção do pedido.
- `director`: receitas de checkout que configuram os builders.
