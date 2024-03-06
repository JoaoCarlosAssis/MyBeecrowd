numbers_array = []

while True:
  numbers_str = input()
  numbers = list(map(int, numbers_str.split()))
  smaller, bigger = min(numbers), max(numbers)
  
  if smaller <= 0 or bigger <= 0:    
    break
  else:
    sum_itens = sum(range(smaller, bigger + 1)) 
    numbers_array.append([smaller, bigger, sum_itens])
    
    
for items in numbers_array:
   sequence = ' '.join(map(str, range(items[0], items[1] +1)))
   print(f"{sequence} Sum={items[2]}")

