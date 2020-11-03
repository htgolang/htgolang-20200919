#!/usr/bin/env bash

bash ./run_code_lines_go.sh \
    | grep lines \
    | awk '{num[$(NF-1)]++}END{all=0;for(i in num)all+=i; printf"All lines of fine end with .go/.cgo in user_manager_proj are: %s\n", all}'
