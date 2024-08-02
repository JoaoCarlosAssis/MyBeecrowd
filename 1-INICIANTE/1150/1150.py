x = int(input())


z = int(input())

while z <= x:
    z = int(input())


count = 0
soma = x

while soma <= z:
    soma += x + 1
    count += 1


print(soma, count)
