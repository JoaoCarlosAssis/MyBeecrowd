values = input()

values = (int(values.split(' ')[0]), int(values.split(' ')[1]))

for i in range(1, values[1]+1):
  if i % values[0] == 0:
    print(f"{i}")
  else:
    print(f"{i}", end=" ")
  
    