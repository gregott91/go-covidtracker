# go-covidtracker

## Overview

This repository will download US daily COVID data from [The COVID Tracking Project](https://covidtracking.com), deserialize it and convert it to a well-known output format. It accepts a single argument, which should contain the filename to save the downloaded data to.

## Building

From the top level of the directory, run the following commands:

```Go
go build
```

## Running

Once the EXE has been build, you must pass it a single command line argument, containing the file path where the data should be saved. For example:

From the top level of the directory, run the following commands:

```
.\covidtracker.exe "C:\data.json"
```