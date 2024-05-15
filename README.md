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

### Arrays
```
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr) // Output: [1 2 3]
}
```

### Slices
```
package main

import "fmt"

func main() {
    slice := []int{1, 2, 3}
    slice = append(slice, 4) // adiciona 4 ao slice
    fmt.Println(slice) // Output: [1 2 3 4]
}
```

### Maps 
```
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2}
    m["c"] = 3
    fmt.Println(m) // Output: map[a:1 b:2 c:3]
}
```

### Principais Funções Built-in
#### make() e new()
```
package main

import "fmt"

func main() {
    p := new(int)    // aloca memória para um int, inicializa com zero
    *p = 100
    fmt.Println(*p) // Output: 100

    s := make([]int, 5) // cria um slice com capacidade e comprimento de 5
    fmt.Println(s) // Output: [0 0 0 0 0]
}
```

#### panic() e recover()
```
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    
    panic("Panic!") // Output: Recovered in f Panic!
}
```

#### defer
```
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello") // Output: hello \n world
}
```

### Structs
```
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p) // Output: {Alice 30}
}
```

### Métodos
```
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    p.Greet() // Output: Hello, my name is Alice and I am 30 years old.
}
```

### Criar e Escrever em Arquivos
```
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    file.WriteString("Hello, World!")
}
```

# Exemplo de Array de JSONs
```
[
    {"name": "Alice", "age": 30},
    {"name": "Bob", "age": 25}
]
```

### Passo 1: Definir a Estrutura
```
package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```

### Passo 2: Declarar a Variável
```
var people []Person
```

### Passo 3: Deserializar o JSON
```
func main() {
    jsonData := `[
        {"name": "Alice", "age": 30},
        {"name": "Bob", "age": 25}
    ]`

    var people []Person
    err := json.Unmarshal([]byte(jsonData), &people)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(people)
    // Output: [{Alice 30} {Bob 25}]
}
```

- Definindo a Estrutura (struct): A estrutura Person tem dois campos, Name e Age, que correspondem às chaves no JSON. As tags json:"name" e json:"age" indicam ao pacote encoding/json como mapear os campos da estrutura para os campos no JSON.

- Declarando a Variável: var people []Person declara uma variável people que é um slice de Person.

- Deserializando o JSON: json.Unmarshal converte os dados JSON em uma estrutura Go. O primeiro argumento é o JSON em formato de []byte e o segundo é um ponteiro para a variável onde os dados deserializados serão armazenados.

### Trabalhando com a Variável
```
for _, person := range people {
    fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
```
Depois de deserializar os dados JSON, você pode trabalhar com a variável people como faria com qualquer outro slice; cada elemento do slice people e imprimir os valores de Name e Age.
