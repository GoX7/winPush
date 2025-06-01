package winpush

type Actions struct {
	Content        string
	Arguments      string
	Icon           string
	ActivationType string
	Placement      string
}

type Notificator struct {
	AppID               string
	Title               string
	Subtitle            string
	Message             string
	Icon                string
	Actions             []Actions
	ActivationType      string
	ActivationArguments string
	Duration            string
}
