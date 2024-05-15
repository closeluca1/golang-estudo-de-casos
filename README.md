# golang-estudo-de-casos

É interessante acompanhar este documento auxiliado do __**Go Playground**__ para obter retornos. 
>
>https://go.dev/play/
>

B<i>o</i>a jornada!
#

### Compreender a Sintaxe Básica da Golang
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

- package main: Declara o pacote principal do programa.
- import "fmt": Importa o pacote fmt, que contém funções para formatação de E/S.
- func main(): Define a função principal do programa.
- fmt.Println("Hello, World!"): Imprime "Hello, World!" na saída padrão.

### Compreender Conceitos Básicos Envolvidos na Go, Como Ponteiros de Memória
Ponteiros são variáveis que armazenam endereços de memória. Em Go, os ponteiros são usados de forma semelhante ao C.
```
package main

import "fmt"

func main() {
    var x int = 10
    var p *int = &x  // p é um ponteiro para x

    fmt.Println("x:", x)   // Output: x: 10
    fmt.Println("p:", p)   // Output: p: (endereço de x)
    fmt.Println("*p:", *p) // Output: *p: 10

    *p = 20  // altera o valor de x para 20 através do ponteiro
    fmt.Println("x após *p = 20:", x) // Output: x após *p = 20: 20
}
```

### Utilizar as Estruturas Básicas da Linguagem, Como Declaração de Variáveis
Variáveis podem ser declaradas de várias maneiras em Go;
```
package main

import "fmt"

func main() {
    var a int = 1   // declaração explícita com tipo
    var b = 2       // inferência de tipo
    c := 3          // declaração curta com inferência de tipo
    var d, e int = 4, 5 // declaração múltipla

    fmt.Println(a, b, c, d, e) // Output: 1 2 3 4 5
}
```
