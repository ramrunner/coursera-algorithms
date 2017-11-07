#!/bin/sh
SRV=http://localhost:8080
NAME=`echo $1 |sed 's/.[^.]*$//'`
echo $NAME
FIRST=1

while IFS= read -r line
do
	if [[ $FIRST -eq 1 ]]; then
		curl "$SRV/create/$NAME/$line/"	
		FIRST=0
	else
		if [[ -z $line ]]; then
			echo "blank line"
		else
			STR=`echo $line | sed 's/\ /\//'`
			curl "$SRV/open/$NAME/$STR/"
		fi
	fi
	

done<"$1"
