The NotificationService is a Java based service with a REST endpoint that allows
for the send of email notifications.

To build, insteall gradle 2.2+ and run the following command from notfication root dir:
gradle build

To run, execute the following command from nofication root dir:
java -jar build/libs/notification-rest-service-0.1.0.jar

Try sending an email via the following:

http://localhost:8080/email?toAddress=jjarrett@vmware.com&subject=TestSubject&message=YourMessage