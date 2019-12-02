#!/bin/sh

setup_git() {
  git config --global user.email "admin@lees.work"
  git config --global user.name "lee"
}

commit_website_files() {
  git add .
  git commit --message "Travis build: $TRAVIS_BUILD_NUMBER"
}

upload_files() {
  git remote add origin-pages https://${GH_TOKEN}@github.com/ConserveLee/leetcode_solution.git > /dev/null 2>&1
  git push --quiet --set-upstream origin-pages master
}

setup_git
commit_website_files
upload_files