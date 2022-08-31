#!/bin/bash
cd /home
bash $PROGRAM_PATH/start-jushin 1 "$@"
sleep 60
bash $PROGRAM_PATH/send-command 1 "$@" LLS
exit