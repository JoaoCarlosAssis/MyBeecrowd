quantidade = int(input())
cont = 0
lista_numeros = []
while cont < quantidade:
  numeros = input().split()
  lista_numeros.append(numeros)
  cont+=1

for numero in lista_numeros:
  if int(numero[1]) == 0:
    total = "divisao impossivel"
  else:
    total = int(numero[0]) / int(numero[1])
  print(total)