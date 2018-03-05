# FileName:   Makefile
# Author:     Fasion Chan
# @contact:   fasionchan@gmail.com
# @version:   $Id$
#
# Description:
#
# Changelog:
#

.DEFAULT_GOAL := run

callee.o: callee.c
	gcc -fPIC -c callee.c

libcallee.so: callee.o
	gcc -shared -o $@ $^

caller: caller.go libcallee.so
	go build caller.go

clean:
	rm -f caller libcallee.so callee.o

run: caller
	./caller
