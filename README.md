# Autobase Twitter
Autobase bot on Twitter to automatically read DMs using certain keyword and send the message as tweets. 
This bot is currently being used on [@UGM_FESS](https://twitter.com/UGM_FESS) (even tho I've sold the twitter account to other person) which can easily process over 15K messages per month.  

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### How to Set up The App

1. Create and set up your own `.env` file according to `.env.example`
2. Run Project by building its binary. This project has two binaries:
   * `worker` : is the bot itself. Run by using `make run-worker`
   * `api` : the REST api of the app which planned to be the dashboard, but currently is not implemented. Run by using `make run-api`
3. Enjoy.

## Directory Structure

This repository is organized in the following directory structure.

```
autobase-twitter
|-- bin                                    # Contains binary of the built app
|-- cmd                                    # Contains executable codes and serves as entry point of the app
|   |-- worker                             
|   |   |-- main.go                        # entry point of the bot (worker) app
|   |-- <other entry point>                # other entry point
|-- config                                 # project level config. Ideal for development purposes because it will be bypassed by ENV vars
|-- constant                               # project level constant
|-- internal                               # Go files in this folder represent the Big-Pictures and Contracts of the system
|   |-- app                                # Contains dependency injection of the app and other app's related configs
|   |   |-- config                         # Configuration struct for the app
|   |-- delivery                           # Delivery layer of the app
|   |   |-- rest                           # REST Server as delivery of the app
|   |   |-- worker                         # Worker to as entry point of the bot itself
|   |-- entity                             # Enterprise Data structures
|   |   |-- message.go                     # Model for message
|   |   |-- <other_entities>.go            # Other data structures, preferrably 1 struct 1 file 
|   |-- gateway                            # Implementations of Gateway for External Service
|   |   |-- autobase_gateway.g             # Gateway implementation for Twitter API
|   |   |-- <other_gateways>.go            # Other Gateway implementations based on interfaces on folder internal.
|   |-- repo                               # Implementations of Repository-pattern to data-sources
|   |   |-- db                             # Database implementation for Archive of Autobase Twitter.
|   |   |  |-- firebase                    # Database implementation using Firebase.
|   |-- usecase                            # Usecases implementations for Application Business Logic
|   |   |-- autobaseuc                     # Business logic for autobase usecase
|   |   |-- archiveuc                      # Business logic for archive usecase  
|   |-- <other_usecases>.go                # Other Usecase implementations based on interfaces on folder internal.
|   |-- gateway.go                         # Interfaces / Contracts of all the gateways (External Services)
|   |-- repo.go                            # Interfaces / Contracts of all the repositories (Repository Pattern)
|   |-- usecase.go                         # Interfaces / Contracts of all the use-cases (Application Business Logic)

```

## Contributing 

TODOs (Pull Requests are welcomed):
- Rules feature (min followers, should be a follower)
- Blocked words
- More clear error messages
- Success tweet reply with the link
- UI for looking at older messages (annoying to have to look at firebase json everytime)
