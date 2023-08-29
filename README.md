# kafka-connect-connector-manager
as the name suggests is a small lib in _**GO**_, which allows you to create, update and delete _**Kafka Connect**_ connectors. 
Currently the lib only supports _**Debezium**_ Connectors.

# Pre-requisites
Since this is a GO Lib, you have to have **_GoLang_** installed. 

WIP: Build support to run this lib via Docker as well, so that this can be triggered by Github Actions
as well.

Additionally, you need to set values for the following env variables

```bash 

AUTH_TYPE="basic_auth" # Currently only supports basic-auth
BASIC_AUTH_USERNAME="<username>" # Auth Username of the user
BASIC_AUTH_PASSWORD="<password>" # Auth Password of the user

APP_ENV="<file-name-in-the-configs-folder>" # Name of the file in the configs folder. 
DB_HOST="<host>"; # Database Hostname
DB_NAME="<db-name>"; # Database Name 
DB_PORT="<db-port"; # Database Port
DB_USERNAME="<username>"; # Database Username  
DB_PASSWORD="<password>"; # Database Password

```
Alternative, you can club all these variables in a shell file and run the application from there.
Checkout [run_sample.sh](run_sample.sh) for reference

# Setup 

To run the project, you need to place a config file in the `configs` directory.
Once added, set the anme of the variable `APP_ENV` to the name of the file which needs to be run from config.

To add connectors, add the following object in the `connectors` section of the yml file added in the configs folder
```yml
{
  id: 6, # ID stands for the connector id to be set, unique for every connector
  name: "dbz_mysql_sample", # Connector Name
  table-name: "mydb.sample", # Name of the table to run CDC from
  type: "debezium", # Connector Type, currently only supports debezium
  topic: "warehouse.mysql.sample" # Name of the kafka topic where the table data will be written
}
```

The `connectors` section is a an array and you can add as many connectors you want from the same database.
There needs to be a different config file for each database.


# Running the project

To run the utility, you need to pass the following arguments 

```bash

--mode # Options are upsert, delete & status
--connector-name # Name of the connector for which the utility is being run
```

## Sample Commands

The sample commands will run the utility using `./run_sample` bash script

### Get Status

```./run_sample --mode status```

### Add or Update a Connector

```./run_sample --mode upsert --connector-name <connector-name>```


### Delete a Connector

```./run_sample --mode delete --connector-name <connector-name>```


