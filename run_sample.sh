export AUTH_TYPE="basic_auth"
export BASIC_AUTH_USERNAME="<username>"
export BASIC_AUTH_PASSWORD="<password>"
export DB_PORT="3306";

export APP_ENV="sample"
export DB_HOST="<host>";
export DB_NAME="sample";
export DB_USERNAME="<username>";
export DB_PASSWORD="<password>";

go run main.go $@
