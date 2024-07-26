i = 0
count = 0
for _ in range(0, 11):
    j = 1 + count
    for _ in range(1, 4):
        formatted_i = "{:.1f}".format(i) if i % 1 != 0 and i < 1.8 else "{:.0f}".format(i)
        formatted_j = "{:.1f}".format(j) if j % 1 != 0 else "{:.0f}".format(j)
        print("I={} J={}".format(formatted_i, formatted_j))
        j += 1
    i += 0.2
    count += 0.2

