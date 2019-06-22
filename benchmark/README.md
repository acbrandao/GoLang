# GoLang vs. Python vs. Node Javascript Benchmark Test

<h1>
 <p align="center">
<img src="https://github.com/acbrandao/Reference/blob/master/img/019-go-lang.png" width="128" style="float: left;">   
VS.
<img src="https://github.com/acbrandao/Reference/blob/master/img/005-python.png" width="128" style="float: right;">
 VS
<img src="https://github.com/acbrandao/Reference/blob/master/img/029-javascript.png" width="128" style="float: left;">
</p>
</h1>

## The Benchmark

This is a very simple benchmark of file processing and string manipulaiton in three common server-side languages. 
The benchmark is very simple but may be increased if there's any interest.  First we read a 1MB text file, and 
captilize every work and display it to the console, and time the results of each acript.

## Requirements

  * Requires Go Lang installed on youir system. [How to Intall Go Lang](https://tecadmin.net/install-go-on-ubuntu/)
  * Requires Python installed on your system . [Instlal Python](https://www.digitalocean.com/community/tutorials/how-to-install-python-3-and-set-up-a-local-programming-environment-on-ubuntu-16-04)
  * Requires NodeJS installed on your system. [Install NodeJS](https://www.digitalocean.com/community/tutorials/how-to-install-node-js-on-ubuntu-16-04)

## Installing

  Simply git clone this repository
  
 ``` git clone git@github.com:acbrandao/GoLang.git ```
 
 then cd into the _benchmark_ directory
 ```cd GoLang/benchmark/``
 
 and ```chmod  +x``` the benchmark shell script.
 
 
 ## Running the benchmarks 
 Run the included shell script:
 
 ```
 !/bin/bash
####################################
#
# benchmarks script, runs three seperate programs in python, nodejs and golang
# which reads a large text file and capitlaizes it to check performance
####################################

echo "Running Python bench.py bench mark"
time python bench.py > /dev/null
echo "Running NodeJS bench.js bench mark"
time node bench.js > /dev/null
echo "Running Go bench.go bench mark"
go build -o benchgo bench.go
time ./benchgo > /dev/null
````

 the benchmark shell script which runs and times each individual language.
 
 ``` chmod +x benchmark ```
 then run the script
 
``` ./benchmark```

# Benchmark Script

To run this benchmark on a Linux system (yes, you probably can run it in Windows with the proper binaries, but I did not test it)

# Benchmark results.

Run this on your system, your actual results of course will vary based on your hardware.   Below are mine run on I7 wuth 16GB or Ram
For comparison on Ubuntu Linux XPS15i7 4.4.0-17763-Microsoft #379-Microsoft Wed Mar 06 19:16:00 PST 2019 x86_64 x86_64 x86_64 GNU/Linux4_

| Program (Source)       | Time (s)  |  Size (MB) |
| ---------------------- | ----------| ----------|
| bench.go               | none      |  1.0 MB  |
| bench.py               | one       |  1.0 MB    |
| bench.js               |           |  1.0 MB    |
