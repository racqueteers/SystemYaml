package racqueteers

type TeamD struct {
	EmergingRacqueteers EmergingRacqueteers `yaml:"emergingRacqueteers"`
}
type Players struct {
	Alpha []string `yaml:"alpha"`
	Beta  []string `yaml:"beta"`
}
type Team struct {
	Name    string  `yaml:"name"`
	Players Players `yaml:"players"`
	Slogan  string  `yaml:"slogan"`
	Captain string  `yaml:"captain"`
}
type Logo struct {
	URL    string   `yaml:"url"`
	Colors []string `yaml:"colors"`
}
type Photos struct {
	URL string `yaml:"url"`
}
type Videos struct {
	URL string `yaml:"url"`
}
type Schedule struct {
	Match          int    `yaml:"match"`
	Day            string `yaml:"day"`
	Date           string `yaml:"date"`
	OppositionTeam string `yaml:"oppositionTeam"`
	Type           string `yaml:"type"`
	Players        string `yaml:"players"`
}
type EmergingRacqueteers struct {
	Team       Team       `yaml:"team"`
	Logo       Logo       `yaml:"logo"`
	Photos     Photos     `yaml:"photos"`
	Videos     Videos     `yaml:"videos"`
	Schedule   []Schedule `yaml:"schedule"`
	Meetings   int        `yaml:"meetings"`
	TotalPosts int        `yaml:"totalPosts"`
}
