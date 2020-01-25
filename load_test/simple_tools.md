CURL
================
curl -i 'https://otus.ru/'

### Цикл на 10 повторений
for i in {1..10}; do curl 'https://yandex.ru/' --write-out '%{http_code} %{time_total}\n' --silent --output /dev/null; done;

### Цикл на 10 повторение и сбор статистики по всем
for i in {1..10}; do curl 'https://yandex.ru/' --write-out '%{http_code} %{time_total}\n' --silent --output /dev/null >> /tmp/test.1.txt; done; cat /tmp/test.1.txt; awk '{ total += $2; count++ } END { print "\nMean: " total/count "\n" }' /tmp/test.1.txt


Apache.Benchmark
================
sudo apt install apache2-utils
### 30 запросов в 10 потоков с keepAlive метод GET
ab -n 30 -c 10 -k -m GET https://yandex.ru/


HTTPerf
================
sudo apt install httperf
httperf --server yandex.ru --num-conn 10 --rate=3
echo '/images/' > /tmp/test.url.txt 
cat /tmp/test.url.txt

### 10 сессий с паузами 0.1 между запросами, которые берутся из файла /tmp/test.url.txt
httperf --server yandex.ru --wsesslog=10,0.1,/tmp/test.url.txt --rate=3