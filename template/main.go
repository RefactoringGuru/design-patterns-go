package main

import "fmt"

func main() {
	// otp := otp{}

	// smsOTP := &sms{
	// 	otp: otp,
	// }

	// smsOTP.genAndSendOTP(smsOTP, 4)

	// emailOTP := &email{
	// 	otp: otp,
	// }
	// emailOTP.genAndSendOTP(emailOTP, 4)
	// fmt.Scanln()
	smsOTP := &Sms{}
	smsOTP.Otp = Otp{
		iOtp: smsOTP,
	}
	smsOTP.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	emailOTP.Otp = Otp{
		iOtp: emailOTP,
	}
	emailOTP.genAndSendOTP(4)

}
