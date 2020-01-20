

def char_range(c1, c2):
    for c in range(ord(c1), ord(c2)+1):
        yield chr(c)


fileOut = open('urls.txt', 'wt')
for c1 in char_range('А', 'Я'):
    for c2 in char_range('а', 'я'):
        fileOut.write(f'/?name={c1}{c2}\n')


fileOut.close()
