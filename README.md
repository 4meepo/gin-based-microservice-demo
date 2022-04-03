# Gin-based Microservice Demo

In this demo, we will use Gin to build a simple cooking application.We will cover how to integrate,deploy,and test the app using the Gin framework.The application will do the following:
* Display the recipes that are submitted by the users, along with their ingredients and instructions.
* Allow anyone to post a new recipe.


## Prerequisites
* go 1.18+
* a MongoDB instance
  * You can use the MongoDB as a Service solution, known as MongoDB Atlas (https://www.mongodb.com/cloud/atlas), to run a free 500 MB database on the cloud. You can deploy a fully managed MongoDB server on AWS, Google Cloud Platform, or Microsoft Azure. 
  * You can run MongoDB locally with a containerization solution such as Docker. Multiple Docker images are available on DockerHub with a MongoDB server configured and ready to use out of the box.