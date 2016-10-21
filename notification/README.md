Notification Service
====================

The NotificationService is a Java based service with a REST endpoint that allows for the sending of email notifications.

Building 
--------
Requires JDK 1.7 or greater. 

To build, install gradle 2.2+ and run the following command from notfication root dir: 
```
gradle build
```

Running
-------
To run, execute the following command from nofication root dir:
```
java -jar build/libs/notification-rest-service-0.1.0.jar
```

Testing
-------
Try sending an email via the following:
```
http://localhost:8080/email?toAddress=jjarrett@vmware.com&subject=TestSubject&message=YourMessage
```
