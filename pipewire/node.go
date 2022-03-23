package pipewire

type Node struct {
	Id          int      `json:"id"`
	Type        *string  `json:"type"`
	Version     int      `json:"version"`
	Permissions []string `json:"permissions"`
	Info        info     `json:"info"`
}

type info struct {
	Cookie     int       `json:"cookie"`
	UserName   string    `json:"user-name"`
	HostName   string    `json:"host-name"`
	Name       string    `json:"name"`
	ChangeMask []string  `json:"change-mask"`
	Props      infoProps `json:"props"`
}

type infoProps struct {
	LogLevel                 int    `json:"log.level"`
	NodeName                 string `json:"node.name"`
	ClientApi                string `json:"client.api"`
	ApplicationProcessId     int    `json:"application.process.id"`
	ApplicationProcessUser   string `json:"application.process.user"`
	ApplicationProcessHost   string `json:"application.process.host"`
	ApplicationProcessBinary string `json:"application.process.binary"`
	ApplicationName          string `json:"application.name"`
	MediaClass               string `json:"media.class"`
	AudioChannels            int    `json:"audio.channels"`
	NodeDescription          string `json:"node.description"`
	PulseModuleId            int    `json:"pulse.module.id"`
	ClientId                 int    `json:"client.id"`
	NodeId                   int    `json:"node.id"`
	ClientName               string `json:"client.name"`
	MediaType                string `json:"media.type"`
	MediaCategory            string `json:"media.category"`
	MediaRole                string `json:"media.role"`
}
