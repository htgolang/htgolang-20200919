#!/bin/bash
#
# ./parse_yaml.sh -p "prefix" -f ./some_file
# ./parse_yaml.sh -f ./some_file
#

PY_PREFIX=
PY_YAML_FILE=

help(){
    echo "help."
}

while getopts h?p:f: OPTION
do
    case $OPTION in
    	h|\?)
            help
            exit 0
            ;;
        p)
            PY_PREFIX=$OPTARG
            ;;
        f)
            PY_YAML_FILE=$OPTARG
            ;;
    esac
done

do_parse(){
    if [ -z ${PY_YAML_FILE} ]; then
        help
        exit 1
    fi
    
    local dis=`cat ${PY_YAML_FILE} | egrep "^[[:space:]].*" | head -n1 | awk '{
        split($0, chars, "");
        count=0;
        for (i=1; i <= length($0); i++) {
            if (chars[i] == " ") count++; else break;
        } 
        print count;
    }'`
    
    if [ -z ${dis} ]; then
        dis=1
    fi
    
    local s='[[:space:]]*' w='[a-zA-Z0-9_]*' fs=$(echo @|tr @ '\034')
    
    sed -ne "s|^\($s\):|\1|" \
        -e "s|^\($s\)\($w\)$s:$s[\"']\(.*\)[\"']$s\$|\1$fs\2$fs\3|p" \
        -e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  ${PY_YAML_FILE} |
    awk -F$fs '{
        indent = length($1)/'${dis}';
        vname[indent] = $2;
        for (i in vname) {if (i > indent) {delete vname[i]}}
        if (length($3) > 0) {
            vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
            printf("%s%s%s=\"%s\"\n", "'${PY_PREFIX}'",vn, $2, $3);
        }
    }'
}

do_parse

