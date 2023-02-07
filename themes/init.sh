#!/bin/bash

cd $(dirname $0)

HUGO_THEME="git@github.com:devops-camp/maupassant-hugo.git"

git clone $HUGO_THEME maupassant-hugo
