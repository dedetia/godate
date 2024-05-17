package constant

type (
	Gender         string
	PremiumPackage string
)

const MaxFileSize = 2 << (10 * 2)

const (
	GenderMale     Gender = "Male"
	GenderFemale   Gender = "Female"
	GenderEveryone Gender = "Everyone"

	PremiumPackageVerifiedLabel PremiumPackage = "verified_label"
	PremiumPackageNoSwipeQuota  PremiumPackage = "no_swipe_quota"
)

func (g Gender) String() string {
	return string(g)
}

func (g Gender) Interest() []string {
	interests := make([]string, 0)
	switch g {
	case GenderEveryone:
		interests = append(interests, GenderMale.String(), GenderFemale.String())
	default:
		interests = append(interests, g.String())
	}
	return interests
}

func (p PremiumPackage) String() string {
	return string(p)
}
