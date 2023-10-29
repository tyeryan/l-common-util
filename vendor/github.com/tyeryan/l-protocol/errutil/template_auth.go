package errutil

import (
	"google.golang.org/grpc/codes"
)

const (
	EmailAlreadyRegistered    ErrorCode = "EMAIL_ALREADY_REGISTERED"
	EmailNotRegistered        ErrorCode = "EMAIL_NOT_REGISTERED"
	InvalidMagiclink          ErrorCode = "INVALID_MAGICLINK"
	InvalidOTP                ErrorCode = "INVALID_OTP"
	InvalidDeviceID           ErrorCode = "INVALID_DEVICE_ID"
	InvalidAppID              ErrorCode = "INVALID_APP_ID"
	InvalidUserID             ErrorCode = "INVALID_USER_ID"
	InvalidScope              ErrorCode = "INVALID_SCOPE"
	InvalidRefreshToken       ErrorCode = "INVALID_TOKEN"
	InvalidCode               ErrorCode = "INVALID_CODE"
	ForceLogout               ErrorCode = "FORCE_LOGOUT"
	NonAdminSignUp            ErrorCode = "NON_ADMIN_SIGNUP"
	CodeExpired               ErrorCode = "CODE_EXPIRED"
	SuspendedAccount          ErrorCode = "SUSPENDED_ACCOUNT"
	CodeAlreadyVerified       ErrorCode = "CODE_ALREADY_VERIFIED"
	InvitationCancelled       ErrorCode = "INVITATION_CANCELLED"
	AccountAlreadyCreated     ErrorCode = "ACCOUNT_ALREADY_CREATED"
	InvitationCancelledResend ErrorCode = "INVITATION_CANCELLED_RESEND"
	InvitationRevoked         ErrorCode = "INVITATION_REVOKED"
	AccountLockWarning        ErrorCode = "ACCOUNT_LOCK_WARNING"
	AccountLocked             ErrorCode = "ACCOUNT_LOCKED"
	AccountPermanentLocked    ErrorCode = "ACCOUNT_PERMANENT_LOCKED"
	GenerateOTPLocked         ErrorCode = "GENERATE_OTP_LOCKED"
	EmailDomainBlocked        ErrorCode = "EMAIL_DOMAIN_BLOCKED"
)

func init() {
	errorResponses := []*Template{
		{
			Code:     EmailAlreadyRegistered,
			Desc:     "email already registered",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     EmailNotRegistered,
			Desc:     "email not registered",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvalidMagiclink,
			Desc:     "invalid magiclink",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvalidOTP,
			Desc:     "invalid otp",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvalidDeviceID,
			Desc:     "invalid device id",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     InvalidAppID,
			Desc:     "invalid app id",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     InvalidUserID,
			Desc:     "invalid user id",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     InvalidScope,
			Desc:     "invalid scope value",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     InvalidRefreshToken,
			Desc:     "token is invalid or expired",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     InvalidCode,
			Desc:     "code is invalid",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     ForceLogout,
			Desc:     "user is forced logout",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     NonAdminSignUp,
			Desc:     "non admin sign up for admin",
			GrpcCode: codes.InvalidArgument,
		},
		{
			Code:     CodeExpired,
			Desc:     "code expired",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     SuspendedAccount,
			Desc:     "account suspended",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     CodeAlreadyVerified,
			Desc:     "code already verified",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvitationCancelled,
			Desc:     "invitation cancelled",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     AccountAlreadyCreated,
			Desc:     "account already created",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvitationCancelledResend,
			Desc:     "invitation cancelled resend",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     InvitationRevoked,
			Desc:     "invitation revoked",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     AccountLockWarning,
			Desc:     "account will be locked after multiple failed attempts",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     AccountLocked,
			Desc:     "account locked",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     AccountPermanentLocked,
			Desc:     "account permanent locked",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     GenerateOTPLocked,
			Desc:     "can generate OTP only after time has lapsed",
			GrpcCode: codes.PermissionDenied,
		},
		{
			Code:     EmailDomainBlocked,
			Desc:     "email domain blocked",
			GrpcCode: codes.PermissionDenied,
		},
	}
	addErrorTemplate(errorResponses)
}
