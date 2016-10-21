package com.vmware.q3training.service.notification;

import java.util.Properties;

import javax.mail.Message;
import javax.mail.MessagingException;
import javax.mail.Session;
import javax.mail.Transport;
import javax.mail.internet.InternetAddress;
import javax.mail.internet.MimeMessage;

public class NotificationService {

	
	public EmailReceipt sendEmail(String toAddress, String subject, String message) {
		EmailReceipt emailReceipt = new EmailReceipt();
	
		String from = "q3training@vmware.com";
        String host = "localhost";

        // Get system properties
        Properties properties = System.getProperties();
        properties.setProperty("mail.smtp.host", host);
        Session session = Session.getDefaultInstance(properties);

        try {
            MimeMessage mimeMessage = new MimeMessage(session);
            mimeMessage.setFrom(new InternetAddress(from));
            mimeMessage.addRecipient(Message.RecipientType.TO, new InternetAddress(toAddress));
            mimeMessage.setSubject(subject);
            mimeMessage.setText(message);

            // Send message
            Transport.send(mimeMessage);
            System.out.println("Sent message successfully....");
            emailReceipt.setSuccess(true);
            emailReceipt.setError("No error, email was sent successfully!");
         }catch (MessagingException mex) {
        	emailReceipt.setSuccess(false);
        	emailReceipt.setError(mex.getMessage());
            mex.printStackTrace();
         }  
  		return emailReceipt;
	}
	
}
