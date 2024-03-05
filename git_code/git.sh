#!/bin/bash

git pull origin master

commit_message="$@"
#for i in $@
#do
#commit_message="$commit_message $i"
#done

git add .
git commit -m "$commit_message"
git push origin master
