# GoLang vs. Python vs. Node Javascript Benchmark Test

## The Benchmark

This is a very simple benchmark of file processing and string manipulaiton in three common server-side languages. 
The benchmark is very simple but may be increased if there's any interest.  First we read a 1MB text file, and 
captilize every work and display it to the console.

## Installing

  Simply git clone this repository
  
 ``` git clone https://github.com/acbrandao/GoLang/new/master/benchmark ```
 
 
 ## Running the benchmarks 
 and run the benchmark shell script which runs and times each individual language.
 
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
