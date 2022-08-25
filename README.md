# Backend

To build and run:

go build main.go

./main

API Routes:

-   /version => Returns the API version and version name
-   /upload_result => Uploads result from user
-   /resources/coins_today => Returns todays coins
-   /resources/archive/:file => Returns archival data

Database Solution:

I think the most suitable solution for this project would be a Redis Datebase acting as cache for a day,
after then it would write out to a csv file and archived for further use.

User -> Redis cache - ( After 24 h ) > csv file archive

Some would prefer writing to a relational database like postgresql, but I think this solution is highly scalable with minimal effort and maintenance

Redis:  
What we need to store in redis is basically this:  
{ "Name":"Bitcoin","Price":"0.1","Action": true/false,"Hour":"19:45" }
