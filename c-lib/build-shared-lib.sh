#!/usr/bin/env bash

gcc -c -fpic greet.c

gcc -shared -o libgreet.so greet.o