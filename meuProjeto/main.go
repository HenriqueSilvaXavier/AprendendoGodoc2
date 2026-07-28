package main

import (
    "fmt"
    "meuProjeto/cachorro" // Importa o pacote local
)

func main(){
    anos := 3
    idade := cachorro.Idade(anos)
    fmt.Println(idade)
}