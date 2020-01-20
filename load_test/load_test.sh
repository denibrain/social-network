# 10 сессий с паузами 0.1 между запросами, которые берутся из файла /tmp/test.url.txt
httperf --server localhost --port 8181 --wsesslog=10,0.1,./urls.txt --rate=3