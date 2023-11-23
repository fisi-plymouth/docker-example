# Docker Exmaple

This is an example of how to create a simple development and production environment using docker. This was created on the back on the last FISI code meetup in Plymouth UK.


## Prerequisites

You need to either install [Docker Desktop](https://www.docker.com/products/docker-desktop/) or if you prefer just to use the CLI, you can install [docker engine](https://docs.docker.com/engine/install/)


## Development

To boot up the development environment run `docker compose up --build` the `--build` will build the images before starting the container, making sure the latest changes gets build.

This will run the containers in your terminal where you can see the output from each container. If you would like to run the containers in the background, you can pass the option `-d` to the command.

To stop the container(s) you run `docker compose down`.

You can read more about the some of the options you can pass to `compose up` [here](https://docs.docker.com/engine/reference/commandline/compose_up/).

The `docker-compose.override.yml` is the file we used in development, as it will merge and override values on the `docker-compose.yml` file. This happens automatically.


## Production

To build the images for production, run `docker compose -f docker-compose.yml build`. By passing in the `-f docker-compose.yml` we are telling it to not use the override file, but explicitely use that compose file. 

Once build you can find the images in your docker desktop or via CLI `docker images`.
