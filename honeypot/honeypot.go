package honeypot

var TAGS = [...]string{
	"health",
	"eyes",
	"brain",
	"kidney",
	"hair",
	"ear",
	"nose",
	"lungs",
	"teeth",
	"muscle",
	"heart",
	"cancer",
	"diabates",
}

type HoneyPot interface {
	PopulatePot(string)
	Download()
}
