package domain

type PurchasePremium struct {
	PackageType string `json:"package_type" validate:"required,oneof=no_swipe_quota verified_label"`
}
