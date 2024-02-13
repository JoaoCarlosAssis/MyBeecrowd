n = int(input())
numbers_array = []

for _ in range(n):
    numbers_str = input()
    numbers = list(map(int, numbers_str.split()))
    smaller, bigger = min(numbers), max(numbers)
    numbers_array.append([smaller, bigger])


def calcular_soma_impares(smaller, bigger):
    return sum(num for num in range(smaller + 1, bigger) if num % 2 != 0)


somas_impares = [calcular_soma_impares(smaller, bigger) for smaller, bigger in numbers_array]

for numero in somas_impares:
    print(numero)

