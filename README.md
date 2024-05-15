# RSSFlow: RSS Blog Aggregator

RSSFlow is a RSS feed aggregator built in Go! It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

## Personal Learning Goals

- Learn how to integrate a Go server with PostgreSQL
- Learn about the basics of database migrations
- Learn about long-running service workers

## Setup

### Configuration

**1. Define Port**

Create a gitignore'd .env file in the root of your project and add the following:

`PORT="8080"`

This file will automatically be loaded using `godotenv.Load()` in the main function.

**2. Connection String:**

After testing the connection, explained below, add the connection string to the .env file.

Don't forget to change username, password, port and database names.

`CONN=postgres://postgres:@localhost:5555/blogator?sslmode=disable`


### Setting up the environment

To install the correct packages and corresponding version, use the nix supplied nix flake and make sure that flakes are enabled.

In order to use it, run:

`nix develop`

## Usage

Run and test the server:

`go build -o out && ./out`

## Connect to Database

**Testing Connection:**

Test your connection string by running psql, for example:

```sh
psql "postgres://username:password@host:port/database"`
```

In my case, I am using:

```sh
psql "postgres://postgres:@localhost:5555/blogator"
```

**Run the Migration:**

To migrate, enter `sql` directory and run:

```sh
goose postgres postgres://postgres:@localhost:5555/blogator up
```
