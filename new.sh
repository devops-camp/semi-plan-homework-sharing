#!/bin/bash

[ -z "$1" ] && exit 

user=homework

hugo new posts/${user}/$1.md