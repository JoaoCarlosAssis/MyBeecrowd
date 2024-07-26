peca1 = input()
peca2 = input()

peca1 = peca1.split()
peca2 = peca2.split()


total_peca1 = int(peca1[1]) * float(peca1[2])
total_peca2 = int(peca2[1]) * float(peca2[2])

total_pagar = total_peca1 + total_peca2

print(f"VALOR A PAGAR: R$ {total_pagar:.2f}")
