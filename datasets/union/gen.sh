#!/bin/sh
typeset -i num
typeset -i sz
sz=$1+2
echo $sz
num=0
while [[ $num -le $1 ]];do
	echo $num $(($num+1))
	num=$num+1
done
echo 1 $num
num=0
while [[ $num -le $1 ]];do
	echo $(($1-1)) $1
	num=$num+1
done
