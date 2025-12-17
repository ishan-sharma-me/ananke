## Thesis:
The current situation is that multiple saas services exist for deployment, database, authentication, etc that have to purchased separately and then joined together which end up being quite expensive. All these essential services are already available production grade in the kubernetes ecosystem. A system that deploys these as kubernetes services in a easy way from a local environment will make this information available to the editor as yaml files that the AI can read and know the exact state of the system and act in that way. The kubernetes ecosystem is also production grade and open source which allows AI to have a better grasp of the services instead of closed source blackbox services with their own proprietary languages and APIs. Later kubernetes services can be added to the system to allow more and more services eventually paralleling the ecosystem of saas services or even surpass it.

## Goal:
The goal of ananke is allow user and AI agents to deploy production grade services directly from a local environment. Initially it will target a single local enviornment and deploy a single kubernetes cluster on a single server by finding the cheapest server that will be able to execute the deployment. Later this will allow resilient database deployments and other services. 

## Tech Stack:
- Uses go as the language as infra is generally written in go. Was trying to use effect ts but that made things more complicated to begin with, might switch later if required. Using effect is like using scala, more resilient programming. Might shift in case manually implementing lot of patterns of effect. Really resonate with the concept of functional programming.
- Single binary throughout, all the tools are deployed as a single binary.
- Initial target is to only run locally through a single laptop and a single server.

## Current Status:
[ ] - Create a basic machine hunter that works with intially only with Azure. Finds the approriate machine for the deployment and creates a kubernetes cluster on it.
[ ] - A single binary for the local system and a single binary for the server.
[ ] - Single local binary is the control plane stores state in a local database in the local system in the home directory.
[ ] - Single server binary is the agent that runs on the server and is responsible for the deployment of the services.
[ ] - Server binary for v1 just sets up a kubernetes cluster on the machine.