# Load Test

Перед нагрузочным тестом было сгенерено 253000 пользователей с различными данными
Дамп: users.sql

Также было сгенерировано порядка 1000 адресов (urls.txt)
Для нагрузочного теста был использован httpref

Запуск httpref: httperf --server localhost --port 8181 --wsesslog=10,0.1,./urls.txt --rate=3 

## Условия тестирования

Веб сервер и httpref были на одной машине.
MySQL локальный

На каджый запрос к веб-серверу происходил запрос к базе
вида SELECT * FROM users WHERE name LIKE '<query-parameter>%' LIMIT 20

## Результаты до добавления индексов

httperf --client=0/1 --server=localhost --port=8181 --uri=/ --rate=3 --send-buffer=4096 --recv-buffer=16384 --wsesslog=10,0.100,./urls.txt
Maximum connect burst length: 1

Total: connections 10 requests 10240 replies 10240 test-duration 425.419 s

Connection rate: 0.0 conn/s (42541.9 ms/conn, <=10 concurrent connections)
Connection time [ms]: min 415642.2 avg 419421.6 max 422751.6 median 0.0 stddev 2356.0
Connection time [ms]: connect 1.0
Connection length [replies/conn]: 1024.000

Request rate: 24.1 req/s (41.5 ms/req)
Request size [B]: 72.0

Reply rate [replies/s]: min 13.2 avg 24.1 max 26.2 stddev 1.6 (85 samples)
Reply time [ms]: response 309.0 transfer 0.0
Reply size [B]: header 124.0 content 3105.0 footer 2.0 (total 3231.0)
Reply status: 1xx=0 2xx=10240 3xx=0 4xx=0 5xx=0

CPU time [s]: user 63.31 system 353.44 (user 14.9% system 83.1% total 98.0%)
Net I/O: 77.6 KB/s (0.6*10^6 bps)

Errors: total 0 client-timo 0 socket-timo 0 connrefused 0 connreset 0
Errors: fd-unavail 0 addrunavail 0 ftab-full 0 other 0

Session rate [sess/s]: min 0.00 avg 0.02 max 1.20 stddev 0.14 (10/10)
Session: avg 1.00 connections/session
Session lifetime [s]: 419.4
Session failtime [s]: 0.0

## Результаты до добавления индексов

httperf --client=0/1 --server=localhost --port=8181 --uri=/ --rate=3 --send-buffer=4096 --recv-buffer=16384 --wsesslog=10,0.100,./urls.txt
Maximum connect burst length: 1

Total: connections 10 requests 10240 replies 10240 test-duration 108.726 s

Connection rate: 0.1 conn/s (10872.6 ms/conn, <=10 concurrent connections)
Connection time [ms]: min 105681.9 avg 105723.0 max 105774.6 median 0.0 stddev 31.3
Connection time [ms]: connect 0.0
Connection length [replies/conn]: 1024.000

Request rate: 94.2 req/s (10.6 ms/req)
Request size [B]: 72.0

Reply rate [replies/s]: min 69.0 avg 95.4 max 98.0 stddev 6.1 (21 samples)
Reply time [ms]: response 2.8 transfer 0.0
Reply size [B]: header 124.0 content 3105.0 footer 2.0 (total 3231.0)
Reply status: 1xx=0 2xx=10240 3xx=0 4xx=0 5xx=0

CPU time [s]: user 28.05 system 80.25 (user 25.8% system 73.8% total 99.6%)
Net I/O: 303.6 KB/s (2.5*10^6 bps)

Errors: total 0 client-timo 0 socket-timo 0 connrefused 0 connreset 0
Errors: fd-unavail 0 addrunavail 0 ftab-full 0 other 0

Session rate [sess/s]: min 0.00 avg 0.09 max 0.00 stddev 0.00 (10/10)
Session: avg 1.00 connections/session
Session lifetime [s]: 105.7
Session failtime [s]: 0.0

## Итоги

Добавление индексов дало сокращение времени ответа в 4 раза.