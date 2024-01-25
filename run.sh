#!/bin/bash
C:/Users/SudharshanMuralidhar/AppData/Local/Programs/Python/Python39/python.exe -m venv .
source ./Scripts/activate      
pip install jenkins-job-builder
jenkins-jobs --conf ./jenkins_jobs.ini update  ./go-lms-pipeline.yaml --delete-old
curl -X POST -u sudharshan:11f3c221d1cb742f65075af435602b1979 \
  --data-urlencode json='{"parameter": [{"name":"PORT", "value":"4000"}]}' \
  http://localhost:8080/job/new-test-jjb-35/build


