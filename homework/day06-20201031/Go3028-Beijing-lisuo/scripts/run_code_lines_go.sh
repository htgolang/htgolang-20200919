#!/usr/bin/env bash

declare -a DIRS
BASE="../user_manager_proj"

DOF=`ls -1 $BASE`

for iterm in $DOF; do 
    if [[ -d $BASE/$iterm ]]; then
      dir=$BASE/$iterm
      if [[ $iterm == "cmd" ]]; then
          for i in `ls -1 $BASE/$iterm`; do
              dir=$BASE/$iterm/$i
              go run ../code_lines/code_lines.go -f $dir
          done
      fi
      if [[ $iterm != "cmd" ]]; then
        go run ../code_lines/code_lines.go -f $dir
      fi
    fi
done
