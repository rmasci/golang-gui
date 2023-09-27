v#!/bin/bash

function diskspace {
  clear
  df -k
}

function availablemem {
  clear
  sysctl -n hw.memsize | awk '{print $1/1024/1024/1024 " GB " }'
}

function loggedinusers {
  clear
  who
}


function menu {
  clear
  echo
  echo -e "\t\t\t System Menu:\n"
  echo -e "\t1. Display Disk Space."
  echo -e "\t2. Dispay Available Memory."
  echo -e "\t3. See what user is logged in."
  echo -e "\t0. Exit menu."
  echo -en "\t\t Enter option: "

  read -n 1 option
}

while [ 1 ]
do
  menu
  case $option in
    0)
      break ;;
    1)
      diskspace ;;
    2) 
      availablemem ;;
    3)
      loggedinusers ;;
    *)
      clear
      echo "That isn't a valid option."
  esac
  echo -en "\n\n\t\t Hit any key to continue..."
  read -n 1 line
done
clear
