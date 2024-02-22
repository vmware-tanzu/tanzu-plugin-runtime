#!/bin/bash

# Copyright 2023 VMware, Inc. All Rights Reserved.
# SPDX-License-Identifier: Apache-2.0

# Script to run lint command on every plugin found/installed.
# Collect terms in violation to be considered for inclusion into the global
# word list.

set -o nounset
set -o pipefail

PLUGINS_JSON=/tmp/plugins_list.json
LINT_OUT=/tmp/plugins_lint.txt
NEW_WORDS=/tmp/new_words.txt

echo > ${LINT_OUT}
echo > ${NEW_WORDS}

tanzu plugin search -o json > ${PLUGINS_JSON}

for i in `cat ${PLUGINS_JSON} | jq -r '.[] | select(.target == "mission-control") | .name'`; do 
   echo "mission-control $i"
   echo "** mission-control $i" >> ${LINT_OUT}
   echo tanzu mission-control $i lint
   tanzu plugin install $i --target mission-control
   tanzu mission-control $i lint >> ${LINT_OUT} 
done

for i in `cat ${PLUGINS_JSON} | jq -r '.[] | select(.target == "operations") | .name'`; do 
   echo "operations $i"
   echo "** operations $i" >> ${LINT_OUT}
   echo tanzu operations $i lint
   tanzu plugin install $i --target operations
   tanzu operations $i lint >> ${LINT_OUT} 
done

for i in `cat ${PLUGINS_JSON} | jq -r '.[] | select((.target == "kubernetes") or (.target == "global")) | .name'`; do
   echo "$i"
   echo "** $i" >> ${LINT_OUT}
   echo tanzu $i lint 
   tanzu plugin install $i
   tanzu $i lint >> ${LINT_OUT} 
done

cat ${LINT_OUT} | perl -pe 's/^.* unknown top-level term (\S+),\s.*$/NEWNOUN $1/' | grep NEWNOUN | sort | uniq | sort >> ${NEW_WORDS}
cat ${LINT_OUT} | perl -pe 's/^.* unknown subcommand term (\S+),\s.*$/NEWTERM $1/' | grep NEWTERM | sort | uniq >> ${NEW_WORDS}
cat ${LINT_OUT} | perl -pe 's/^.* unexpected flag (\S+),\s.*$/NEWFLAG $1/' | grep NEWFLAG | sort | uniq >> ${NEW_WORDS}
cat ${LINT_OUT} | perl -pe 's/^.* unexpected global flag (\S+),\s.*$/NEWGLOBALFLAG $1/' | grep NEWGLOBALFLAG | sort | uniq >> ${NEW_WORDS}
