package pipewire

type Node struct {
	Id          int
	Type        *string
	Version     int
	Permissions []string
	Info        info
}

type info struct {
	Cookie     int
	UserName   *string
	HostName   *string
	Version    *string
	Name       *string
	ChangeMask []string
	Props      infoProps
}

type infoProps struct {
	ConfigName                  *string `json:"config.name"`
	LinkMaxBuffers              int     `json:"link.max-buffers"`
	CoreDaemon                  bool    `json:"core.daemon"`
	CoreName                    *string `json:"core.name"`
	DefaultClockMinQuantum      int     `json:"default.clock.min-quantum"`
	CpuMaxAlign                 int     `json:"cpu.max-align"`
	DefaultClockRate            int     `json:"default.clock.right"`
	DefaultClockQuantum         int     `json:"default.clock.quantum"`
	DefaultClockMaxQuantum      int     `json:"default.clock.max-quantum"`
	DefaultClockQuantumLimit    int     `json:"default.clock.quantum-limit"`
	DefaultVideoWidth           int     `json:"default.video.width"`
	DefaultVideoHeight          int     `json:"default.video.height"`
	DefaultVideoRateNum         int     `json:"default.video.rate.num"`
	DefaultVideoRateDenom       int     `json:"default.video.rate.denom"`
	LogLevel                    int     `json:"log.level"`
	ClockPowerOfTwoQuantum      bool    `json:"clock.power-of-two-quantum"`
	MemWarnMLock                bool    `json:"mem.warn-mlock"`
	MemAllowMLock               bool    `json:"mem.allow-mlock"`
	SettingsCheckQuantum        bool    `json:"settings.check-quantum"`
	SettingsCheckRate           bool    `json:"settings.check-rate"`
	ObjectId                    int     `json:"object.id"`
	ObjectSerial                int     `json:"object.serial"`
	ModuleName                  *string `json:"module.name"`
	ModuleAuthor                *string `json:"module.author"`
	ModuleDescription           *string `json:"module.description"`
	ModuleUsage                 *string `json:"module.usage"`
	ModuleVersion               *string `json:"module.version"`
	NiceLevel                   int     `json:"nice.level"`
	ModuleId                    int     `json:"module.id"`
	FactoryName                 *string `json:"factory.name"`
	FactoryTypeName             *string `json:"factory.type.name"`
	FactoryTypeVersion          int     `json:"factory.type.version"`
	FactoryUsage                *string `json:"factory.usage"`
	NodeName                    *string `json:"node.name"`
	NodeGroup                   *string `json:"node.group"`
	PriorityDriver              int     `json:"priority.driver"`
	FactoryId                   int     `json:"factory.id"`
	ClockQuantumLimit           int     `json:"clock.quantum-limit"`
	NodeDriver                  bool    `json:"node.driver"`
	NodeFreewheel               bool    `json:"node.freewheel"`
	PipewireProtocol            *string `json:"pipewire.protocol"`
	PipewireSecPid              int     `json:"pipewire.sec.pid"`
	PipewireSecUid              int     `json:"pipewire.sec.uid"`
	PipewireSecGid              int     `json:"pipewire.sec.gid"`
	ClientApi                   *string `json:"client.api"`
	PulseServerType             *string `json:"pulse.server.type"`
	ApplicationProcessId        int     `json:"application.process.id"`
	ApplicationProcessUser      *string `json:"application.process.user"`
	ApplicationProcessHost      *string `json:"application.process.host"`
	ApplicationProcessBinary    *string `json:"application.process.binary"`
	ApplicationName             *string `json:"application.name"`
	ApplicationLanguage         *string `json:"application.language"`
	ApplicationProcessMachineId *string `json:"application.process.machine-id"`
	CoreVersion                 *string `json:"core.version"`
	PipewireAccess              *string `json:"pipewire.access"`
	ObjectLinger                bool    `json:"object.linger"`
	MediaClass                  *string `json:"media.class"`
	AudioChannels               int     `json:"audio.channels"`
	AudioPosition               *string `json:"audio.position"`
	NodeDescription             *string `json:"node.description"`
	MonitorChannelVolumes       bool    `json:"monitor.channel-volumes"`
	PulseModuleId               int     `json:"pulse.module.id"`
	FactoryMode                 *string `json:"factory.mode"`
	AudioAdaptFollower          *string `json:"audio.adapt.follower"`
	LibraryName                 *string `json:"library.name"`
	WireplumberDaemon           bool    `json:"wireplumber.daemon"`
	WireplumberExportCore       bool    `json:"wireplumber.export-core"`
	WireplumberScriptEngine     *string `json:"wireplumber.script-engine"`
	ClientId                    int     `json:"client.id"`
	DeviceApi                   *string `json:"device.api"`
	FormatDsp                   *string `json:"format.dsp"`
	ObjectPath                  *string `json:"object.path"`
	PortName                    *string `json:"port.name"`
	PortAlias                   *string `json:"port.alias"`
	PortId                      int     `json:"port.id"`
	PortPhysical                bool    `json:"port.physical"`
	PortTerminal                bool    `json:"port.terminal"`
	PortDirection               *string `json:"port.direction"`
	NodeId                      int     `json:"node.id"`
	ApiAlsaCard                 int     `json:"api.alsa.card"`
	LoopCancel                  bool    `json:"loop.cancel"`
	ClientName                  *string `json:"client.name"`
	JackShortName               bool    `json:"jack.short-name"`
	JackSelfConnectMode         *string `json:"jack.self-connect-mode"`
	MediaType                   *string `json:"media.type"`
	MediaCategory               *string `json:"media.category"`
	MediaRole                   *string `json:"media.role"`
	NodeAlwaysProcess           bool    `json:"node.always-process"`
	NodeLockQuantum             bool    `json:"node.lock-quantum"`
	NodeTransportSync           bool    `json:"node.transport.sync"`
	ApiV4L2Path                 *string `json:"api.v4l2.path"`
	DeviceBusPath               *string `json:"device.bus-path"`
	DeviceCapabilities          *string `json:"device.capabilities"`
	DeviceDescription           *string `json:"device.description"`
	DeviceEnumApi               *string `json:"device.enum.api"`
	DeviceName                  *string `json:"device.name"`
	DevicePluggedUsec           *string `json:"device.plugged.usec"`
	DeviceProductName           *string `json:"device.product.name"`
	DeviceSubsystem             *string `json:"device.subsystem"`
	DeviceSysfsPath             *string `json:"device.sysfs.path"`
	ApiV4L2CapDriver            *string `json:"api.v4l2.cap.driver"`
	ApiV4L2CapCard              *string `json:"api.v4l2.cap.card"`
	ApiV4L2CapBusInfo           *string `json:"api.v4l2.cap.bus_info"`
	ApiV4L2CapVersion           *string `json:"api.v4l2.cap.version"`
	ApiV4L2CapCapabilities      int     `json:"api.v4l2.cap.capabilities"`
	ApiV4L2CapDeviceCaps        int     `json:"api.v4l2.cap.device-caps"`
	ApiAcpAutoPort              bool    `json:"api.acp.auto-port"`
	ApiAcpAutoProfile           bool    `json:"api.acp.auto-profile"`
	ApiAlsaCardLongname         *string `json:"api.alsa.card.longname"`
	ApiAlsaCardName             *string `json:"api.alsa.card.name"`
	ApiAlsaPath                 *string `json:"api.alsa.path"`
	ApiAlsaUseAcp               bool    `json:"api.alsa.use-acp"`
	ApiDbusReserveDevice1       *string `json:"api.dbus.ReserveDevice1"`
	DeviceBus                   *string `json:"device.bus"`
	DeviceBusId                 *string `json:"device.bus-id"`
	DeviceIconName              *string `json:"device.icon-name"`
	DeviceNick                  *string `json:"device.nick"`
	DeviceProductId             int     `json:"device.product.id"`
	DeviceSerial                *string `json:"device.serial"`
	DeviceVendorId              int     `json:"device.vendor.id"`
	DeviceVendorName            *string `json:"device.vendor.name"`
	AlsaCard                    int     `json:"alsa.card"`
	AlsaCardName                *string `json:"alsa.card_name"`
	AlsaLongCardName            *string `json:"alsa.long_card_name"`
	AlsaDriverName              *string `json:"alsa.driver_name"`
	DeviceString                int     `json:"device.string"`
	DeviceFormFactor            *string `json:"device.form-factor"`
	DeviceId                    *string `json:"device.id"`
	NodePauseOnIdle             bool    `json:"node.pause-on-idle"`
	AudioChannel                *string `json:"audio.channel"`
	PortMonitor                 bool    `json:"port.monitor"`
	AlsaClass                   *string `json:"alsa.class"`
	AlsaDevice                  int     `json:"alsa.device"`
	AlsaId                      *string `json:"alsa.id"`
	AlsaName                    *string `json:"alsa.name"`
	AlsaResolutionBits          int     `json:"alsa.resolution_bits"`
	AlsaSubclass                *string `json:"alsa.subclass"`
	AlsaSubdevice               int     `json:"alsa.subdevice"`
	AlsaSubdeviceName           *string `json:"alsa.subdevice_name"`
}
