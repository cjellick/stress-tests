# Stress tests

A set of stress tests for Rancher

The eventual goal is to have a CLI interface for running various stress tests. The CLI will record and report on things like stack launch time and number of containers that failed to deploy, Right now, we have one "stress test" and nono CLI.

### chatty-cow
This is a stack that you can deploy that create services that attempt to make an HTTP request to another random "chatty-cow" every second. If the request succeeds, the container reports healthy on its healthcheck. If the request fails, it reports unhealthy. In this way, you can see if networking begins to fail as you scale up the service

The template for this stack is in stack-templates/chatty-cow. 
