package com.vmware.q3training.service.notification;

public class EmailReceipt {

		private boolean success;
		private String error;
		public boolean isSuccess() {
			return success;
		}
		public void setSuccess(boolean success) {
			this.success = success;
		}
		public String getError() {
			return error;
		}
		public void setError(String error) {
			this.error = error;
		}
		
}
