

names = set([])
fin = open('names.csv', 'rt')
for line in fin:
    name, lastName = line.strip().split(',')
    if name == 'DELETED':
        continue

    if len(name) > 60:
        print("Long name", name)
        continue

    names.add(name[0:2])


fin.close()

fileOut = open('urls.txt', 'wt')
for prefix in names:
    fileOut.write(f'/?name={prefix}\n')

fileOut.close()
