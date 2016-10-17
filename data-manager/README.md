# data-manager
Microservice to manage the persistance of lab data to a blob service.

# Getting Started
Include this repository in your GOPATH at src/github.com/vmtrain/data-manager, then build it:
    cd data-manager
    ./build.sh containerize

Then, you can start the microservice:

	./build.sh start

Note: The microservice will start on the default port of 6001

To stop the microservice:

	./build.sh stop

## The API

**Create a Reservation**

`POST /reservations`

Request body:

```
{
  "name": "my reservation",
  "start_date": "2016-01-01T06:00:00.000Z",
  "end_date": "2099-12-31T00:00:00.000Z",
  "server_name": "server1"
}
```

Response codes:

HTTP code   | Description
----------- | -------------
200         | Indicates the reservation was created.
400         | Indicates the reservation was not created due to invalid payload.
500         | Indicates request was unsucessful due to an unexpected condition.

Success response body:

```
{
  "uuid": "dbbeb1e7-2847-4886-b5cb-9690ad29e16c",
  "name": "my reservation",
  "start_date": "2016-01-01T06:00:00.000Z",
  "end_date": "2099-12-31T00:00:00.000Z",
  "server_name": "server1",
  "approved": false
}
```
**Get a List of Reservations**

`GET /reservations`

Response codes:

HTTP code | Description
--------- | -------------
200       | Indicates request was successful and the reservations are returned.
500       | Indicates request was unsucessful due to an unexpected condition.


Success response body:

```
[
  {
    "uuid": "b8e70070-a2d0-4458-965a-67d29a632090",
    "name": "my reservation",
    "start_date": "2016-01-01T06:00:00.000Z",
    "end_date": "2099-12-31T00:00:00.000Z",
    "server_name": "server1",
    "approved": false
  },
  ...
]
```

**Create a Server**

`POST /servers`

Request body:

```
{
  "name": "server1"
}
```

Response codes:

HTTP code   | Description
----------- | -------------
200         | Indicates the server was created.
400         | Indicates the server was not created due to invalid payload.
500         | Indicates request was unsucessful due to an unexpected condition.

Success response body:

```
{
  "uuid": "c79fe768-d1a3-4ff8-bfdc-f8e48ea39837",
  "name": "server1"
}
```

**Get a List of Servers**

`GET /servers`

Response codes:

HTTP code | Description
--------- | -------------
200       | Indicates request was successful and the servers are returned.
500       | Indicates request was unsucessful due to an unexpected condition.


Success response body:

```
[
  {
    "uuid": "c79fe768-d1a3-4ff8-bfdc-f8e48ea39837",
    "name": "server1"
  },
  ...
]
```

**Create a User**

`POST /users`

Request body:

```
{
  "name": "user1"
}
```

Response codes:

HTTP code   | Description
----------- | -------------
200         | Indicates the user was created.
400         | Indicates the user was not created due to invalid payload.
500         | Indicates request was unsucessful due to an unexpected condition.

Success response body:

```
{
  "uuid": "e70d2787-ea9f-4902-acf7-f993554b80b4",
  "name": "user1"
}
```

**Get a List of Users**

`GET /users`

Response codes:

HTTP code | Description
--------- | -------------
200       | Indicates request was successful and the users are returned.
500       | Indicates request was unsucessful due to an unexpected condition.


Success response body:

```
[
  {
    "uuid": "e70d2787-ea9f-4902-acf7-f993554b80b4",
    "name": "user1"
  },
  ...
]
```

# Dependencies
This service requires a valid Go language environment.

# License and Author
Copyright: Copyright (c) 2015 VMware, Inc. All Rights Reserved

Author: Luis M. Valerio, VMware, Inc.

License: MIT

For details of the license, see the LICENSE file.
