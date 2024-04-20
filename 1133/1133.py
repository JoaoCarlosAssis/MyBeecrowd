x = int(input())
y = int(input())

x, y = min(x, y), max(x, y)

for number in range(x + 1, y):
  if number % 5 == 2 or number % 5 == 3:
    print(number)

