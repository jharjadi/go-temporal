# go-temporal

gcloud components update
gcloud config set auth/disable_credentials true
gcloud config set project tms-local
gcloud config set api_endpoint_overrides/spanner http://localhost:9020/

gcloud spanner instance-configs list

# Create instance
gcloud spanner instances create test-instance --config=emulator-config --description="Test Instance" --nodes=1
# Create a database
gcloud spanner databases create test-database --instance test-instance --ddl "CREATE TABLE TestTable (Key INT64, Value STRING(MAX)) PRIMARY KEY (Key)"
# Write into database
gcloud spanner rows insert --table=TestTable --database=test-database --instance=test-instance --data="Key=1,Value=TestValue1"
# Read from database
gcloud spanner databases execute-sql test-database --instance test-instance --sql "select * from TestTable"