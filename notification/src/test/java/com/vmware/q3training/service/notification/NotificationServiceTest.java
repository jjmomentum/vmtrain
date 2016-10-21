package com.vmware.q3training.service.notification;

import org.junit.Test;
import org.junit.Assert;

public class NotificationServiceTest {

	@Test
	public void testSendEmailSuccess() throws Exception {
		NotificationService notificationService = new NotificationService();
		EmailReceipt emailReceipt = notificationService.sendEmail("jjarrett@vmware.com", "Test Success Send", "Someone is testing your code!");
		
		Assert.assertNotNull("EmailReceipt cannot be null if sent successfully!", emailReceipt);
		Assert.assertTrue("Email success must be true if sent successfully!", emailReceipt.isSuccess());
	}
}
