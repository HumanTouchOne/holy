#!/usr/bin/env bash
git add -A;
git commit -m "$USER $(date +%H:%M:%S)";
git push  ;
#git push upstream master;

# git tag -l pattern "v0.0.*" | xargs -n 1 git push origin --delete
# git tag -l pattern "v0.0.*" | xargs -n 1 git tag -d