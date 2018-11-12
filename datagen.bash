#!/usr/bin/env bash
#--- 
#--- datagen.bash [record length] [number of records] [records per second]
#--- defaults are datagen.bash 1024 1000 100
#--- works for Linux and Mac OS X (Darwin)
#--- in mac os x make sure you do brew install coreutils pv
#--- 
#--- gustavo.maurizio@gmail.com
#--- 
#set -ex
# get values from command line arguments
# Argument 1 is record length (last byte will be a newline character)
# Argument 2 is the number of records that will be generated
# Argument 3 is the rate at which records will be generated (approximate) per second
reclen=${1:-1024}
numrecs=${2:-1000}
rate=${3:-100}
# subtract one to allow for the newline at the end
recbody=$((reclen - 1))
# Mac OS X date does not handle microseconds and we need to use GNU date
# that can be installed with brew install coreutils
# pv is also not standard, so brew install pv
MACOSX=false
LINUX=false
[[ "$(uname -s)" == "Darwin" ]] && MACOSX=true 
[[ "$(uname -s)" == "Linux"  ]] && LINUX=true
$MACOSX && datecmd="gdate"
$LINUX  && datecmd="date"
# iterate and create the records to be sent to stdout into pv for rate management
>&2 echo "$0 will generate $numrecs records of $reclen size at $rate messages per second"
for i in $(seq 1 $numrecs)
do
  # get the time as part of the body of the message
  printf "%0${recbody}f\n" $(${datecmd} +%s.%N)
done | pv --line-mode --rate-limit "$rate"
>&2 echo "$0 ended"
