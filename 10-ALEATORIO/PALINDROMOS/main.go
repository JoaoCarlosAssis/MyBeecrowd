import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sair = false
	for !sair {
		menu(&sair)
	}
}

func menu(sair *bool) {
	fmt.Print("Digite uma opção:\n1-Verificar Palindromo\n2-Sair\n")
	var option string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite uma palavra: ")
	palavra, _ := reader.ReadString('\n')
	palavra = strings.TrimSpace(palavra)

	switch option {
	case "1":
		verificar_palindromo()
	case "2":
		*sair = true
	}
}

func verificar_palindromo() {
	var palavra string
	fmt.Print("Digite uma palavra: ")
	fmt.Scanln(&palavra)

	palavra = strings.ReplaceAll(palavra, " ", "")
	palavra = strings.ToLower(palavra)

	palavra_inversa := []string{}

	for _, letra := range palavra {
		letra_string := string(letra)
		palavra_inversa = append([]string{letra_string}, palavra_inversa...)
	}

	palavra_invertida := strings.Join(palavra_inversa, "")

	if palavra_invertida == palavra {
		fmt.Printf("%s é um palindromo!\n", palavra)
	} else {
		fmt.Printf("%s não é um palindromo!\n", palavra)
	}

}
