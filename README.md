# Setting up Cypress with Docker
A tutorial on setting up cypress to an application with docker

## What is E2E testing and Why?
End-to-end testing (E2E testing for short) is a technique that aims at testing the flow of an application from "end to end". Or how I understood it - writing tests that try to mimic the way the users would interact with your application. 

In contrast to the unit and integration testing you are not breaking down the application into separate smaller parts to test it, E2E testing aims to test the application in its entirety. Because even if the unit and integration tests seem to work well by themselves, they might not work well when put together as a whole, and according to Edward Robinson - "*The faster code fails, the less bugs we find in QA, the faster our QA cycles are*".

Now that we have a better understanding of what E2E testing is, let's talk about some E2E frameworks next! 

## Cypress - a replacement for Selenium?
When looking around for E2E testing tools, a common name that pops up is "Selenium". It is one of the more popular and oldest browser automation tools and is based on Java. However, Selenium is notorious for being difficult to write and a pain to install/set up. 

If you think that I am talking about trees when I talk about 'cypress' then great! It means that you have not heard about Cypress.io before and this is your first introduction to it. Cypress.io is an open-source framework built for end-end testing web apps. In comparison with Selenium, Cypress is relatively easy to install and set up. For this reason, it has started to gain traction with front-end developers that want to implement testing in their web apps!

Unfortunately, Cypress tests can only be written in JavaScript which means you would need to learn JavaScript to use it. On the other hand, Selenium provides support for several languages, including Java, C#, Python, Ruby, R, Dart, Objective-C, and, yes, even JavaScript. Furthermore, Selenium has support for several browsers such as Firefox, Safari, Edge, and IE - whilst the Cypress test runner only works on Chrome... This is only scratching the surface of how Selenium and Cypress differ from each other and if you want to learn more about it I think this [article](https://applitools.com/blog/cypress-vs-selenium-webdriver-better-or-just-different/) gives a good starting point. 

## Not everyone has a Node.js stack or wants to...
For developers that are experienced in working with Node.js, installing Cypress can be seen as something easy. However, for developers that work with Python or Go, having to use npm can be problematic. A solution would be to have a Docker image with Cypress pre-installed. 

In this tutorial, we will use `docker-compose` to run the Cypress tests in one container and the demo/tutorial app in another container. This way we can decouple our test framework from the app. 

## Simple Web App
The tutorial app is a very simple - the user clicks on the "Click Me!" button which will re-direct the user to another page. The project layout is the following:

```
cypress-docker-example/
|-- Dockerfile                      <--- docker file that will tell how to run the web app's container
|-- main.go                         <--- src code for our little simple web app
|-- e2e                             <--- folder that contains all the e2e tests
    |-- docker-compose.yml          <--- docker-compose file that will link the two containers together
    |-- cypress.json                <--- cypress configurations
    |-- cypress
        |-- videos
        |-- integration
            |-- spec.js             <--- the defined e2e tests 
```

## Run the Web App locally
First and foremost, let's test that our web app works like it should and whether or not we can run our app in a Docker container.

1. Start with cloning the Github repository: 

`git clone https://github.com/emetha/cypress-docker-tutorial.git`

2. Go into your new project folder:

`cd cypress-docker-tutorial`

3. Now we want to build a Docker image and we will call it 'simple-app':

Note: you might need to run the `docker` and `docker-compose` commands with `sudo`. 

`docker build --tag simple-app`

4. To run our app in a docker container, we simply use the command: 

`docker run --interactive --tty --env PORT=1234 publish 1234:1234 simple-app`

- `tty`: allocate a pseudo-TTY
- `interactive`: keep STDIN open even if not attached
- `env PORT=1234`: set the port environment variable to 1234
- `publish 1234:1234`: publish the container's 1234 port to the host's 1234 port.

## Creating Our First Cypress E2E Test!
There are three files we need to implement to start writing our Cypress E2E test: cypress.json, docker-compose.yml and integration/spec.js.

### Writing our cypress.json file

The cypress.json file specifies the configurations of Cypress. Let create a simple cypress.json file:

```
$EDITOR /e2e/cypress.json
```
Copy the JSON snippet below:

```json
{
  "pluginsFile": false,
  "supportFile": false
}
```

These were set to false to tell Cypress not to generate unneccessary helper files. 

### Writing our docker-compose.yml file
The docker-compose.yml file tells how the two containers should communicate with each other. To create the docker-compose.yml we use the command:

```
$EDITOR /e2e/docker-compose.yml
```

Then we copy this YAML snippet below to our docker-compose.yml:

```YAML
version: '3.2'
services:
  simple-app:
    build: ../
    environment:
      - PORT=1234
  cypress:
    image: "cypress/included:4.4.0"
    depends_on:
      - simple-app
    environment:
      - CYPRESS_baseUrl=http://simple-app:1234
    working_dir: /e2e
    volumes:
      - ./:/e2e
```
Here, we have more things to unpack and we will go through the lines that are of interest:

```YAML
image: "cypress/included:4.4.0"
```
Cypress have provided their own [Docker images](https://docs.cypress.io/examples/examples/docker.html#Images) that we can build from. These images have Cypress already installed and there are various types:
- `cypress/base:<node version>`: includes the dependencies that are required from the operating system to run Cypress.
- `cypress/browsers:<tag>`: an extension of the base images where browsers are pre-installed.
- `cypress/included:<cypress version>`: an extension of the base images where Cypress versions are pre-installed. 
In our case, we have chosen to build from the version 4.4.0 of `cypress/included`, this is to make sure that Cypress executes tests as soon as its container starts up. 

```YAML
depends_on: 
  - simple-app
```
This is to make sure that our simple app is set up and running before Cypress starts executing its tests.

```YAML
environment: 
  - CYPRESS_baseUrl=http://simple-app:1234
```
With this we can allow the Cypress container to send network requests to our app container through the given URL. Note that we can use the container name (which we set for our app previously) as our hostname because both containers reside in the same Docker Compose configuration.  


### Writing our spec.js file
Now comes the part we have all been waiting for... writing E2E tests! Since all our app does is re-direct the user to another page once they click on a button, the tests are relatively simple; check that we are on the correct page after the button has been clicked. 



