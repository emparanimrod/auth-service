# Auth Service

This service is responsible for Authentication of all users. Among some other functionalities it handles, some of its
functions require the `User Service` in order to complete some functions e.g. `User Registration`

## Installation
The application uses a postgres as the backend database.


### Configuring

After cloning the source code, we need to configure some env variables before we build the application. The application
reads a `.env` file and populates the necessary `environment variables`. We need to create this `.env` file.

```bash
$ cp .env.example .env
```

The above command copies the example `dotenv` file which looks like this:

```dotenv
# Database configuration variables
AUTHAPP_DBUSER=user
AUTHAPP_DBPASSWORD=user
AUTHAPP_DBHOST=localhost
AUTHAPP_DBPORT=5432
AUTHAPP_DBNAME=auth_service

# Application specific configurations
# secret key used to sign jwt tokens
AUTHAPP_SECRETKEY=TCfTk$a!KzZd5mReXFRntBbtFZ&H6z7KB
# token validity period in minutes: default is set to 2days(2880 mins)
AUTHAPP_TOKENDURATION=2880

# Application grpc port
AUTHAPP_GRPC_PORT=7301
```

You can change the `.env` file with the required env variables of the deployment environment.

### Building and Running

#### Using the Binary

```bash
$ go build -o auth-service cmd/main.go
```

This will produce a binary named `auth-service`. The binary can be run locally, but make sure the `.env` file is in the
same directory you are running the binary.

```bash
$ ./auth-service
```

The application will start only a gRPC server on the port specified by the line `AUTHAPP_GRPC_PORT=` configured on
`.env` file. 
