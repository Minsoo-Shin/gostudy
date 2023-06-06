package errors

const (
	ErrAuthUnauthorized = "AUTH_UNAUTHORIZED"
	ErrAuthForbidden    = "AUTH_FORBIDDEN"

	ErrOutfitNotFound           = "OUTFIT_NOT_FOUND"
	ErrCommentNotFound          = "OUTFIT_COMMENT_NOT_FOUND"
	ErrReportDuplicated         = "OUTFIT_REPORT_DUPLICATED"
	ErrOutfitContentLimitExceed = "OUTFIT_CONTENT_LIMIT_EXCEED"
	OutfitClothesNotFound       = "OUTFIT_CLOTHES_NOT_FOUND"
	RedisNotFound               = "REDIS_NOT_FOUND"
	ErrMarketItemNotFound       = "MARKET_ITEM_NOT_FOUND"

	ErrUserNotFound                    = "USER_NOT_FOUND"
	ErrUserTokenInvalid                = "USER_TOKEN_INVALID"
	ErrUserPasswordIncorrect           = "USER_PASSWORD_INCORRECT"
	ErrUserNicknameDuplicated          = "USER_NICKNAME_DUPLICATED"
	ErrUserMobileDuplicated            = "USER_MOBILE_DUPLICATED"
	ErrMobileVerificationLimitExceeded = "MOBILE_VERIFICATION_LIMIT_EXCEEDED"
	ErrMobileVerificationExpired       = "MOBILE_VERIFICATION_EXPIRED"
	ErrMobileVerificationNotMatched    = "MOBILE_VERIFICATION_NOT_MATCHED"
	ErrMobileSmsFailToSend             = "MOBILE_SMS_FAIL_TO_SEND"

	ErrEmailVerificationNotMatched   = "EMAIL_VERIFICATION_NOT_MATCHED"
	ErrBanUserNotFound               = "BAN_USER_NOT_FOUND"
	ErrBanDuplicated                 = "BAN_DUPLICATED"
	ErrMarketItemReviewAlreadyExists = "MARKET_ITEM_REVIEW_ALREADY_EXISTS"
	ErrNotBuyerAndSeller             = "NOT_BUYER_AND_SELLER"
	ErrPrivateOutfit                 = "PRIVATE_OUTFIT"
	ErrDeletedOutfit                 = "DELETED_OUTFIT"
	ErrInteractBannedUser            = "INTERACT_BANNED_USER"
	ErrClothesNotFound               = "CLOTHES_NOT_FOUND"
	ErrCountryCodeNotFound           = "COUNTRY_CODE_NOT_FOUND"
	ErrCurrencyNotFound              = "CURRENCY_NOT_FOUND"
	ErrMarketItemNotSold             = "MARKET_ITEM_NOT_SOLD"
	ErrMarketItemBookmarkNotFound    = "MARKET_ITEM_BOOKMARK_NOT_FOUND"
	ErrMarketItemAlreadyBookmarked   = "MARKET_ITEM_ALREADY_BOOKMARKED"
	ErrMarketItemReviewNotFound      = "MARKET_ITEM_REVIEW_NOT_FOUND"

	ErrSendbirdTokenNotFound      = "SENDBIRD_TOKEN_NOT_FOUND"
	ErrSendbirdTokenInvalid       = "SENDBIRD_TOKEN_INVALID"
	ErrSendbirdTokenAlreadyExists = "SENDBIRD_TOKEN_ALREADY_EXISTS"
	ErrSendbirdUnknown            = "SENDBIRD_UNKNOWN"

	ErrInferenceFailed = "INFERENCE_FAILED"

	ErrInvalidShopID         = "INVALID_SHOP_ID"
	ErrInvalidShopTaskID     = "INVALID_SHOP_TASK_ID"
	ErrInvalidShopTaskItemID = "INVALID_SHOP_TASK_ITEM_ID"
	ErrShopNotFound          = "SHOP_NOT_FOUND"
	ErrShopTaskItemNotFound  = "SHOP_TASK_ITEM_NOT_FOUND"
)
