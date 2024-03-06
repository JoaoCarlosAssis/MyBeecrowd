def soma_nao_multiplos_de_13(numero_01, numero_02):
    soma = 0

    if numero_01 > numero_02:
       numero_01, numero_02 = numero_02, numero_01
    
    for num in range(numero_01, numero_02 + 1):
       if num % 13 != 0:
          soma += num


    return soma



numero_01 = int(input())
numero_02 = int(input())

resultado = soma_nao_multiplos_de_13(numero_01, numero_02)


print(resultado)


