#!/bin/bash

cd public
git init
git config user.name "Thinh Tran"
git config user.email "duythinht@gmail.com"

echo "5kbps.io" > CNAME

git add .
git commit -m "Deploy from Darkness"
git push --force --quiet "https://${GITHUB_TOKEN}@github.com/duythinht/duythinht.github.io.git" master:master
