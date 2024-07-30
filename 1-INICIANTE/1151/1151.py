n = int(input())

numeros = [0, 1]
while len(numeros)<n:
  next_value = numeros[-1] + numeros[-2]
  numeros.append(next_value)

for i in numeros:
  if i == numeros[-1]:
    print(f"{i}")
  else: 
    print(f"{i}", end=" ")