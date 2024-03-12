#!/bin/bash

commit_message="$@"

git add .
git commit -m "$commit_message"
git pull origin master
git push origin master
