#!/bin/bash

# application that will be run
APPNAME="ymlgenerator"

# Available commands
declare -A commands=(
	["start"]="start" ["stop"]="stop" ["restart"]="restart" ["status"]="status"
)

# text color
redColor='\033[0;31m'
greenColor='\033[0;32m'
yellowColor='\033[1;33m'
noColor='\033[0m'

# command param from args
comParam=$1


# function help
help_func(){
	echo -e "How to use this launcher script"
	echo -e "${yellowColor}List Command\t : ${greenColor}${commands[@]}"
	echo -e "${yellowColor}Usage \t\t : ${greenColor}$0 CommandName"
	echo -e "${yellowColor}Example \t : ${greenColor}$0 start"
	echo -e "${noColor}"
}

# run service function
# param 1 argument application name
run_service(){
	echo -e "${greenColor}starting $APPNAME ..."
	if [ -e $APPNAME.pid ]; then
		echo APPNAME already running, pid=`cat $APPNAME.pid`
	else
		nohup ./$APPNAME > /dev/null &
		pid=$!
		if [ -z $pid ]
			then echo "cannot get pid program `$1`"
		else
			echo $pid > $APPNAME.pid
		fi
		echo -e "$APPNAME Started${noColor}"
	fi
	echo -e "${noColor}"
	exit 1
}

# stop service function
# param 1 argument application name
stop_service(){
	echo -e "${redColor}kill $APPNAME"
	ps aux | grep -m 1 "${APPNAME}"
	if [ -e $APPNAME.pid ]; then
		kill `cat $APPNAME.pid`
        rm $APPNAME.pid
		echo -e "$APPNAME has been killed.${noColor}"
        else
		echo $APPNAME pid file not found
        fi
	echo -e "${noColor}"
}

# check args
if [ -z $comParam ]
then
	help_func
	exit 0
fi

case "$comParam" in
start)
   run_service
   ;;
stop)
   stop_service
   ;;
restart)
   echo -e "${redColor}Restarting $APPNAME$ ...${noColor}"
   stop_service
   run_service
   echo -e "${redColor}$APPNAME$ Restarted.${noColor}"
   ;;
status)
   echo -e "${yellowColor}"
   if [ -e $APPNAME.pid ]; then
      ps aux | grep -m 1 "${APPNAME}"
      echo $APPNAME is running, pid=`cat $APPNAME.pid`
      ps -p `cat $APPNAME.pid` -o pid,%cpu,%mem,cmd
   else
      echo $APPNAME is NOT running
   fi
   echo -e "${noColor}"
   exit 1
   ;;
*)
   help_func
esac

exit 0
