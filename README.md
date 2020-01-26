# Yet Another Social Network

Education project (homework) for High Load Architect OTUS course

# Command line options

## host

-e &lt;host>:&lt;port>

e.g.: -e 0.0.0.0:80

default: 0.0.0.0:8181

## database

-d &lt;user>:&lt;password>@&lt;host>/&lt;database name>

e.g.: -d 'admin:admin@localhost/social_network'

default: admin:admin@localhost/social_network

## database read-only replica 

-r &lt;user>:&lt;password>@&lt;host>/&lt;database name>

e.g.: -r 'admin:admin@localhost/social_network'

default: when the option is omitted, app uses the DB which
is specified in d 