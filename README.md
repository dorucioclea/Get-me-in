[![CircleCI](https://circleci.com/gh/ProjectReferral/Get-me-in/tree/master.svg?style=svg&circle-token=632ab80f9b534a6dab955b1f27f267b00b700ac4)](https://circleci.com/gh/ProjectReferral/Get-me-in/tree/master)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Version](https://badge.fury.io/gh/tterb%2FHyde.svg)](https://badge.fury.io/gh/tterb%2FHyde)
[![GitHub last commit](https://img.shields.io/github/last-commit/google/skia.svg?style=flat)]()
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://opensource.org/)



# Referral Marketing System

## Overview

A web based system directed in helping a job seeker to land a job at their desired company. It involves a referrer and a referee. 

At the current scope, the referrer will advertise the job opportunity which will be available for referees to "apply". The referee can then choose to communicate with the referrer via instant messaging.

## Technical Overview

All the current microservices are built using GO with a mix of request-driven and event-driven architecture. For event-driven, we using RabbitMQ to broadcast messages.

#### Current services:
- Authentication Service(auth-service) - handles the lifecycle of JSON Web Tokens(JWT).
- Account Service(account-service) - handles all the CRUD operations to do with users.
- Marketing Service(marketing-service) - handles all the CRUD operations to do with job adverts.

#### New services under development/analysis:
- Customer Service(customer-service) - handles email confirmations, reset passwords and any other communications between the consumer and producer.
- Messaging Service(msg-service) - handles instant messaging between users.

#### Front-end:
Front end will be designed using React and Redux.

#### Deploy process:
We using CircleCI to manage our continuous integration environment. To manage our infrastructure, we are using Docker and AWS.






[Diagram to be added]

