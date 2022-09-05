#!/bin/bash
cd /home
bash $PROGRAM_PATH/send-command 1 "$@" LLT
sleep 60
bash $PROGRAM_PATH/send-command 1 "$@" SED
exit