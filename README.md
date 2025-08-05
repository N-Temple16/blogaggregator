# Blog Aggregator (blogaggregator)

## Prerequisites
In order to run the `blogaggregator` program, Postgres and Go must be installed.

1. - Postgres installation macOS ```brew install postgresql@15```
   - Postgres installation Linux/WSL ```sudo apt update``` ```sudo apt install postgresql postgresql-contrib```
2. For Go installation, follow the official installer guide for your OS ```https://go.dev/doc/install```

## Installation
To install the `blogaggregator` CLI, you will need Go installed and your Go bin directory on your system `PATH`.
1. Install the CLI
- ```go install github.com/N-Temple16/blogaggregator@latest```
2. Add the Go bin directory to your PATH (if it is not already):
- ```export PATH="$PATH:$(go env GOPATH)/bin"```

This will install a gator executable to your Go bin directory (typically ~/go/bin), allowing you to run it from anywhere in your terminal.


## Configuration and Commands
Before using `blogaggregator`, you need to create a configuration file that stores your database connection string and tracks the currently logged-in user.

1. Create the config file

In your home directory, create a file named `.gatorconfig.json`:
- ```touch ~/.gatorconfig.json```

2. Add the following contents to the configuration file:
- ```
  {
    "db_url": "postgres://username:password@localhost:5432/dbname?sslmode=disable",
    "current_user_name": ""
  }
You will need to replace the `db_url` with your actual Postgres credentials. The `current_user_name` field will be automatically filled in after you log in.

3. Once your configuration file is set up, you can begin using the `blogaggregator` CLI.

Here are some of the available commands at this current time:

`blogaggregator addfeed` - Add a new RSS feed to the current user

`blogaggregator agg` - Start the aggregator to collect posts from followed feeds

`blogaggregator browse` - Browse from feeds the current user is following

`blogaggregator feeds` - List all available RSS feeds

`blogaggregator follow` - Follow a feed for the current user

`blogaggregator following` - List all feeds the current user is following

`blogaggregator login` - Log in as an existing user

`blogaggregator register` - Register a new user

`blogaggregator reset` - Reset the database (clears all data)

`blogaggregator unfollow` - Unfollow a feed the current user is following

`blogaggregator users` - List all registered users

