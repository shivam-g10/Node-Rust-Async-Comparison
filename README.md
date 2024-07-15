# Server for load testing NodeJS vs Rust
## Setup
Rust
```shell
cd rust_server
cargo build --release
```
Node
```shell
cd node_server
pnpm install
```
PM2
```shell
npm install -g pm2
```
Postgres with docker
```shell
sudo docker compose -f dev_infra/docker-compose.yml up -d
```
Apache Bench for load testing
```shell
sudo apt-get update
sudo apt-get install apache2-utils
```
## Run
```shell
pm2 start pm2.config.js
```
## Load test
Node
```shell
# 10,000 requests with 1000 concurrent
ab -n 10000 -c 1000 http://localhost:3000/users/1
```
Rust
Node
```shell
# 10,000 requests with 1000 concurrent
ab -n 10000 -c 1000 http://localhost:3000/users/1
```
## Results
### RAM
Start up
```shell
[PM2] App [node_server] launched (1 instances)
[PM2] App [rust_server] launched (1 instances)
┌────┬────────────────────┬──────────┬──────┬───────────┬──────────┬──────────┐
│ id │ name               │ mode     │ ↺    │ status    │ cpu      │ memory   │
├────┼────────────────────┼──────────┼──────┼───────────┼──────────┼──────────┤
│ 0  │ node_server        │ fork     │ 0    │ online    │ 0%       │ 63.0mb   │
│ 1  │ rust_server        │ fork     │ 0    │ online    │ 0%       │ 4.5mb    │
└────┴────────────────────┴──────────┴──────┴───────────┴──────────┴──────────┘
```
Node Running
```shell
┌────┬────────────────────┬──────────┬──────┬───────────┬──────────┬──────────┐
│ id │ name               │ mode     │ ↺    │ status    │ cpu      │ memory   │
├────┼────────────────────┼──────────┼──────┼───────────┼──────────┼──────────┤
│ 0  │ node_server        │ fork     │ 0    │ online    │ 0%       │ 105.2mb  │
│ 1  │ rust_server        │ fork     │ 0    │ online    │ 0%       │ 4.6mb    │
└────┴────────────────────┴──────────┴──────┴───────────┴──────────┴──────────┘
```
Rust Running
```shell
┌────┬────────────────────┬──────────┬──────┬───────────┬──────────┬──────────┐
│ id │ name               │ mode     │ ↺    │ status    │ cpu      │ memory   │
├────┼────────────────────┼──────────┼──────┼───────────┼──────────┼──────────┤
│ 0  │ node_server        │ fork     │ 0    │ online    │ 0%       │ 31.8mb   │
│ 1  │ rust_server        │ fork     │ 0    │ online    │ 100%     │ 39.2mb   │
└────┴────────────────────┴──────────┴──────┴───────────┴──────────┴──────────┘
```

5min after tests
```shell
┌────┬────────────────────┬──────────┬──────┬───────────┬──────────┬──────────┐
│ id │ name               │ mode     │ ↺    │ status    │ cpu      │ memory   │
├────┼────────────────────┼──────────┼──────┼───────────┼──────────┼──────────┤
│ 0  │ node_server        │ fork     │ 0    │ online    │ 0%       │ 36.2mb   │
│ 1  │ rust_server        │ fork     │ 0    │ online    │ 0%       │ 35.6mb   │
└────┴────────────────────┴──────────┴──────┴───────────┴──────────┴──────────┘
```

10,000 concurrent requests
```shell
┌────┬────────────────────┬──────────┬──────┬───────────┬──────────┬──────────┐
│ id │ name               │ mode     │ ↺    │ status    │ cpu      │ memory   │
├────┼────────────────────┼──────────┼──────┼───────────┼──────────┼──────────┤
│ 0  │ node_server        │ fork     │ 3    │ online    │ 0%       │ 101.9mb  │
│ 1  │ rust_server        │ fork     │ 3    │ online    │ 0%       │ 319.6mb  │
└────┴────────────────────┴──────────┴──────┴───────────┴──────────┴──────────┘
```

### Single request
Node
```shell
Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /users/1
Document Length:        140 bytes

Concurrency Level:      1
Time taken for tests:   0.038 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      348 bytes
HTML transferred:       140 bytes
Requests per second:    26.08 [#/sec] (mean)
Time per request:       38.344 [ms] (mean)
Time per request:       38.344 [ms] (mean, across all concurrent requests)
Transfer rate:          8.86 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    1   0.0      1       1
Processing:    37   37   0.0     37      37
Waiting:       35   35   0.0     35      35
Total:         38   38   0.0     38      38
```
Rust
```shell
Benchmarking localhost (be patient).....done


Server Software:        
Server Hostname:        localhost
Server Port:            3001

Document Path:          /users/1
Document Length:        142 bytes

Concurrency Level:      1
Time taken for tests:   0.011 seconds
Complete requests:      1
Failed requests:        0
Total transferred:      251 bytes
HTML transferred:       142 bytes
Requests per second:    94.46 [#/sec] (mean)
Time per request:       10.586 [ms] (mean)
Time per request:       10.586 [ms] (mean, across all concurrent requests)
Transfer rate:          23.15 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.0      0       0
Processing:    10   10   0.0     10      10
Waiting:       10   10   0.0     10      10
Total:         10   10   0.0     10      10
```
### 500 concurrent requests
Node
```shell
Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /users/1
Document Length:        140 bytes

Concurrency Level:      500
Time taken for tests:   8.748 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3480000 bytes
HTML transferred:       1400000 bytes
Requests per second:    1143.14 [#/sec] (mean)
Time per request:       437.391 [ms] (mean)
Time per request:       0.875 [ms] (mean, across all concurrent requests)
Transfer rate:          388.49 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   5.8      0      36
Processing:    29  425  63.1    419     584
Waiting:       14  425  63.5    419     584
Total:         50  427  60.9    420     600

Percentage of the requests served within a certain time (ms)
  50%    420
  66%    434
  75%    461
  80%    471
  90%    515
  95%    534
  98%    557
  99%    568
 100%    600 (longest request)
```

Rust
```shell
Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3001

Document Path:          /users/1
Document Length:        142 bytes

Concurrency Level:      500
Time taken for tests:   3.379 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      2510000 bytes
HTML transferred:       1420000 bytes
Requests per second:    2959.05 [#/sec] (mean)
Time per request:       168.973 [ms] (mean)
Time per request:       0.338 [ms] (mean, across all concurrent requests)
Transfer rate:          725.31 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    4   6.5      1      36
Processing:    27  158  17.6    160     202
Waiting:       27  157  17.1    159     201
Total:         60  162  14.6    162     220

Percentage of the requests served within a certain time (ms)
  50%    162
  66%    167
  75%    169
  80%    171
  90%    178
  95%    184
  98%    192
  99%    199
 100%    220 (longest request)
```
### 1000 concurrent requests
Node
```shell
Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3000

Document Path:          /users/1
Document Length:        140 bytes

Concurrency Level:      1000
Time taken for tests:   11.857 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      3480000 bytes
HTML transferred:       1400000 bytes
Requests per second:    843.35 [#/sec] (mean)
Time per request:       1185.747 [ms] (mean)
Time per request:       1.186 [ms] (mean, across all concurrent requests)
Transfer rate:          286.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  100 384.0      0    3048
Processing:    70  585 875.3    417    8743
Waiting:       34  584 875.3    416    8743
Total:        146  685 1114.9    419    9767

Percentage of the requests served within a certain time (ms)
  50%    419
  66%    436
  75%    468
  80%    485
  90%   1010
  95%   1651
  98%   6390
  99%   7686
 100%   9767 (longest request)
```
Rust
```shell
Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3001

Document Path:          /users/1
Document Length:        142 bytes

Concurrency Level:      1000
Time taken for tests:   3.367 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      2510000 bytes
HTML transferred:       1420000 bytes
Requests per second:    2969.97 [#/sec] (mean)
Time per request:       336.704 [ms] (mean)
Time per request:       0.337 [ms] (mean, across all concurrent requests)
Transfer rate:          727.99 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    9  20.4      1      97
Processing:    41  304  46.0    316     364
Waiting:       40  303  45.7    314     363
Total:        133  313  30.8    318     392

Percentage of the requests served within a certain time (ms)
  50%    318
  66%    326
  75%    330
  80%    333
  90%    340
  95%    343
  98%    350
  99%    373
 100%    392 (longest request)
```

### 10000 concurrent
Node
```shell
Benchmarking localhost (be patient)
apr_socket_recv: Connection reset by peer (104)
Total of 7972 requests completed
```
Rust
```shell
Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            3001

Document Path:          /users/1
Document Length:        142 bytes

Concurrency Level:      10000
Time taken for tests:   50.814 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      25100000 bytes
HTML transferred:       14200000 bytes
Requests per second:    1967.95 [#/sec] (mean)
Time per request:       5081.426 [ms] (mean)
Time per request:       0.508 [ms] (mean, across all concurrent requests)
Transfer rate:          482.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   54 167.2      0     745
Processing:   299 4783 849.0   5037    5417
Waiting:       54 4783 849.0   5037    5416
Total:        799 4837 706.4   5041    5632

Percentage of the requests served within a certain time (ms)
  50%   5041
  66%   5095
  75%   5136
  80%   5161
  90%   5231
  95%   5322
  98%   5354
  99%   5378
 100%   5632 (longest request)
 ```