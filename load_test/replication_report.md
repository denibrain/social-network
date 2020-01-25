# Отчет по Реплицации

Испытательный стенд был расширен еще двумя дополнительными узалми
MySQL master и slave (Ранее MySQL был установлен на узле приложения)

Каждый узел Ubuntu с 512Mb RAM, 25Gb HDD

Используюя инструкцию 
https://www.digitalocean.com/community/tutorials/how-to-set-up-master-slave-replication-in-mysql


## Результаты эксперимента

В качестве инструменты был использован httpref

Когда в базе было всего 250,000 то разница скорости ответа сервера была
невелика и если использовалась только одна нода для записи и чтения 
то результаты даже были лучше

### Только мастер:
 httperf --server 146.185.177.181 --port 8181 --wsesslog=100,0.1,./urls.txt --rate=1000
httperf --client=0/1 --server=146.185.177.181 --port=8181 --uri=/ --rate=1000 --send-buffer=4096 --recv-buffer=16384 --wsesslog=100,0.100,./urls.txt
httperf: warning: open file limit > FD_SETSIZE; limiting max. # of open files to FD_SETSIZE
Maximum connect burst length: 5

Total: connections 100 requests 102400 replies 102400 test-duration 271.717 s

Connection rate: 0.4 conn/s (2717.2 ms/conn, <=100 concurrent connections)
Connection time [ms]: min 266867.1 avg 269793.0 max 271658.1 median 0.0 stddev 1056.4
Connection time [ms]: connect 68.6
Connection length [replies/conn]: 1024.000

Request rate: 376.9 req/s (2.7 ms/req)
Request size [B]: 78.0


### Мастер для записи, и слейв для чтения:
httperf --server 146.185.177.181 --port 8181 --wsesslog=100,0.1,./urls.txt --rate=1000
httperf --client=0/1 --server=146.185.177.181 --port=8181 --uri=/ --rate=1000 --send-buffer=4096 --recv-buffer=16384 --wsesslog=100,0.100,./urls.txt
httperf: warning: open file limit > FD_SETSIZE; limiting max. # of open files to FD_SETSIZE
Maximum connect burst length: 3

Total: connections 100 requests 102400 replies 102400 test-duration 276.257 s

Connection rate: 0.4 conn/s (2762.6 ms/conn, <=100 concurrent connections)
Connection time [ms]: min 271461.3 avg 274307.2 max 276191.0 median 0.0 stddev 1010.0
Connection time [ms]: connect 68.5
Connection length [replies/conn]: 1024.000

Request rate: 370.7 req/s (2.7 ms/req)
Request size [B]: 78.0

Reply rate [replies/s]: min 304.6 avg 372.0 max 396.6 stddev 16.9 (55 samples)
Reply time [ms]: response 166.8 transfer 0.6

## 10,000,000 записей

## Подключение только к мастеру

➜  load_test git:(master) ✗ httperf --server 146.185.177.181 --port 8181 --wsesslog=10,0.1,./urls.txt --rate=10
httperf --client=0/1 --server=146.185.177.181 --port=8181 --uri=/ --rate=10 --send-buffer=4096 --recv-buffer=16384 --wsesslog=10,0.100,./urls.txt
httperf: warning: open file limit > FD_SETSIZE; limiting max. # of open files to FD_SETSIZE
^CMaximum connect burst length: 1

Total: connections 10 requests 10369 replies 10364 test-duration 1274.598 s

Connection rate: 0.0 conn/s (127459.8 ms/conn, <=10 concurrent connections)
Connection time [ms]: min 0.0 avg 0.0 max 0.0 median 0.0 stddev 0.0
Connection time [ms]: connect 73.0
Connection length [replies/conn]: 0.000

Request rate: 8.1 req/s (122.9 ms/req)
Request size [B]: 77.0

Reply rate [replies/s]: min 0.0 avg 8.0 max 56.2 stddev 11.6 (254 samples)
Reply time [ms]: response 1128.6 transfer 0.3
Reply size [B]: header 124.0 content 4784.0 footer 2.0 (total 4910.0)
Reply status: 1xx=0 2xx=10364 3xx=0 4xx=0 5xx=0

## Подключение на чтение и к мастеру и к реплике

Total: connections 10 requests 9023 replies 9013 test-duration 779.785 s

Connection rate: 0.0 conn/s (77978.5 ms/conn, <=10 concurrent connections)
Connection time [ms]: min 0.0 avg 0.0 max 0.0 median 0.0 stddev 0.0
Connection time [ms]: connect 71.0
Connection length [replies/conn]: 0.000

Request rate: 11.6 req/s (86.4 ms/req)
Request size [B]: 77.0

Reply rate [replies/s]: min 0.0 avg 11.6 max 44.4 stddev 8.3 (155 samples)
Reply time [ms]: response 733.3 transfer 0.3
Reply size [B]: header 124.0 content 4750.0 footer 2.0 (total 4876.0)
Reply status: 1xx=0 2xx=9013 3xx=0 4xx=0 5xx=0




