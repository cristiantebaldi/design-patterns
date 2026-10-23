# Abstract Factory em Go

Adaptacao do [exemplo TypeScript](../node/abstractFactory/abstractFactory.ts).
Uma fabrica cria uma **familia de produtos relacionados**: voo, hotel e transfer.
O sistema de reservas depende da interface da fabrica, nao dos produtos concretos.

## Executar

Na pasta `abstract-factory-go`, usando Go 1.27.1 ou superior:

```sh
go run .
go test ./...
go vet ./...
```

## Estrutura e papeis

- [travel/interfaces.go](travel/interfaces.go): produtos abstratos e `PackageFactory`.
- [travel/domestic.go](travel/domestic.go): produtos e fabrica da familia domestica.
- [travel/international.go](travel/international.go): produtos e fabrica da familia internacional.
- [travel/booking.go](travel/booking.go): cliente `BookingSystem`, pedido e pacote criado.
- [main.go](main.go): composicao e exibicao dos dois exemplos.

## A forma Go

```go
booking := travel.NewBookingSystem(travel.DomesticFactory{})
pkg := booking.CreateTravelPackage(travel.PackageRequest{
    Origin:        "New York",
    Destination:   "Los Angeles",
    Airline:       "Delta",
    HotelCity:     "Los Angeles",
    HotelLocation: "Downtown",
    TransferType:  "Bus",
})
fmt.Println(pkg.Flight.Details())
```

Trocar `DomesticFactory{}` por `InternationalFactory{}` troca a familia inteira,
sem alterar o sistema de reservas.

- Interfaces sao satisfeitas implicitamente; nao existe `implements`.
- As fabricas usam `struct{}` e receptores por valor, pois nao possuem estado.
  Nao precisam de construtores que apenas retornariam structs vazias.
- `PackageRequest` substitui seis argumentos posicionais por campos nomeados.
- O cliente retorna o pacote; somente o exemplo imprime seus detalhes.
- Como no TypeScript, os produtos apenas descrevem os dados: nao fazem reservas,
  chamadas externas ou validacao de destinos. Por isso, nao retornam erros artificiais.
- Os metodos `Details()` continuam iguais nas duas familias para manter o exemplo
  original. Os testes verificam tambem os tipos concretos, nao apenas o texto.

As interfaces dos produtos possuem a mesma assinatura neste exemplo didatico.
Em Go, portanto, elas sao estruturalmente equivalentes; nomes diferentes nao
impedem a troca manual de produtos. A fabrica fornece os produtos da familia
correta, mas nao constitui uma restricao nominal entre esses tipos.

**Factory Method x Abstract Factory:** no exemplo anterior, a fabrica criava um
processador de pagamento. Aqui, cada fabrica cria varios tipos de produtos
relacionados, mantendo a escolha da familia em um unico lugar.
