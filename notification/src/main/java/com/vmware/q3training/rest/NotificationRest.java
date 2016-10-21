package com.vmware.q3training.rest;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.vmware.q3training.service.notification.EmailReceipt;
import com.vmware.q3training.service.notification.NotificationService;

@RestController
public class NotificationRest 
{
    
    // Member Variables
    private NotificationService notificationService = new NotificationService();

    @RequestMapping("/email")  //, defaultValue="medium"
    public EmailReceipt sendEmail(@RequestParam(value="toAddress") String toAddress, 
    		@RequestParam(value="subject") String subject,
    		@RequestParam(value="message") String message) {
    	
    	return notificationService.sendEmail(toAddress, subject, message);
    }
}
