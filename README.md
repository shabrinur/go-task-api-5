# Assignment: Go REST API
Personal repository for basic Go REST API assignment.

## Configuration
Configuration for this project can be found in file **/config-*.properties**. To set which configuration file to load, modify this line in **/app/config/Configuration.go**:

    viper.SetConfigName("set configuration file name here")

Project requires **PostgreSQL** to run.

## Documentation

Swagger UI for this application can be accessed from path [/swagger/index.html](#). Documentation for API requests & responses can be found under **/docs**.