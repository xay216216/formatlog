





#!/usr/bin/env python
# -*- coding: utf-8 -*-
import sys
from pyspark import SparkContext
from operator import add
import  re

def main():
    sc = SparkContext(appName= "wordsCount")
    lines = sc.textFile('words.txt')
    counts = lines.flatMap(lambda  x: x.split(' '))\
                .map( lambda  x : (x, 1))\
                .reduceByKey(add)
    output = counts.collect()
    print output
    for (word, count) in output:
        print "%s: %i" %(word, count)

    sc.stop()

if __name__ =="__main__":
    main()
