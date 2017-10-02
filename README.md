tbcnv
======

Tabular converter

Basic Usage
------------

```sh
tbcnv --help
```

```sh
Usage of tbcnv:
  -d string
        デリミタを指定します。 (default "\t")
  -f string
        出力する列名を指定します。指定されない場合、すべての列を出力します。
  -t string
        出力形式を指定します。指定可能な形式: tsv,csv,json,yaml (default "tsv")
```

```sh
# content of sample table
$ cat in.csv
a,b,c
1,2,3
3,4,5

# convert to json
$ tbcnv -d , -t json in.csv 
[{"a":"1","b":"2","c":"3"},{"a":"3","b":"4","c":"5"}]

# convert to yaml
$ tbcnv -d , -t yaml in.csv
- a: "1"
  b: "2"
  c: "3"
- a: "3"
  b: "4"
  c: "5"
  
# filtering
$ tbcnv -d , -t csv -f a,b in.csv
a,b
1,2
3,4

# pipe
$ cat in.csv | tbcnv -d , -t csv -f a,b
a,b
1,2
3,4
```

Setup
------

### require

Go

### install

```sh
go get -u github.com/0x75960/tbcnv
```
