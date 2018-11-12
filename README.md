# datagen
## Utility to create data records at a specific rate for testing purposes
### TL;DR
A tool to generate unique records of a given length and at a given rate expressed in
records per second. Comes in two basic flavors (a golang version is on the works),
a bash script, and a python program.
#### The Linux way: datagen.bash
In order to use it, some prerequisite components are needed:
- Pipe Viewer tool (`pv`).
- For Mac OS X the gnu version of date (`gdate`) that can be installed using `brew install coreutils`.
Usage:
```
datagen.bash [record_length|1024] [number_of_records|1000] [rate|100]
```
For instance, `./datagen.bash 1512 3000 250` will generate 3,000 records of 1,512 bytes (each record terminated by newline '\n') at 250 records per second.

![output from command](/images/images/Screen_Shot_datagen.png?raw=true "Output from datagen.bash").

You can check the [output and progress bar here](https://gfycat.com/GiddyOfficialAvocet).

## Background: What problem is datagen trying to solve?
Test data generation is an important part of the tasks a software engineer needs to face.
While there are many open source and commercial tools available that solve many of the
problems, developer still struggle with some (relatively) simple tasks.
This tool helps in the particular space of testing APIs and micro-services, as well as
queueing systems, from an infrastructure and middleware perspective.
The **datagen** tool has been designed to solve the problem of quickly generating
testing data at a particular rate (expressed as records or messages per second),
and with a particular distribution of record or message lengths. The tool follows
the principles of Unix, in the sense that it writes its output to STDOUT in order
for it to be fed into other programs in a simple, automated, and scripted manner.
