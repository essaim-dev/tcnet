package tcnet

type ApplicationCode uint32

const (
	ApplicationCodePublic                     ApplicationCode = 0x0000
	ApplicationCodePioneerDJ                  ApplicationCode = 0x0AA0 // http://www.pioneerdj.com
	ApplicationCodeTCSupplyShowKontrol        ApplicationCode = 0x0AAA // http://www.showkontrol.com
	ApplicationCodeTCSupplyPyrotechnicSystems ApplicationCode = 0x0AAB // http://www.tc-supply.com
	ApplicationCodeTCSupplyRideControlSystems ApplicationCode = 0x0AAC // http://www.tc-supply.com
	ApplicationCodeAvolitesLighting           ApplicationCode = 0x0AB0 // http://www.avolites.com
	ApplicationCodeMALighting                 ApplicationCode = 0x0AB1 // http://www.malighting.com
	ApplicationCodeChamsysLighting            ApplicationCode = 0x0AB3 // http://www.chamsys.co.uk
	ApplicationCodeObsidianControl            ApplicationCode = 0x0AB4 // http://www.obsidiancontrol.com
	ApplicationCodeArkaosSoftware             ApplicationCode = 0x0ABA // http://www.arkaos.net
	ApplicationCodeBLCKBOOKTimeCodeSync       ApplicationCode = 0x0ABB // http://www.timecodesync.com
	ApplicationCodeResolumeSoftware           ApplicationCode = 0x0ABC // http://www.resolume.com
	ApplicationCodeGreenHippo                 ApplicationCode = 0x0ABD // http://www.green-hippo.com
	ApplicationCodeRDShowCockpit              ApplicationCode = 0x0ABE // http://www.showcockpit.com
	ApplicationCodeDisguise                   ApplicationCode = 0x0ABF // http://disguise.one
	ApplicationCodeOrangePI                   ApplicationCode = 0x0ACA // http://orangepi.dmx.org
	ApplicationCodeRedPillVR                  ApplicationCode = 0x0ACB // http://www.redpillvr.com
	ApplicationCodePublic2                    ApplicationCode = 0xFFFF
)
