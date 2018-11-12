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

    sys.stderr.write('>>> [{}] {:6.2f}{} ...{}\r'.format(bar, percents, '%', status))
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
    parser.add_argument("-j", "--jitterlen",
                        default=0,
                        type=int,
                        help="jitter in the record length")
    parser.add_argument("-f", "--jitterrate",
                        default=0,
                        type=float,
                        help="jitter in the record rate")

    args = parser.parse_args()

    print("{} will generate {:,} records of {:,} [+/- {:,}]  bytes at {:,.2f} [+/- {:,.2f}] records per second".format(myName, args.number, args.length, args.jitterlen, args.rate, args.jitterrate), file=sys.stderr)
    ### get the actual format string, that will print time in Unix format with nanoseconds 
    formatlen  = args.length - 1
    formatlenj = formatlen
    formatstr = "%0" + str(formatlenj) + "f"
    ### calculate the time between two successive records to be generated based on the record rate specified
    ### formula is simple, waittime = 1.0 / float(args.rate) 
    waittime   = 1.0 / args.rate
    ### count the bytes actually sent
    bytecount  = 0
    ### get the start time to compute rate
    time_start = time.time()
    for i in range(args.number):
        time_now = time.time()
        if args.jitterlen != 0:
            formatlenj = random.randint(formatlen - args.jitterlen, formatlen + args.jitterlen)
        formatstr = "%0" + str(formatlenj) + "f"
        bytecount += formatlenj + 1
        print(formatstr % (time_now))
        if i % max(5, args.number / 50) == 0:
            progress(i, args.number,       status='{:,d} records @ {:,.2f} rps. Total bytes: {:,d} avg record length: {:,.2f}'.format(i, float(i) / (time_now - time_start), bytecount, float(bytecount)/float(i+1)) )
        if args.jitterrate != 0:
            high_resolution_sleep(random.uniform(waittime * (1-args.jitterrate/args.rate), waittime * (1+args.jitterrate/args.rate) ))
        else:
            high_resolution_sleep(waittime)
    progress(args.number, args.number, status='{:,d} records @ {:,.2f} rps. Total bytes: {:,d} avg record length: {:,.2f}'.format(args.number, float(args.number) / (time_now - time_start), bytecount, float(bytecount)/float(args.number)) )
    print("\n%s ended" % (myName), file=sys.stderr)

if __name__ == "__main__":
    main()

