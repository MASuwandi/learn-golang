
### Running app

# Install Dependency
- install all dependencies:
    go get .

- install all dependencies recursively:
    go get ./...
    go get -d ./...

... is a special pattern, tells to go down recursively.

- compile file -> generate binary file
go build filename

- run app
./filename

# Run on Dev env
go run filename

# Error not in GOROOT
old GOPATH, new Go Modules
- activate old GOPATH
go env -w GO111MODULE=off

- activate Go Modules
go env -w GO111MODULE=on

# Notes
- data type 32 -> 8
127 to -128

Progress: To Do, On Progress, Done

# ----- Go Topic / Subject -----: - Progress ------#
# -------------------------------------------------#
#  1. Dasar                : Done        : 61 chap #
#  2. Modules              : Done        :  8 chap #
#  3. Unit Test            : Done        : 16 chap #
#  4. Goroutines           : Done        : 28 chap #
#  5. Context              : Done        :  9 chap #
#  6. DB MySequel          : Done        : 15 chap #
#  7. Embed                : Done        :  8 chap #
#  8. Web                  : Done        : 29 chap #
#  9. HttpRouter           : Done        : 11 chap #
# 10. JSON                 : Done        : 11 chap #
# 11. RESTful API          : Done        : 25 chap #
# 12. Dependency Injection : To Do       : 19 chap #
# 13. Logging              : To Do       : 13 chap #
# 14. Generic              : To Do       : 17 chap #
# 15. DB Migration         : To Do       : 17 chap #
# 16. Validation           : To Do       : 26 chap #
#     GRaphQL              : To Do                 #
# -------------------------------------------------#


<details>
<summary>### 1. Go Dasar</summary>
<br>
<!-- EOL 245 -->

### 1. Go Dasar

## Data Type
# Number
- int       : - -> +
- uint      : 0 -> +
- float     : decimal

# Alias
- byte     : uint
- rune     : int

# Boolean
# String


## Variable
// Declare
var name string  

// Declare with type
const num2 uint8 = 32

// Declare and assign w/o keyword and type
num3 := 33

// Declare mutiple
var (
    firstName = "firstName"
    lastName = "lastName"
)

name = "masuwandi"  // Assign

## Conversion Data Type
var nilai32 int32 = 128
var nilai64 int64 = int64(nilai32)

var name = "Baim"
var i = name[2]
var eString = string(i)

## Type Dclarations
type NIM string
byte() // check avail alias

## Math Opration
*, /, %

// Augmented Assignments
// += , -= , *= , /= , %=

// Unary Operator
// ++ , -- , - , + , !

## Boolean Opration
&&, ||, !

## Array Data Type
const values = [2]string // empty array
const values = [2]int{1,2}
len(array/string)

## Slice Data Type

## Map Data Type

## If Expression

## Switch Expression

## Loop

# For Range

## Function

# Return Value

# Multiple Values

# Named Return Values
variable return value

## Variadic Function
spread like param and argument
can send slice with ...

## Function Value
function as variable

## Function as Parameter

## Type Declaration

## Anonymous Function

## Closure

## Defer, Panic & Recover

## Struct

# Struct Method

## Interface

# Interface Kosong

## error Interface

## Type Asserions

## Pointer
ex:
    var pX *int = &x;
	var y int = *pX

ex: pX *int = pX type pointer
- * = Arterisk
- * placed next to a type
- it modified the type meaning that var is now a pointer.
- pX log as address = 0x1400018e008

ex: var y int = *pX
- * alone or dereference
- grab value of pointer / the thing pointed to by pX
- y log as value = 1

ex: &x = address of x
- & = Ampersand
- & = the address of
- byte integer (pointer)

- stack of memoery allocation
- active frame
- heap allocation

# Best Practices Consideration / When:
- pass / copy big data/object/property size
- called a lot
- performance / optimize memory
- update state

- (v int) {} -> immutability
- (p *int) {} -> efficiency

# In Function

# In Method

## Activate GOPATH

# GOPATH

## Error Package and Import

## Package and Import

## Access Modifier
- Capital letter for export / global var

## Package Initialization

## Package os

## Package flag

## Package strings

## Package strconv

## Package math

## Package container list

## Package container ring

## Package sort

## Package time

## Package reflec

## Package regexp

### Noted Feature:

- Slice Data Type
- Map Data Type
- For Range
- Multiple Return Values
- Named Return Values
- Variadic Function
- Defer, Panic & Recover
- Type Dclarations
- Struct
- Struct Method
- Interface
- Interface Kosong
- Nil
- error Interface
- Type Asserions
- Pointer
- Pointer In Function Param
- Pointer In Struct Method 
- Activate GOPATH
- GOPATH
- Error Package and Import
- Package and Import
- Access Modifier
- Package Initialization
- Blank Identifier
var, import

# Package
os, flag, strings, strconv, math,
container list, container ring,
sort, time, reflec, regexp

## Enable gopls
cd current project root
- empty go work
go work init

- go.work with contains
go work use ./repo1 ./repo2

### 1. end
<!-- SOL 50 -->
</details>

<details>
<summary>### 2. Go Modules</summary>
<br>
<!-- EOL 445 -->

### 2. Go Modules

like NPM

# Enable go modules
- activate Go Modules
go env -w GO111MODULE=on

## Create / Membuat Module
go mod init namaModule tanpa protocol
ex: github/user/repo name

## Release / Rilis Module
create Tag on repo Git / VCS
git tag v1.0.0
git push origin v1.0.0

## Add / Menambah Module / Dependency
go get namaModule tanpa protocol
add dependency to go mod
ex: go get github/MASuwandi/go-say-hello
import namaVar source
ex: import go_say_hello github/MASuwandi/go-say-hello

## Upgrade Module
add another tag
git tag v1.5.0
git push origin v1.5.0

## Upgrade Dependency
update tag version in go mod to newest version
ex: v1.5.0
go get

## Major
version structure v[major].[minor].[patch]
ex: v1.2.1

# Criteria
- break / error happen if app using new package

# Ex:
- change of function name, parameter, type

# If major changes occur
# on package
- change module name
ex: github/MASuwandi/go-say-hello/v2

- upgrade major version tag
reset the mayor and patch to 0
ex: v2.0.0
git tag v2.0.0
git push origin v2.0.0

# on app that use the package
- update use package in go mod
1. remove old package
2. check module name in go mod package tag v2
ex: code -> tag -> v2.0.0
3. go get new module name from go mod package
github/MASuwandi/go-say-hello/v2
4. update module name use by app
ex: import go_say_hello github/MASuwandi/go-say-hello/v2
5. fix the break or error
- add args

### 2. end
<!-- SOL 314 -->
</details>

<details>
<summary>### 3 Go-Lang Unit Test</summary>
<br>
<!-- EOL 314 -->

### 3 Go-Lang Unit Test
## Pengenalan Software Testing
# Test Cat:
- UI Test / End to End
need Front End

- Service Test / Integration Test
service level / one Back End service
need DB

- Unit Test (this session)
small component
Function, Method, Branch
Controller, Service, Repository
when one component tested,
    response from other component set to fixed
- all scenario: success and fail
- 1 to 4 func

## Pengenalan testing Package
- testting.T : struct T
unit test
- testting.M : sMruct M
life cycle testing
- testting.B : struct B
benchmarking / kecepatan kode program

## Membuat Unit Test
# Format Rule:
Test File Name 
end with _test.go

Func Name
TestFuncName

# add module to workspace
go work use ..

# Running:
go test

- show function
go test -v
go test -v -run funcName // prefix like LIKE %

- run from parent
go test -v ./...

## Menggagalkan Test
t.Fail   : fail and continue execution
t.FailNow: fail, stop and go to next function

- Best Practice
t.Error  : log error, then call Fail()
t.Fatal  : log error, then call FailNow()

## Assertions
assertion library Testify
- assert    : Fail, call Fail()
- require    : Fail, call FailNow()

## Skip Test
skip for UT which can only be run on specific os

## Before dan After Test
Goals:
- reduce repetition
- run in one package

## Sub Test
t.Run()

go test -v -run TestSubTest/SubUT

## Table Test
Konsep menyediakan data slice,
contains param dan expect hasil UT
- lalu iterasi slice dengan sub test

## Mock
object
UT ideal nya:
not running db, ex API / 3rd party API.
dynamic response

ex: query db, service layer / business logic, repo layer / bridge to db
need: kontrak / interface

## Benchmark
speed test app
repetitive iteration call in a duration

using:
package.struct
testing.B

mekanism:
create iteration N amount

## Membuat Benchmark
# Run Benchmark:
- run test and benchmark
go test -v -bench=.

- run benchmark without test
go test -v -run=NotMatchUnitTest -bench=.

- run specific benchmark without test
go test -v -run=NotMatchUnitTest -bench=BenchmarkHelloWorld2

- run all test and benchmark all package
go test -v -run=NotMatchUnitTest -bench=. ./...

- run all benchmark all package without test
go test -v -run=ExcludeTest -bench=. ./...

## Sub Benchmark

## Table Benchmark

### 3. end
<!-- SOL 255 -->
</details>

<details>
<summary>### 4. Goroutines</summary>
<br>
<!-- EOL 605 -->

### 4. Goroutines

##  Pengenalan Concurrency dan Parallel Programming
# Parallel
memecahkan masalah dengan cara membaginya menjadi lebih kecil, dan dijalankan secara bersamaan pada waktu yang bersamaan

menjalankan aplikasi di waktu yg sama
ex: office, editor, browser, spotify bersamaan

# Process vs Thread
Process                 Thread
eksekusi program        segmen dari process
konsum big ram          kosum small ram
isolate from other      can com to other thread in same process
long to start or stop   short to start or stop

# Concurrency
- menjalankan beberapa pekerjaan secara bergantian.
- can use idle time to work on other task
ex: wait response 1 sec from db, can be use to working on other task
- in parallel need many Thread,
- in concurrency need only few Thread

default Go Lang Concurrency
can also support paralel

# CPU bound
# I/O bound

# Why use golang
low Thread usage because of concurrency

##  Pengenalan Goroutine
-  Goroutines run in thread / light thread manage by Go Runtime.
- size 2kb.

# How Goroutine Work
- Goroutine run by Go Scheduler inside thread,
where amount of thread as much GOMAXPROCS ( CPU core amount )
- Thread manage by Go Scheduler automatically

# How Go Scheduler Work
- G: Goroutine
- M: Thread ( Machine )
- P: Processor

##  Membuat Goroutine
go invoker()

##  Goroutine Sangat Ringan

## Channel
##  Pengenalan Channel
- tempat kom secara synchronous by goroutine
- sender and receiver by difrent goroutine
- alternatif async await

## Karakteristik Channel

##  Membuat Channel
type chan
make()

no one take channel data
- panic: send on closed channel

take empty channel data
- all goroutines are asleep - deadlock!


##  Channel Sebagai Parameter
default pass by reference

##  Channel In dan Out


##  Buffered Channel
- Buffer Capacity
kapasitas antrian di dalam buffer
set 5, receive 5 data di buffer

##  Range Channel
- data yg di terima tidak di ketahui jumlahnya
- dikirim terus oleh pengirim
- tidak jelas akan berhenti menerima

##  Select Channel

##  Default Select

##  Race Condition

## Mutex
##  sync Mutex
struct sync.Mutex
struct suitable with mutex
- locking: 
- unlock: 

race condition solution

##  sync RWMutex
case accessed by many go routine

##  Deadlock
proses goroutine saling menunggu lock,
sehingga goroutine tidak bisa jalan

## Wait Group
##  sync WaitGroup
menunggu proses goroutines selesai
Add(int)    : Increment group proses
Done()      : Decrement group proses
Wait()      : Wait until proses become 0

##  sync Once
run function only once

##  sync Pool
design pattern bernama object pool patern

ex: DB connection, quota of DB connection
safe from race condition

##  sync Map
aman untuk concurrent menggunakan goroutine
safe from race condition

- Store(key, Value)
- Load(key)
- Delete(key)
- Range(func(key, value))

##  sync Cond
locking base condition
use Mutex or RWMutex

##  Atomic
locking bisa menggunakan Atomic

##  time Timer
send event when time expire

## Ticker
##  time Ticker
process every periode
ex: every 5 sec

##  GOMAXPROCS
config Thread amount
set or get Threah amount

### 4. end
<!-- SOL 455 -->
</details>

<details>
<summary>### 5. Context</summary>
<br>
<!-- EOL 650 -->

### 5. Context
## Pengenalan Context
- data yg membawa value, sinyal cancel, sinyal timeout dan sinyal deadline.

- created per request

- make us easier to forward value and signal between process.

## Why Need?
- membatalkan semua proses, cukup mengirim sinyal ke context.

- Hampir semua bagian di Golang memanfaatkan context, seperti database, http server, http client dll.

## Membuat Context
Background()
Todo()

## Parent dan Child Context
impact signal and data: parent and child
immutable: cannot be change

when value added: new child context will be created

## Context With Value
.withvalue(parent, key, value)

tidak bisa mengambil data child

## Context With Cancel
WithCancel(parent)

## Context With Timeout
WithTimeout(parent, duration)

## Context With Deadline
WithDeadline(parent, time)

### 5. end
<!-- SOL 615 -->
</details>

<details>
<summary>### 6. DB MySequel</summary>
<br>
<!-- EOL 842 -->

### 6. DB MySequel
## Pendahuluan
- package database
- standar interface
- akses db apapun dengan kode yg sama
- hanya query yg berbeda

# Cara Kerja:
         call               call            call
Aplikasi ----> DB Interface ----> DB Driver ----> DBMS
purpose:
                                - for connection
                                - translator to

# Install
https://dev.mysql.com/doc/refman/8.0/en/macos-installation.html

# Online resources:
https://dev.mysql.com/doc/
https://www.mysql.com
https://www.oracle.com

# Config
ls /usr/local

- set .zshrc / .bashrc file
export PATH=/usr/local/mysql/bin:$PATH

- reload CLI
source ~/.zshrc

db client:
mysql --version

db server / daemon:
mysqld

# mysql shell
mysql -u user -p


## Menambah Database Driver
List supported DB:
github.com/golang/go/wiki/SQLDrivers

Usage Indicator:
- popularity / users: solve issue

Installation:
go get -u github.com/go-sql-driver/mysql

import _ "github.com/go-sql-driver/mysql" // init

## Membuka Koneksi ke Database
- sql.Open(drive, dataSourceName)
- driver "mysql"
- dataSourceName: username:password@tcp(host:port)/database_name

- default port: 3306

- Connection Leak if not close
- connection hanging even thou app already stop
- max open connection: out of connection quota
Close()


## Database Pooling
- Idle Conns: Jumlah Max Koneksi Idle dalam pool.
Koneksi idle adalah koneksi yang terbuka dan tersedia untuk digunakan, tetapi saat ini tidak sedang digunakan untuk eksekusi query atau transaksi.

Pool koneksi database memungkinkan koneksi untuk tetap terbuka setelah digunakan agar dapat digunakan kembali secara efisien tanpa perlu membuka dan menutup koneksi baru setiap kali diperlukan.

- Open Conns: Jumlah Max Koneksi aktif yg dapat dibuka dan digunakan secara bersamaan dalam pool.

Koneksi aktif adalah koneksi yang sedang digunakan untuk eksekusi query atau transaksi pada saat tertentu. Jumlah koneksi aktif biasanya harus dikontrol agar tidak melebihi batas maksimum yang ditentukan. Jika jumlah koneksi aktif melebihi batas yang ditentukan, maka permintaan koneksi baru harus menunggu hingga koneksi aktif berkurang, yang dapat menyebabkan waktu respons yang lambat dan kinerja aplikasi yang buruk.

Dengan menggunakan db.SetMaxOpenConns, Anda dapat mengontrol berapa banyak koneksi aktif yang diperbolehkan dalam pool koneksi database. Nilai ini harus disesuaikan dengan kapasitas database Anda dan batas maksimum koneksi yang dapat ditangani oleh server database.

- Conn Life time: batas waktu maksimum sebuah koneksi dalam pool.
Batas waktu ini menentukan seberapa lama sebuah koneksi dapat tetap aktif sebelum harus dianggap sudah kadaluarsa dan ditutup.

Ketika aplikasi menggunakan pool koneksi database, koneksi-koneksi tersebut dapat dibuka dan ditutup secara otomatis sesuai permintaan query atau transaksi. Namun, dalam beberapa kasus, server database atau jaringan mungkin dapat menyebabkan koneksi menjadi tidak valid setelah berjalan untuk jangka waktu tertentu. Misalnya, server database dapat melakukan restart atau pemeliharaan rutin yang menyebabkan koneksi menjadi tidak berlaku.

- Conn Idle Time: Jumlah max Idle Time
DB.SetConnMaxIdleTime menetapkan durasi maksimum koneksi dapat menganggur sebelum ditutup. Hal ini menyebabkan sql.DB menutup koneksi yang telah menganggur lebih lama dari durasi yang diberikan.

By default, when an idle connection is added to the connection pool, it remains there until it is needed again. When using DB.SetMaxIdleConns to increase the number of allowed idle connections during bursts of parallel activity, also using DB.SetConnMaxIdleTime can arrange to release those connections later when the system is quiet.

# Config:
(DB)SetMaxIdleConns(number) : max idle con
(DB)SetMaxOpenConns(number) : max open con
(DB)SetConnMaxIdleTime(duration) : con idle exp
(DB)SetConnMaxLifetime(duration) : con duration

## Eksekusi Perintah SQL
untuk oprasi yg tidak membutuhkan hasil
(DB)ExecContext(context, sql, params)

## Query SQL
(DB)QueryContext(context, sql, params)

## Tipe Data Column
# Mapping Tipe Data
Tipe Data DB        Tipe Data GO
------------------------------------
varchar, char     : string
int, bigint       : int32, int64
float, double     : float32, float64
boolean           : bool
date, datetime,
time, timestamp   : time.Time

# error time
?parseTime=true

# error null
# Tipe Data Nullable
Tipe Data GO      Tipe Data Nullable
------------------------------------------
string          : database/sql.NullString
bool            : NullBool
float64         : NullFloat64
int32           : NullInt32
int64           : NullInt64
time.Time       : NullTime

return struct

## SQL Injection
- don't trust user input
- "#" equal to comment
- string concat bad practice

## SQL dengan Parameter
# Qeury
SELECT username
FROM user
WHERE username = ?
AND password = ?
LIMIT 1

#Insert
INSERT INTO user(username, password)
VALUES(?, ?)

db.QueryContext(ctx, query, username, password)
db.ExecContext(ctx, query, username, password)

## Auto Increment
(result)LastInsertId()
in exec

## Prepare Statement
can be manual

// Best practice
insert multiple row with one connection
prepare statement dengan satu koneksi

// How to
(DB)Prepare(context, sql)
database/sql.Stmt
Close() // connection leak

// notes
query dan exec dengan param, tidak ada jaminan insert dengan satu koneksi atau koneksi yg sama

## Database Transaction
- by default, all SQL query that we sent with Golang, will be commit automatically

- feature transaction, so it will not be automatic

- (DB)Begin()
struct

- func(Tx) Commit(), Rollback()

## Repository Pattern
DAO / repository pattern
- in Domain-Driven Design by Eric Evans:
"repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of objeects"

- as bridge between business logic with SQL command to DB

- all SQL command write within Repository,
while business logic in controller / agregator

- SQL won't be in controller

# How it works:
Business --> Repository -----> Entitiy/
Logic            ⬇             Model
                 ⬇
             Repository
             Implementation -> Database

- Repository: Kontrak / Interface
- Repo Imple: Struct
- Entity /
Model /                 -> Struct
Data Representation

# Best Practice:
- One Repo for each table

### 6. end
<!-- SOL 660 -->
<details>


<details>
<summary>### 7. Embed</summary>
<br>
<!-- EOL 853 -->

### 7. Embed
## Pengenalan Embed Package
go > 1.16.0
go version

- membaca isi file pada saat compile time secara otomatis dimasukan isi file ke variable.

import _ "embed"

must be store outside function

## Embed File ke String
//go:embed version.txt
var version string

## Embed File ke []Byte
gambar, video, music dll

## Embed Multiple Files
tipe data embed.FS

## Path Matcher (Multi Feature)
path.Match

use pattern regex
golang.org/pkg/path/#Match

* -> all

## Hasil Embed di Compile
hsil embed oleh package ember, permanent dan data file yg dibaca disimpan dalam binary file golang

- bukan dilakukan secara realtime membaca file yg ada diluar

- jika binary file golang sudah di compile, kita tidak butuh lagi file external, dan jika diubah file externalnya, isi variable nya tidak akan berubah lagi.
karena berada di dalam variable

### 7. end
<!-- SOL 853 -->
</details>

<details>
<summary>### 8. Web</summary>
<br>
<!-- EOL 898 -->

### 8. Web
## Pengenalan Web
# Web:
- Web berisi informasi berbentuk: teks, gambar, audio, video dll.
- Web berjalan di aplikasi Web Server, aplikasi untuk menyimpan dan menyampaikan isi informasi Web.
Butuh Web Server untuk menjalankan web.

# Web Host:
penyedia komputer untuk Web Server
ex: cloud provider

# Domain:
Komputer Web memiliki alamat
Alamat adalah ip address
ex: 172.217.194.94

Nama Domain adalah alamat alias ke ip address
Ask domain server for ip address

# Web Browser:
- aplikasi untuk mengakses Web
- kita bisa mengakses Web secara langsung tanpa Web Browser, namun Web Server akan memberikan informasi bahasa mesin seperti HTML, JavaScript, CSS, Gambar, Video, dan dll.

- Dengan menggunakan Web Browser, semua bahasa mesin tersebut bisa ditampilkan secara visual sehingga kita bisa menyerap informasinya dengan lebih mudah.


## Client dan Server
- Web merupakan aplikasi berbasis Client dan Server.
- arsitektur app menghubungkan dua pihak, sistem client dan sistem server.
- Sistem client dan sistem server yg saling berkomunikasi melalui jaringan komputer, internet atau bisa di komputer yg sama.

# Tugas Client dan Server
# Client:
- mengirim request ke server dan menerima response dari server.
# Server:
- menerima request dari Client, memproses data, dan mengembalikan hasil proses data / mengirim response ke Client.

# Keuntungan Client dan Server
- perubahan aplikasi bisa dilakukan dengan mudah di server, tanpa harus membuat perubahan di client.
- bisa digunakan oleh banyak client pada saat yg bersamaan, meskipun server tidak banyak.
- bisa diakses dari mana saja, asal terhubung satu jaringan dengan server.

# Contoh Client dan Server
- Web merupakan salah satu contoh arsitektur client server.
- Applikasi yg bertugas sebagai Client adalah Web Browser:
Chrome, Firefox, Edge dan Safari.
- Applikasi yg bertugas sebagai Server adalah Web Server, dimana di dalam web server terdapat kode program Web kita.


## Golang Web
- Web API (Backend)
- sudah terdapat package untuk membuat web
package net/http

- tidak perlu library atau framework

- golang web sudah teersedia webserver

# Recommend:
- ketika membuat web skala besar, direkomendasikan menggunakan framewoek karena dipermudah oleh web framework.
- lebih terstruktur

# Cara Kerja
1. Web Browser akan melakukan HTTP Request ke Web Server.
2. Golang menerima HTTP Request, lalu mengeksekusi request tersebut sesuai dengan yg diminta.
3. Setelah melakukan eksekusi request, Golang akan mengembalikan data dan di render sesuai dengan kebutuhanya, misal HTML, CSS, JaavaScript, dll.
4. Golang akan mengembalikan content hasil render tersebut sebagai HTTP Response ke Web Browser.
5. Web Browser menerima content dari Web Server, lalu me-render content tersebut sesuai dengan tipe contentnya.

## Server
- struct yg terdapat di package net/http yg digunakan sebagai representasi Web Server

- untuk membuat web, perlu membuat Server
- saat membuat data server, beberapa hal yg perlu kita tentukan, seperti host, port tempat dimana Web kita berjalan.
2 digit port butuh akses admin.

# Run Server:
- ListenAndServe()

## Handler
Server hanya bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yg masuk ke Server, kita butuh yg Handler

- representasi dalam interface, dalam kontrak terdapat func ServeHTTP(), akan di exec ketika menerima HTTP request.

## ServeMux
like router
implementasi Handler support multiple endpoint

# URL Pattern
- long url priority
- akhiran / , akan menerima awalan yg sama dengan akhiran apapun.
ex:
    endpoint: /images/
    request end: /images/a, /images/a/b

# Recommend:
- unique URL / specific

## Request
- struct representasi HTTP Request
- semua informasi request bisa kita dapatkan di Request.
ex: URL, method, header, body, dll.
URI / URL

## HTTP Test
- package httptest
simulasi test / request
- httptest.NewRequest()
- cookie, header, dll.

- recorder := httptest.NewRecorder()
recorder that the HTTP handler writes to as its http.ResponseWriter, and it captures all of the changes that would have been returned to a client caller. Using this, there’s no need to start your server: just hand the recorder directly to the function and it invokes it the same way it would if the request came in over HTTP. After the handler call, the recorder’s Result call provides the values written to it for checking any behaviors you may need to to assert in the rest of your test.

- handler func
- handler call
- string(body)

## Query Parameter
- query parameter ditempatkan pada URL
ex: ?key=value

request.URL
lots of URL method
Query()

## Header
informasi tambahan yg terdapat
di HTTP Request dan Response

- tidak case sensitive

- dengan browser, biasanya otomatis header ditambahkan oleh browser:
    content-type

# Location in browser
Inspect El -> Network -> Headers

# Request Header
- Request.Header

# Response Header
- Response.Writer.Header()

## Form Post
form bisa submit data, dengan method GET atau POST
- GET, data di query param (URL)
- POST, data di body HTTP request

# How to use
- Rquest.ParseForm() (postForm, error)
- Request.PostForm


## Response Code
- atau status code, status dari request yg dikirim sukses diproses oleh server atau gagal

1. Informational responses (100 - 199)
2. Successful responses (200 - 299)
3. Redirects (300 - 399)
4. Client errors (400 - 499)
5. Server errors (500 - 599)

- developer.mozilla.org/en-US/docs/Web/HTTP/Status

# Mengubah Response Code
- Default 200 OK
- ResponseWriter.WriteHeader(int)
- package http
- github.com/golang/go/blob/master/src/net/http/status.go

## Cookie (feature save state)
# Case:
- server tidak akan menyimpan data untuk mengingat setiap request dari client
- bertujuan agar mudah melakukan scalability di sisi server
- cara server mengingat client? setelah login server harus otomatis tahu jika client sudah login, sehingga request selanjutnya di page atau api yg berbeda, tidak perlu diminta login lagi.
- Untuk hal ini, kita bisa memanfaatkan Cookie.

# What is Cookie and how it works?
- cookie adalah fitur HTTP dimana server bisa memberi response cookie (key-value) dan client akan menyimpan cookie tersebut di browser. (ketika login success)
- request selanjutnya, client akan selalu membawa cookie tersebut secara otomatis.
- dan server akan secara otomatis akan selalu menerima data cookie yg dibawa oleh client setiap kalau client mengirim request.

# Flow:
Client                                  Server
1. request login   -   -   -   -   ->  2. proses success

4. browser save
    cookie        <-   -   -   -   -   3. response cookie

5. next request 
will include cookie -   -   -   -   -> 6. validate cookie

- endpoint / resource
- cookie dibuat di server

# Membuat Cookie
http.SetCookie()

## FileServer
- FileServer adalah handler, jadi bisa kita tambahkan ke dalam http.Server atau http.ServeMux
- bisa membuat Handler yg digunakan sebagai static file server
- dengan FileServer, kita tidak perlu manual me-load file lagi

# 404 Not Found using FileServer
prefix problem

- Jika kita coba jalankan, saat kita membuka misal /static/index.js, maka akan dapat error 404 Not Found.
- Kenapa ini terjadi?
- Hal ini dikarenakan FileServer akan membaca URL, lalu mencari file berdasarkan URL nya, jadi jika kita membuat /static/index.js, maka FileServer akan mencari ke file /resources/static/index.js
- Hal ini menyebabkan 404 Not Found karena memang filenya tidak bisa ditemukan
- Oleh karena itu, kita bisa menggunakan function http.StripPrefix() untuk menghapus prefix di URL.

- Like Embed

# 404 Not Found using golang embed
prefix problem

- Jika kita coba jalankan, saat kita membuka misal /static/index.js, maka akan dapat error 404 Not Found.
- Kenapa ini terjadi?
- Hal ini dikarenakan di Go-lang embed, nama folder ikut menjadi nama resource nya, misal resources/index.js, jadi untuk mengaksesnya kita perlu mengunakan URL /static/resources/index.js
- Jika kita ingin langsung mengakses file index.js tanpa menggunakan resources, kita bisa menggunakan function fs.Sub() untuk mendapatkan sub directory.

## ServeFile
- static file sesuai yg di inginkan
- http.ServeFile()
- dengan function ini, kita bisa menentukan file mana yg ingin kita tulis ke http response

## Template
# Web Dinamis
- Sampai saat ini kita hanya membahas tentang membuat response menggunakan String dan juga static file.
- Pada kenyataannya, saat kita membuat web, kita pasti akan membuat halaman yg dinamis, bisa berubah - ubah sesuai dengan data yg diakses oleh user.
- Di Go-lang terdapat fitur "HTML Template", yaitu fitur template yg bisa kita gunakan untuk membuat HTML yg dinamis.

# HTML Template / Template
- Fitur HTML template terdapat di package html/template
- Sebelum menggunakan HTML template, kita perlu terlebih dahulu membuat membuat template nya.
- Template bisa berubah file atau string.
- Bagian dinamis pada HTML Template, adalah bagian yg menggunakan tanda {{ }}

# Membuat Template
- Saat membuat templat dengan string, kita perlu memberi tahu nama template nya.
- Dan untuk membuat text templat, cukup buat text html, dan untuk konten yg dinamis, kita bisa gunakan tanda {{.}}, contoh:
- <html><body>{{.}}</body></html>

# Template File
- Selain membi=uat template dari string, kita juga bisa membuat template langsung dari file.
- Hal ini mempermudah kita, karena bisa langsung membuat file html.
- Saat membuat template menggunakan file, secara otomatis nama file akan menjadi nama template nya, misal jika kita punya file simple.html, maka nama template nya adalah simple.html.

t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

# Template Directory
t := template.Must(template.ParseGlob("./templates/*.gohtml"))

# Template Golang Embed
t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

## Template Data
- Saat kita membuat template, kadang kita ingin menambahkan banyak data dinamis.
- Hal ini bisa kita lakukan dengan cara menggunakan data struct atau map.
- Tapi perlu dilakukan perubahan di dalam text template nya, kita perlu memberi tahu Field atau Key mana yg akan kita gunakan untuk mengisi data dinamis di template.
- Kita bisa menyebutkan dengan cara seperti ini {{.NameField}}

## Template Action
- Go-lang template mendukung perintah action, seperti percabangan, perulangan, dll.

# If Else
{{if.Value1}}T1{{else if.Value2}}T2{{else}}T3{{end}}

# Operator Perbandingan
- eq :  arg1 == arg2
- ne :  arg1 != arg2
- lt :  arg1 <  arg2
- le :  arg1 <= arg2
- gt :  arg1 >  arg2
- ge :  arg1 >= arg2

# ex:
if [operator] [variable] [value/variable]
{{if ge .FinalValue 80}}

# Kenapa Operatornya di depan?
- Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
- Jadi saat kita menggunakan {{eq First Second}}, sebenarnya dia akan memanggil function eq dengan parameter First dan Second : eq(First, Second)

# Range (iteration in golang template)
- Range digunakan untuk melakukan iterasi data template.
- Tidak ada perulangan biasa seperti menggunakan for di Go-Lang template.
- Yang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map atau channel.
- {{range $indexx, $el := .Value}}T1{{end}}, jika value memiliki data, maka T1 akan dieksekusi sebanyak el value, dan kita bisa menggunakan $index untuk mengakses index dan $el untuk mengakses el.

# With
- Kadang kita sering membuat nested struct.
- Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue.
- Di template terdapat action with, yg bisa digunakan mengubah scope dot menjadi object yg kita mau.
- {{with .Value}}T1{{end}}, jika value memiliki data, di T1 semua dot akan merefer ke value.
- {{with .Value}}T1{{else}}T2{{end}}, sama seperti sebelumnya, namun jika value kosong, maka T2 yg akan dieksekusi.

## Template Layout
- Saat kita membuat halaman website, kadang ada beberapa bagian yg selalu sama, misal header dan footer.
- Best practice nya jika terdapat bagian yg selalu sama, disarankan untuk disimpan pada template yg terpisah, agar bisa digunakan di template lain.
- Go-lang template mendukung import dari template lain.

# Import Template
Untuk melakukan import, kita bisa menggunakan perintah berikut:
- {{template"nama"}}, artinya akan meng-import template "nama" dengan memberikan data apapun.
- {{template"nama".Value}}, artinya akan meng-import template "nama" dengan memberikan data value.

## Template Function
- Selain mengakses field, dalam template, func juga bisa diakses.
- Cara mengakses func sama seperti mengakses field, namun jika func tersebut memiliki parameter, kita bisa gunakan tambahkan parameter ketika memanggil function di template nya.
- {{.FunctionName}}, memanggil field FunctionName atau function FunctionName()
- {{.FunctionName"eko", "kurniawan"}}, memanggil function FunctionName ("eko", "kurniawan")

# Global Function
- Go-lang template memiliki beberapa global func
- Global function adalah function yg bisa digunakan secara langsung, tanpa menggunakan template data.
- Berikut adalah beberapa global function di Go-Lang template.

github.com/golang/go/blob/master/src/text/template/funcs.go

# Menambah Global Func
- Kita juga bisa menambahkan global func
- Untuk menambah global func, kita bisa menggunakan method Funcs pada template
- Perlu diingat, bahwa menambahkan global function harus dilakukan sebelum melakukan parsing template.

# Function Pipelines
- Go-lang template mendukung function pipelines, artinya hasil dari function bisa dikirim ke function berikutnya.
- Untuk menggunakan function ppipelines, kita bisa menggunakan tanda | , misal:
- {{ sayHello .Name | upper }}, artinya akan memanggil global function sayHello(Name) hasil dari sayHello(Name) akan dikirim ke function upper(hasil).
- Kita bisa menambahkan function pipelines lebih dari satu.

## Template Caching
- Kode-kode diatas yg sudah kita praktekan sebenarnya tidak efisien
- Hal ini dikarenakan, setiap Handler dipanggil, kita selalu melakukan parsing ulang template nya.
- Idealnya template hanya melakukan parsing satu kali diawal ketika aplikasinya berjalan.

- Selanjutnya data template akan di caching (disimpan di memory), sehingga kita tikda perlu melakukan parsing lagi.
- Hal ini akan membuat web kita semakin cepat.


## XSS Cross Site Scripting
- XSS adalah salah satu security issue yg biasa terjadi ketika membuat web
- XSS adalah celah keamanan, dimana orang bisa secara sengaja memasukkan parameter yg mengandung JavaScript agar dirender oleh halaman website kita.
- Biasanya tujuan dari XSS adalah mencuri cookie browser pengguna yg sedang mengakses website kita.
- XSS bisa menyebabkan account pengguna kita diambil alih jika tidak ditangani dengan baik.

# Auto Escape
- Berbeda dengan bahasa pemrograman lain seperti PHP, pada Go-Lang template, masalah XSS sudah diatasi secara otomatis.
- Go-Lang template memiliki fitur Auto Escape, dimana dia bisa mendeteksi data yg perlu ditampilkan di template, jika mengandung tag-tag html atau script, secara otomatis akan di escape.

- Semua function escape bisa dilihat disini:
    - github.com/golang/go/blob/master/src/html/template/escape.go
    - golang.org/pkg/html/template/#hdr-Contexts

- text/template: not support auto escape
- html/template: automatic support auto escape

ex: 
- Input not read as html tag

# Mematikan Auto Escape
- Jika kita mau, auto escape juga bisa kita matikan.
- Namun, kita perlu memberi tahu template secara eksplisit ketika kita menambahkan template data.
- Kita bisa menggunakan data.
- template.HTML, jika ini adalah data html.
- template.CSS, jika ini adalah data css.
- template.JS, jika ini adalah data JavaScript.

# Masalah XSS (Cross Site Scripting)
- Saat kita mematikan fitur auto escape, bisa dipastikan masalah XSS akan mengintai kita.
- Jadi pastikan kita benar-benar percaya terhadap sumber / input data yg kita matikan auto escape nya.

## Redirect
- Saat membuat website, kadang kita butuh melakukan redirect.
- Misal setelah selesai login, kita lakukan redirect ke halaman dashboard.
- Redirect sudah standard di HTTP
developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
- Kita hanya perlu membuat response code 3xx dan menambahkan header Location.
- Namun untungnya di Go-Lang, ada function yg bisa kita gunakan untuk mempermudah ini.

http.Redirect()

## Upload File
- Saat membuat web, selain menerima input data berupa form dan query param, kadang kita juga menerima input data berupa file dari client.
- Go-Lang Web sudah memiliki fitur untuk management upload file.
- Hal ini memudahkan kita ketika butuh membuat web yg menerima input file upload.

# MultiPart
- Saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan Request.FormFile(name), di dalam nya secara otomatis melakukan parsing terlebih dahulu.
- Hasilnya merupakan data-data yg terdapat pada package multipart, seperti multipart.File sebagai representasi file nya, dan multipart.FileHeader sebagai informasi file nya.

Manual parsing:
request.ParseMultipartForm(100 << 20) // max 100 mb
Automatic parsing:
Behind it will parse, def size max 32mb
request.FormFile()

## Download File
- Selain upload file, kadang kita ingin membuat halaman website yg digunakan untuk download sesuatu.
- Sebenarnya di Go-Lang sudah disediakan menggunakan FileServer dan ServeFile.
- Dan jika kita ingin memaksa file di download(tanpa di render oleh browser, kita bisa menggunakan header Content-Disposition)
- developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition

- As a response header for the main body.
default inline = display / render picture / HTML
attachment = download default url name
attachment; filename="filename.jpg" = download custom name

Response Header:
Content-Disposition: inline
Content-Disposition: attachment
Content-Disposition: attachment; filename="filename.jpg"

## Middleware
- Middleware / Filter / Interceptor
- Dalam pembuatan web, ada konsep yg bernama middleware atau filter atau interceptor.
- Middleware adalah sebuat fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler di eksekusi

# Diagram Middleware
Middleware -> Handler

# Middleware di Go-Lang web
- di Go-Lang web tidak ada middleware.
- Namun karena struktur handler yg baik menggunakan interface, kita bisa membuat middleware sendiri menggunakan handler.

# Error Handler
- Kadang middleware juga biasa digunakan untuk melakukan error handler.
- Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan mengubah panic tersebut menjadi error response.
- Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan.
- catch panic

## Routing Library
- Walaupun Go-Lang sudah menyediakan ServeMux sebagai handler yg bisa menghandle beberapa endpoint atau istilahnya adalah routing.
- Tapi kebanyakan programmer Go-Lang biasanya akan menggunakan library untuk melakukan routing.
- Hal ini dikarenakan ServeMux tidak memiliki advanced fitur seperti path variable, auto binding parameter dan middleware.
- Banyak alternatif lain yg bisa kita gunakan untuk library routing selain ServeMux.

# Recommend Library for routing:
- github.com/julienschmidt/httprouter
- github.com/gorilla/mux
- github.com/julienschmidt/go-http-routing-benchmark

- advantage advanced feature:
    path variable, auto binding parameter dan middleware.

### 8. end
<!-- SOL 898 -->
</details>

<details>
<summary>### 9. HttpRouter</summary>
<br>
<!-- EOL 1350 -->

### 9. HttpRouter
## Pengenalan HttpRouter
- HttpRouter merupakan salah satu OpenSource Library yg populer untuk Http Handler di Go-Lang.
- HttpRouter terkenal dengan kecepatannya dan juga sangat minimalis.
- Hal ini dikarenakan HttpRouter hanya memiliki fitur untuk routing saja, tidak memiliki fitur apapun selain itu.

- github.com/julienschmidt/httprouter

# Project Set Up
# Menambah HttpRouter ke Project
- go get github.com/julienschmidt/httprouter

UT Library, help with assert
- go get github.com/stretchr/testify

## Router
- Library HttpRouter adalah sturct ROuter
- Router ini merupakan implementasi dari http.Handler, sehingga kita bisa dengan mudah menambahkan ke dalam http.Server.
- Untuk membuat Router, kita bisa menggunakan function httprouter.New(), yg akan mengembalikan Router pointer.

- Kenapa bentuk pointer, agar menggunakan router yg sama, tidak di copy.

# HTTP Method
- Router mirip dengan ServeMux, dimana kita bisa menambahkan route ke dalam Router.
- Kelebihan dibandingkan dengan ServeMux adalah, pada Router, kita bisa menentukan HTTP Method yg ingin kita gunakan, misal GET, POST, PUT, dll.
- Cara menambahkan route ke dalam Router adalah menggunakan function yg sama dengan HTTP Method nya, misal router.GET(), .POST(), dll.

# httprouter.Handle
- Saat kita menggunakan ServeMux, ketika menambah route, kita bisa menambahkan http.Handler.
- Berbeda dengan Router, pada Router kita tidak menggunakan http.Handler lagi, melainkan menggunakan type httprouter.Handle.
- Perbedaan dengan http.Handler adalah, pada httprouter.Handle, terdapat parameter ke tiga yaitu Params, yg akan kita bahas nanti di chpater tersendiri.

- beda kontrak

## Params
- httprouter.Handle memiliki parameter yg ketiga, yaiituParams.
Untuk apa kegunaan Params?
- Params merupakan tempat untuk menyimpan parameter yg dikirim dari client.
- Namun Params ini bukan query parameter, melainkan paremeter di URL.
- Kadang kita butuh membuat URL yg tidak gix, alias bisa berubah-ubah, misal:
/products/{product_id},
/products/1,
/products/2,
dll.
- ServeMux tidak mendukung hal tersebut, namun Router mendukung hal tersebut.
- Parameter yg dinasmis yg terdapat di URL, secara otomatis dikumpulkan di Params.
- Namun, agar Router tahu, kita hatus memberi tahu ketika menambahkan Route, dibagian mana kita akan buat URL path nya menjadi dinamis.

## Router Patterns
- Sekarang kita sudah tahu bahwa dengan menggunakan Router, kita bisa menambah params di URL.
- Sekarang pertanyaannya, bagaimana pattern (pola) pembuatan parameter nya?

# Named Parameter
- Named parameter adalah pola pembuatan parameter dengan menggunakan nama.
- Setiap nama parameter harus diawali dengan: (titik dua), lalu diikuti dengan nama parameter.
- Contoh, jika kita memiliki pattern seperti ini:

Pattern             /user/:user
/user/eko           match
/user/you           match
/user/eko/profile   no match
/user/              no match

# Catch All Parameter
- Selain named parameter, ada juga yg benama catch all parameter, yaitu menangkap semua parameter.
- Catch all parameter harus diawali dengan * (bintang), lalu diikuti dengan nama parameter.
- Catch all parameter harus berada di posisi akhir URL.

Pattern                 /src/*filepath
/src/                   no match
/src/somefile           match
/src/subdir/somefile    match

## Serve File
- Pada materi Go-Lang Web, kita sudah pernah membahas tentang Serve File.
- Pada Router pun, mendukung serve static file menggunakan function ServeFiles(Path, FileSystem).
- Dimana pada Path, kita harus menggunakan Catch All Parameter.
- Sedangkan pada FileSystem kita bisa melakukan manual load dari folder atau menggunakan golang embed, seperti yg pernah kita bahas di materi Go-Lang Web.

## Panic Handler / Error Handler
- Apa yg terjadi jika terjadi panic pada logic Handler yg kita buat?
- Secara otomatis akan terjadi error, dan web akan berhenti mengembalikan response.
- Kadang saat terjadi panic, kita ingin melakukan sesuatu, misal memberitahu jika terjadi kesalahan di web, atau bahkan mengirim informasi log kesalahan yg terjadi.
- Sebelumnya, seperti yg sudah kita bahas di materi Go-Lang Web, jika kita ingin menagani panic, kita harus membuat Middleware khusus secara manual.
- Namun di Router, sudah disediakan untuk menangani panic, caranya dengan menggunakan attribute PaniceHandler: func(http.ResponseWriter, *http.Request, interface{})

## Not Found Handler
- Selain panic handler, Router juga memiliki not found handler.
- Not foundhandler adalah handler yg dieksekusi ketika client mencoba melakukan request URL yg memang tidak terdapat di Router.
- Secara default, jika tidak ada route tidak ditemukan, Router akan melanjutkan request ke http.NotFound, namun kita bisa mengubah nya.
- Caranya dengan mengubah router.NotFound = http.Handler.

## Method Not Allowed Handler
- Saat menggunakan ServeMux, kita tidak bisa menentukan HTTP Method apa yg digunakan untuk Handler.
- Namun pada Router, kita bisa menentukan HTTP Method yg ingin kita gunakan, lantas apa yg terjadi jika client tidak mengirim HTTP Method sesuai dengan yg kita tentukan?
- Maka akan terjadi error Method Not Allowed.
- Secara default, jika terjadi error seperti ini, maka Router akan memanggil function http.Error.
- Jika kita ingin mengubahnya, kita bisa gunakan router.MethodNotAllowed = http.Handler.

## Middleware
- HttpRouter hanyalah library untuk http router saja, tidak ada fitur lain selain router.
- Dan karena Router merupakan implementasi dari http.Handler, jadi untuk middleware, kita bisa membuat sendiri, seperti yg sudaj kita bahas pada course Go-Lang Web.

### 9. end
<!-- SOL 1350 -->
</details>

<details>
<summary>### 10. JSON</summary>
<br>
<!-- EOL 1463 -->

### 10. JSON
## Pengenalan Package json
- JSON singkatan dari JavaScript Object Notation, merupakan struktur format data yg bentuknya seperti Object di JavaScript.
- JSON merupakan struktur format data yg paling banyak digunakan saat kita membuat RESTful API.
- Dan pada kelas ini kita akan menggunakan JSON juga.

- json.org/jsonien.html

# Package json
- Go-Lang sudah menyediakan package json, dimana kita bisa menggunakan package ini untuk melakukan konversi data ke JSON (encode) atau sebaliknya (decode).

- pkg.go.dev/encoding/json

## Encode JSON
- Go-Lang telah menyediakan function untuk melakukan konversi data ke JSON, yaitu menggunakan function json.Marshal(interface{}).
- Karena parameter nya adalah interface{}, maka kita bisa masukan tipe data apapun ke dalam function Marshal.

## JSON Object
- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya.
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namu tidak sesuai dengan kontrak JSON.
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array.

# Struct
- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct.
- Dimana tiap attribute di JSON Object merupakan attribute di Struct.

## Decode JSON
- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON.
- Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang(Decode), kita bisa menggunakan function json.Unmarshal(byte[], interface{}).
- Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa berupa pointer.

## JSON Array
- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array.
- Array di JSON mirip dengan Array di JavaScript, dia bisa berisikan tipe data primitif, atau tipe data kompleks(Object atau Array).
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice.
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice.

# Decode JSON Array
- Selain menggunakan Array pada attribute di Object.
- Kita juga bisa melakukan encode atau decode langsung JSON Array nya.
- Encode dan Decode JSON Array bisa menggunakan tipe data Slice.

## JSON Tag
- Secara default atrubut yg terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut yg sama (case sensitive).
- Kadang ada style yg berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase.
- Untungnya, package json mendukung Tag Reflection.
- Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yg kita inginkan ketika konversi dari atau ke JSON.

## Map
- Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic.
- Artinya atribut nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap.
- Pada kasus seperti itu, menggunakan Struct akan menyulitkan, karena pada Struct, kita harus menentukan semua atribut nya.
- Untuk kasus seperti ini, kita bisa menggunakan tipe data map[string]interface{}.
- Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map.
- Namun karena value berupa interface{}, makan kita harus lakukan konversi secara manual jika ingin mengambil value nya.
- Dan tipe data Map tidak mendukung JSON Tag lagi.

## Streaming Decoder
- Sebelumnya kita belajar package json dengan melakukan konversi data JSON yg sudah dalam bentuk variable dan data string atau []byte.
- Pada kenyataanya, kadang data JSON nya berasal dari input berupa io.Reader(File, Network, Request Body).
- Kita bisa saja membaca semua datanya terlebih dahulu, lalu simpan di variiable, baru lakukan konversi dari JSON, namun hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur untuk membaca data dari Stream.

# json.Decoder
- Untuk membuat json Decoder, kita bisa menggunakan function json.NewDecoder(reader).
- Selanjutnya untuk membaca isi input reader dan  konversikan secara langsung ke data di Go-Lang, cukup gunakan function Decode(interface{}).

## Streaming Encoder
- Selain decoder, package json juga mendukung membuat Encoder yg bisa digunakan untuk menulis langsung JSON nya ke io.Writer.
- Dengan begitu, kita tidak perlu menyimpan JSON datanya terlebih dahulu ke dalam variable string atau []byte, kita bisa langsung tulis ke io.Writer.

# json.Decoder
- Untuk membuat json Encoder, kita bisa menggunakan function json.NewEncoder(writer).
- Dan untuk menulis data sebagai JSON langsung ke writer, kita bisa gunakan function Encode(interface{}).

### 10. end
<!-- SOL 1463 -->
</details>

<details>
<summary>### 11. RESTful API</summary>
<br>
<!-- EOL 1490 -->

### 11. RESTful API

# Aplikasi CRUD
- Kita akan membuat CRUD sederhana.
- Tujuannya untuk belajar RESTful API, bukan untuk membuat aplikasi.
- Kita akan membuat aplikasi CRUD untuk data category.
- Dimana data category memiliki atribut id(number) dan name(string).
- Kita akan membuat API untuk semua operasi nya, Create Category, Get Category, List Category, Update Category dan Delete Category.
- Semua API akan kita tambahkan Authentication berupa API-Key.

## Setup Project

# go mod init

# Menambah Dependency
- Drive MySQL:  github.com/go-sql-driver/mysql
- HTTP Router:  github.com/julienschmidt/httprouter
- Validation:   github.com/go-playground/validator

## Membuat OpenAPI
- Kontrak API

- github.com/OAI/OpenAPI-Specification/blob/main/examples

# API Spec List Categories
# API Spec Create Category
# API Spec Get Category
# API Spec Update Category
# API Spec Delete Category

# API Spec Security
"components": {
        "securitySchemes": {
            "CategoryAuth": {
                "type": "apiKey",
                "in": "header",
                "name": "X-API-Key",
                "description": "Authentication for Category API"
            }
        },
}

## Membuat Database
# Setup Database
- create DB
- create Table
- check if success

## Domain
# Membuat Model DB / Domain / Entity
- create "model" folder
- create "domain" sub folder

# Best practices:
- pascal case for domain
- eventho in db snake case column

## Repository
## Repository Implementation
# Membuat Repository / Data Access Layer
- create "repository" folder
- kontrak / interface
- struct / implementation

- sql.Tx is a struct so it must be pointer.

## Service
## Service Implementation
# Membuat Service
- Bisnis Logic
- Kontrak
- 1 API 1 Function biasanya

- create "web" folder
- request file
- response file
- helper ToCategoryResponse
- helper CommitOrRollback

- do transaction by service level
- service -> repository

## Validation
# Membuat Validation
- write validation in struct level

## Controller
# Membuat Controller
- create "controller" folder
- create kontrak controller interface

## Controller Implementation
- create controller impl file
- create struct
- decode helper
- encode helper

## HTTP Router
- NewCategoryController in controller impl
- NewCategoryService in service impl
- NewCategoryRepository in repository impl
- main.go
- call function
- set router
- set database config

## HTTP Server
- 
server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()

## JSON Tag
Set JSON Attribute for struct from Pascal Case to match your apispec / API Contract for response.

ex: json:"name"

- create, update request struct
- web response struct

## Manual Test API
- create file for trigger endpoint test.http / .rest
- or use postman

- test all the endpoint

# Error:
# Recommended:
- Set folder name without space for location trace.

# Terminal:
- Case1:
    json: unsupported type: func() string
meaning: try to make function as json,
    should be string.

- Case2:
    [mysql] read: connection reset by peer driver: bad connection

solution:
- Case1:
    exception := because err.(validator.ValidationErrors) return a function.
    exception.Error()
    - ctrl + space

- Case2:
    check db config string connection:
    - sql.Open(drive, dataSourceName)
    - driver "mysql"
    - dataSourceName: username:password@tcp(host:port)/database_name

    - check if mysql running
    - check host
    - check port
    - try without protocol / http://

# Check Query
- Repository level

# Check Validation String
- function use

# Invalid Connection
- close rows

# Unexpected Response
- check all layer function use
- in service layer -> service.CategoryRepository.Save()

## Error Handler
- Internar Server Error
- Not Found
- Validation Error
- Client Error

## Authentication
- create "middleware" folder
- auth_middleware file
- ServeHTTP function base on handler interface


## Unit Test
- go get github.com/stretchr/testify
# Integration Test
- Setup Router
- Setup Test DB Config
- Create DB for Test
- Create Tables
- Create Test Scenario
- Run Test Scenario:
    go test -v ./server/repository -run=TestCommentInsert
- Set Delete tables data before run scenario

# Unit Test

- Case:
- validator dependency: check validator package must use /v[version number]
- care for json response property dif
- type assertion: need to explicit define type
- no need content type for get
- 301 redirect: could be wrong url -> last / could affect url.

- data: [map[id:1 name:Gadget] map[id:2 name:Computer]]
- Process: responseBody["data"].([]map[string]interface{})

- Response: panic: interface conversion: interface {} is []interface {}, not []map[string]interface {} [recovered]
- Solution:
    []map[string]interface{} -> []interface{}

- Reason:
    - FindAll() func return -> []web.CategoryResponse.
    type CategoryService interface{
	    FindAll(ctx context.Context) []web.CategoryResponse
    }
    - data type is interface{} not map.
	type WebResponse struct{
        Data	interface{}	`json:"data"`
    }


### 11. end
<!-- SOL 1490 -->
</details>


<details>
<summary>### 12. Dependency Injection</summary>

### 12. Dependency Injection

### 12. end
<!-- SOL 1490 -->
</details>


<details>
<summary>### 13. Logging</summary>

### 13. Logging

### 13. end
<!-- SOL 1490 -->
</details>


<details>
<summary>### 14. Generic</summary>

### 14. Generic

### 14. end
<!-- SOL 1490 -->
</details>


<details>
<summary>### 15. DB Migration</summary>

### 15. DB Migration

### 15. end
<!-- SOL 1490 -->
</details>


<details>
<summary>### 16. Validation</summary>

### 16. Validation

### 16. end
<!-- SOL 1490 -->
</details>


## Materi Selanjutnya:
- Go-Lang Deployment
- Go-Lang Docker
- Go-Lang Library
- Pracite Building App

### GRaphQL

# Run UT in folder
go test -v ./server/repository -run=TestCommentInsert


## Aditional
- golang formater
- Fprint vs Fprintf

Algoritma:
var
A : [1,2,3,4,5]
ascending
a - b
function(param) {
    for([1,2,3,4,5]) {
        
    }
}

<!--
error:
panic: template: pattern matches no files: `templates/*.gohtlm`
-->
cause: space can produce error
case: 
- //go:embed
- //go:embed [file name has space]
- //go:embed [if still error, re type location]
- sida space URL

case:
_, err :=
no new variables on left side of :=compiler
:= -> =

Download File:
- sometimes catch will effect change response.
- try other file to test.

# Pointer
koneksi database, itu adalah object yang kompleks, jadi harus menggunakan pointer

# go test cache
- go clean -testcache: expires all test results.
- use non-cacheable flags on your test run. The idiomatic way is to use -count=1

