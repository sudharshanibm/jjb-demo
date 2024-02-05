#!/bin/bash
username=$(awk -F "=" '/user/ {print $2}' jenkins_job.ini)
password=$(awk -F "=" '/password/ {print $2}' jenkins_job.ini)
echo "${username}:${password}"
C:/Users/SudharshanMuralidhar/AppData/Local/Programs/Python/Python39/python.exe -m venv .
source ./Scripts/activate      
pip install jenkins-job-builder
jenkins-jobs --conf ./jenkins_jobs.ini update ./matrix-pipeline.yaml --delete-old
curl -X POST -u ${username}:${password} \
    http://localhost:8080/job/multi-config-job-1/buildWithParameters?PARAMETER=Successful


