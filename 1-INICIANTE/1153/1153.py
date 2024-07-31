def calcular_fatorial(n):
    fatorial = 1
    for i in range(1, n + 1):
        fatorial *= i
    return fatorial

n = int(input())
resultado = calcular_fatorial(n)
print(f"{resultado}")

