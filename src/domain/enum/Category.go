package enum


type Category int

const (
	Influencer Category = iota
	Sports
	NewMedia
	Business
	Brand
	Organization
	Others
)

func (g Category) String() string {
	return [...]string{"Influencer", "Sports", "NewMedia", "Business", "Brand", "Organization", "Others" }[g]
}


