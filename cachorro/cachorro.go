package cachorro

import "fmt"

// A função Idade recebe a idade do cachorro em anos e retorna a idade equivalente em anos humanos.
func Idade(anos int) int {
    return anos * 7
}

func main(){
    anos := 3
    idade := Idade(anos)
    fmt.Println(idade)
}