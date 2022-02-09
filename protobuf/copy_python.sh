#!/usr/bin/env bash
#
#  SPDX-License-Identifier: MIT
#
#    Copyright (c) 2021, SCANOSS
#
#    Permission is hereby granted, free of charge, to any person obtaining a copy
#    of this software and associated documentation files (the "Software"), to deal
#    in the Software without restriction, including without limitation the rights
#    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
#    copies of the Software, and to permit persons to whom the Software is
#    furnished to do so, subject to the following conditions:
#
#    The above copyright notice and this permission notice shall be included in
#    all copies or substantial portions of the Software.
#
#    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
#    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
#    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
#    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
#    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
#    THE SOFTWARE.

# This file copies the Python library code into the scanoss.py repo (on the local file system)

export b_dir=$(dirname "$0") # script location
if [ "$b_dir" = "" ]; then
  export b_dir=.
fi
#
# Print help info
#
help() {
  echo "Usage: copy_python [-h]
               [-s source-folder]
               [-t target-folder]"
  exit 2
}

#
# Confirm before proceeding
# Args:
#      1 - force: true/false
#      2 - question: text
#
confirm() {
  if [ "$1" != "true" ] ; then
    read -r -p "$2 [Y/n] " response
    response=$(echo "$response" | tr '[:upper:]' '[:lower:]')
    if [[ $response =~ ^(yes|y| ) ]] || [[ -z $response ]]; then
      echo "Proceeding..."
    else
      echo "Aborting."
      exit 1
    fi
  fi
}

#
# Parse command options and take action
#
force=false
while getopts ":s:t:,hf" o; do
  case "${o}" in
  t)
    t=${OPTARG}
    ;;
  s)
    s=${OPTARG}
    ;;
  h)
    help
    ;;
  f)
    force=true
    ;;
  *)
    echo "Unknown option: $o"
    help
    ;;
  esac
done
shift $((OPTIND - 1))

if [ -z "${s}" ]; then
  s=python
fi
if [ -z "${t}" ]; then
  t=../../scanoss.py/src
fi

scanoss_dir=scanoss

confirm $force "Copy Python library?"
cd "$s"
if [ ! -d "$scanoss_dir" ] ; then
  echo "Error: No Python code to copy from $s/$scanoss_dir"
  exit 1
fi
if ! tar cf - "$scanoss_dir" | tar xvf - -C "$t" ; then
  echo "Error: Failed to copy $s/$scanoss_dir to $t"
  exit 1
fi

exit 0
