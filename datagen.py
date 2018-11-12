#!/usr/bin/env python
from __future__ import print_function

import time
import os
import sys
import random
import argparse

def high_resolution_sleep(duration):
    end = time.time() + duration
    if duration > 0.02:
        time.sleep(duration)
    while time.time() - end < 0:
        time.sleep(0)

def progress(count, total, status=''):
    bar_len = 60
    filled_len = int(round(bar_len * count / float(total)))

    percents = round(100.0 * count / float(total), 1)
    bar = '=' * filled_len + '-' * (bar_len - filled_len)

    sys.stderr.write('[%s] %6.2f%s ...%s\r' % (bar, percents, '%', status))
    sys.stderr.flush() 


def main():
    ### get the name of this script removing path information if present
    myName              = os.path.basename(os.path.splitext(sys.argv[0])[0])
    ### set defaults and a more easy to understand command line argument style
    parser = argparse.ArgumentParser(description="datagen utility")
    
    parser.add_argument("-n", "--number",
                        default=10,
                        type=int,
                        help="number of records")
    parser.add_argument("-l", "--length",
                        default=100, 
                        type=int,
                        help="record length")
    parser.add_argument("-r", "--rate",
                        default=1,
                        type=float,
                        help="record rate in records per second")

    args = parser.parse_args()

    print("%s will generate %d records of %d bytes at %f records per second" % (myName, args.number, args.length, args.rate), file=sys.stderr)
    ### get the actual format string, that will print time in Unix format with nanoseconds 
    formatlen = args.length - 1
    formatstr = "%0" + str(formatlen) + "f"
    ### calculate the time between two successive records to be generated based on the record rate specified
    ### formula is simple, waittime = 1.0 / float(args.rate) 
    waittime   = 1.0 / args.rate
    ### get the start time to compute rate
    time_start = time.time()
    for i in range(args.number):
        time_now = time.time()
        print(formatstr % (time_now))
        if i % max(5, args.number / 50) == 0:
            progress(i, args.number,       status='%10d rate %7.2f records per second' % (i, float(i) / (time_now - time_start)) )
        high_resolution_sleep(waittime)
    progress(args.number, args.number, status='%10d records at an average rate of %7.2f records per second' % (args.number, float(args.number) / (time_now - time_start)) )
    print("\n%s ended" % (myName), file=sys.stderr)

if __name__ == "__main__":
    main()

