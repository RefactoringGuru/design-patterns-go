package main

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

// type otp struct {
// }

// func (o *otp) genAndSendOTP(iOtp iOtp, otpLength int) error {
// 	otp := iOtp.genRandomOTP(otpLength)
// 	iOtp.saveOTPCache(otp)
// 	message := iOtp.getMessage(otp)
// 	err := iOtp.sendNotification(message)
// 	if err != nil {
// 		return err
// 	}
// 	iOtp.publishMetric()
// 	return nil
// }

type otp struct {
	iOtp iOtp
}

func (o *otp) genAndSendOTP(otpLength int) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}
