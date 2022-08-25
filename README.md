# Cryptoguess Backend

## Running and building API

Run the API using `go run main.go`

Build it by running `go build -o cryptoguesser-api main.go`

That should give you the executable for deployment

## How does it work?

### Boot up

When the API starts it tries to connect to the Redis database on localhost:6379 without password

Then it checks if the archive has been set up and if not creates one

After that it checks for remaining transactions in redis and saves them to file

Next step involves choosing today's coins which are picked from coingecko API

After the coins have been picked it schedules asynchronous tasks such as:

-   Updating coin prices every 15 minutes
-   Saving transactions to file from Redis every day at 00:00

Meanwhile the router prepares the necessery routes and CORS settings for functioning properly

### User submits TX

When user submits transaction and it passes checks it is written to the Redis database and the user is returned a status code and a response corresponding to the result.

## Enviroment variables:

API uses .env file as a store for secrets

For now these include cookie salt and hashed admin password

Of course this will be changed in the future.

Example .env:

    COOKIESECRET=b826937235ee76c752a08843ca55621443b6a88ba385617e5283f7867581c75b
    ADMINPASSWORD=b826937235ee76c752a08843ca55621443b6a88ba385617e5283f7867581c75b

## API Routes:

-   /version => Returns the API version and version name
-   /upload_result => Uploads result from user ( Requires session )
-   /resources/coins_today => Returns todays coins
-   /resources/archive/:file => Returns archival data ( Admins only)
-   /auth/login => Logs in the session
-   /auth/logout => Logouts session
-   /auth/test => Returns status of the session
