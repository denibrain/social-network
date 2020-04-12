# Galera cluster

Для создания кластера из баз MySQL с технологией Galera
было использовано Perciona Xtra DB Cluster

Для того чтобы обеспечить приложению доступ к кластеры
принято решение исползовать ProxySQL и ETCD
В паре они обеспечивают достаточно стабильное подключение к БД

## Стенд

Сервер приложения был использован без изменеий из пред. домашней работы
Сервер данный: Ubuntu 18.04 2GB RAM 50GB HDD

На сервере данных было развернуто 5 докер контейнеров 
ETCD - Сервис провайдер нод кластера
ProxySQL - Прокси слой для приложения для единого подключения к БД
3 ноды Percona XtraDB Cluster

Детали: https://github.com/denibrain/social-network/tree/master/cluster

## Под нагрузкой

коммнада запуска 
httperf --server 146.185.177.181 --port 8181 --wsesslog=10,0.1,./urls.txt --rate=10

#№# Переключение нод, сессия 1

Request rate: 25.9 req/s (38.5 ms/req)
Request size [B]: 77.0

Reply rate [replies/s]: min 0.2 avg 26.0 max 52.4 stddev 13.9 (70 samples)
Reply time [ms]: response 282.6 transfer 0.6
Reply size [B]: header 124.0 content 6863.0 footer 2.0 (total 6989.0)
Reply status: 1xx=0 2xx=9024 3xx=0 4xx=0 5xx=111


### Переключение нод, сессия 2

Request rate: 26.2 req/s (38.2 ms/req)
Request size [B]: 77.0

Reply rate [replies/s]: min 0.0 avg 26.3 max 50.6 stddev 14.4 (136 samples)
Reply time [ms]: response 271.2 transfer 0.6
Reply size [B]: header 124.0 content 6867.0 footer 2.0 (total 6993.0)
Reply status: 1xx=0 2xx=17672 3xx=0 4xx=0 5xx=208

### Нарощение кол-ва нод с 1 по 3 с интервалом 1-2 мин

Request rate: 28.2 req/s (35.5 ms/req)
Request size [B]: 77.0

Reply rate [replies/s]: min 4.6 avg 28.0 max 48.0 stddev 11.3 (33 samples)
Reply time [ms]: response 253.0 transfer 0.5
Reply size [B]: header 124.0 content 6900.0 footer 2.0 (total 7026.0)
Reply status: 1xx=0 2xx=4636 3xx=0 4xx=0 5xx=40


### Уменшение кол-во нод с 3 до 1 с интервалом 1-2 мин

Reply rate [replies/s]: min 11.2 avg 27.7 max 45.6 stddev 9.1 (14 samples)
Reply time [ms]: response 244.2 transfer 0.5
Reply size [B]: header 124.0 content 6698.0 footer 2.0 (total 6824.0)
Reply status: 1xx=0 2xx=1878 3xx=0 4xx=0 5xx=62

CPU time [s]: user 10.93 system 61.24 (user 15.1% system 84.5% total 99.6%)
Net I/O: 180.3 KB/s (1.5*10^6 bps)


### Сессия без изменений

Reply rate [replies/s]: min 2.8 avg 29.6 max 52.0 stddev 11.3 (78 samples)
Reply time [ms]: response 235.5 transfer 0.5
Reply size [B]: header 124.0 content 6882.0 footer 2.0 (total 7008.0)
Reply status: 1xx=0 2xx=11522 3xx=0 4xx=0 5xx=142

CPU time [s]: user 55.23 system 336.03 (user 14.0% system 85.5% total 99.5%)
Net I/O: 205.3 KB/s (1.7*10^6 bps)

### Вывод

500 ошибки приложения видны только под нагрузкой приложения.
500 возникали (с той же скоростью) когда кластер не изменял конфигурацию
При смене состояния кластера с достаточным запасом ресурсов они не были обнаружены
Возможно можно улучшить чтобы нивелировать и их.