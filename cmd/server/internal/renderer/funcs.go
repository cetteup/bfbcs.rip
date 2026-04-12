package renderer

import (
	"fmt"
	"html/template"
	"math"
	"strconv"
	"strings"
	"time"
)

func seq(start, end int) []int {
	n := end - start
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, start+i)
	}
	return s
}

func add(a, b float64) float64 {
	return a + b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func gt(a, b float64) bool {
	return a > b
}

func gte(a, b float64) bool {
	return a >= b
}

func formatNumber(value float64) string {
	i := int64(value)
	s := strconv.FormatInt(i, 10)
	var parts []string
	for len(s) > 3 {
		parts = append([]string{s[len(s)-3:]}, parts...)
		s = s[:len(s)-3]
	}
	if s != "" {
		parts = append([]string{s}, parts...)
	}
	return strings.Join(parts, " ")
}

func formatDuration(value float64) string {
	i := int(value)
	hours := i / 3600
	minutes := (i % 3600) / 60
	seconds := i % 60
	return fmt.Sprintf("%0.2dh %0.2dm %0.2ds", hours, minutes, seconds)
}

func formatTime(value string) string {
	t, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return value
	}
	return t.Format("2006-01-02 15:04")
}

func getRankName(rank float64) string {
	switch int(rank) {
	case 0:
		return "RECRUIT"
	case 1:
		return "PRIVATE I"
	case 2:
		return "PRIVATE II"
	case 3:
		return "PRIVATE III"
	case 4:
		return "SPECIALIST I"
	case 5:
		return "SPECIALIST II"
	case 6:
		return "SPECIALIST III"
	case 7:
		return "CORPORAL I"
	case 8:
		return "CORPORAL II"
	case 9:
		return "CORPORAL III"
	case 10:
		return "SERGEANT I"
	case 11:
		return "SERGEANT II"
	case 12:
		return "SERGEANT III"
	case 13:
		return "STAFF SERGEANT I"
	case 14:
		return "STAFF SERGEANT II"
	case 15:
		return "STAFF SERGEANT III"
	case 16:
		return "MASTER SERGEANT I"
	case 17:
		return "MASTER SERGEANT II"
	case 18:
		return "MASTER SERGEANT III"
	case 19:
		return "FIRST SERGEANT I"
	case 20:
		return "FIRST SERGEANT II"
	case 21:
		return "FIRST SERGEANT III"
	case 22:
		return "WARRANT OFFICER I"
	case 23:
		return "WARRANT OFFICER II"
	case 24:
		return "WARRANT OFFICER III"
	case 25:
		return "CHIEF WARRANT OFFICER I"
	case 26:
		return "CHIEF WARRANT OFFICER II"
	case 27:
		return "CHIEF WARRANT OFFICER III"
	case 28:
		return "SECOND LIEUTENANT I"
	case 29:
		return "SECOND LIEUTENANT II"
	case 30:
		return "SECOND LIEUTENANT III"
	case 31:
		return "FIRST LIEUTENANT I"
	case 32:
		return "FIRST LIEUTENANT II"
	case 33:
		return "FIRST LIEUTENANT III"
	case 34:
		return "CAPTAIN I"
	case 35:
		return "CAPTAIN II"
	case 36:
		return "CAPTAIN III"
	case 37:
		return "MAJOR I"
	case 38:
		return "MAJOR II"
	case 39:
		return "MAJOR III"
	case 40:
		return "LIEUTENANT COLONEL I"
	case 41:
		return "LIEUTENANT COLONEL II"
	case 42:
		return "LIEUTENANT COLONEL III"
	case 43:
		return "COLONEL I"
	case 44:
		return "COLONEL II"
	case 45:
		return "COLONEL III"
	case 46:
		return "BRIGADIER GENERAL I"
	case 47:
		return "BRIGADIER GENERAL II"
	case 48:
		return "BRIGADIER GENERAL III"
	case 49:
		return "GENERAL"
	case 50:
		return "GENERAL OF THE ARMY"
	default:
		return fmt.Sprintf("%.0f", rank)
	}
}

type weapon struct {
	Index       int
	Key         string
	Name        string
	Description template.HTML // Using template.HTML to render special characters correctly
	Image       string
	IsPistol    bool
}

func getWeapons(platform string) []weapon {
	// Type 88 sniper uses a different key on consoles
	qbx88 := "qbu88"
	if platform == "ps3" || platform == "xbox360" {
		qbx88 = "qby88"
	}

	return []weapon{
		{Index: 0, Key: "aek", Name: "AEK-971 Vintovka", Description: "An assault rifle capable of 800 rounds per minute equipped with a 30 round magazine and recoil damper.", Image: "aek971.png"},
		{Index: 1, Key: "xm8", Name: "XM8 Prototype", Description: "A 30 round Experimental fully automatic assault rifle capable of firing 750 rounds per minute.", Image: "xm8.png"},
		{Index: 2, Key: "f2000", Name: "F2000 Assault", Description: "A bullpup 30 round fully automatic assault rifle that enables the operator to utilize the weapon in any mission.", Image: "f2000.png"},
		{Index: 3, Key: "aug", Name: "Stg.77 AUG", Description: "A durable fully automatic 30 round assault rifle possessing low accuracy which is off set by high mobility.", Image: "aug.png"},
		{Index: 4, Key: "an94", Name: "AN-94 Abakan", Description: "A high powered two round burst assault rifle suited for long range combat while firing 600 rounds per minute.", Image: "an94.png"},
		{Index: 5, Key: "m416", Name: "M416", Description: "A 30 round fully automatic assault rifle capable of firing 700 rounds per minute.", Image: "hk416.png"},
		{Index: 6, Key: "m16", Name: "M16A2", Description: "A 3 round burst assault rifle boasting a high accuracy rate that counters lack of mobility.", Image: "m16.png"},
		{Index: 7, Key: "9a91", Name: "9A-91 Avtomat", Description: "20 round, fully automatic carbine that's suited for close quarters combat packing significant stopping power.", Image: "9a91.png"},
		{Index: 8, Key: "scar", Name: "SCAR-L Carbine", Description: "Fully automatic 30 round carbine possessing stellar mobility making it suitable for close quarters.", Image: "fnscarl.png"},
		{Index: 9, Key: "xm8c", Name: "XM8 Compact", Description: "30 round fully automatic prototype carbine rifle which sacrifices accuracy for a high level of mobility.", Image: "xm8c.png"},
		{Index: 10, Key: "aks74u", Name: "AKS-74U Krinkov", Description: "30 round fully automatic carbine capable of 750 rpm. Low accuracy is countered by high mobility and silencer.", Image: "aks74u.png"},
		{Index: 11, Key: "uzi", Name: "UZI", Description: "32 round fully automatic sub machine gun equipped with silencer capable of firing 900 rounds per minute.", Image: "imiuzi.png"},
		{Index: 12, Key: "pp2", Name: "PP-2000 Avtomat", Description: "40 round fully automatic submachine gun that's lack of power is countered by high mobility and silencer.", Image: "pp2000.png"},
		{Index: 13, Key: "ump", Name: "UMP-45", Description: "25 round fully automatic sub machine gun possessing optimal mobility suiting it for close quarters combat.", Image: "ump.png"},
		{Index: 14, Key: "pkm", Name: "PKM LMG", Description: "A 7.62x54mm GPMG with a 100 round belt. It fires 650 rounds/min and is capable of engaging targets at long range.", Image: "pkm.png"},
		{Index: 15, Key: "m249", Name: "M249 Saw", Description: "200 round belt feed, 5.56 mm light machine gun firing 800 rounds/min. with an effective range of 300 to 1,000 meters.", Image: "m249fnminimi.png"},
		{Index: 16, Key: "qju88", Name: "Type 88 LMG", Description: "200 round belt feed, 5.8mm LMG firing 650 rounds/min. It's long range offsets its high recoil and heavy weight.", Image: "qjy88.png"},
		{Index: 17, Key: "m60", Name: "M60 LMG", Description: "7.62mm LMG firing 550 rounds/min from a 100 round belt. It provides lethal, accurate, long range fire.", Image: "m60.png"},
		{Index: 18, Key: "xm8lmg", Name: "XM8 LMG", Description: "LMG variant of the XM8 featuring a 100 round drum, bipod and integrated scope. It has a cyclic rate of 750 rounds/min.", Image: "xm8lmg.png"},
		{Index: 19, Key: "mg36", Name: "MG36", Description: "5.56mm LMG variant of the G36 utilizing a 100 round dual drum magazine, bipod and integrated scope. It has a cyclic rate of 750 rounds/min.", Image: "mg36.png"},
		{Index: 20, Key: "mg3", Name: "MG3", Description: "German LMG firing 1,000 rnds/min from a 7.62mm, 100 round belt it puts an incredible amount of rounds on target.", Image: "mg3.png"},
		{Index: 21, Key: "m24", Name: "M24 Sniper", Description: "A 7.62 mm bolt-action sniper rifle capable of engaging targets at extreme ranges with incredible stopping power. Capacity: 5 rounds.", Image: "m24.png"},
		{Index: 22, Key: qbx88, Name: "Type 88 Sniper", Description: "A 10 round, semiautomatic rifle with a rugged design produces average accuracy and stopping power.", Image: "qbu88.png"},
		{Index: 23, Key: "sv98", Name: "SV98 Snaiperskaya", Description: "A 10 round, bolt action sniper rifle capable of engaging targets up to 1,000 m.", Image: "sv98.png"},
		{Index: 24, Key: "svu", Name: "SVU Snaiperskaya Short", Description: "A silenced, 10 round, semi automatic sniper rifle capable of extremely accurate fire at extended ranges.", Image: "svu.png"},
		{Index: 25, Key: "gol", Name: "GOL Sniper Magnum", Description: "An extremely accurate, bolt action sniper rifle capable of neutralizing targets at long range. Capacity: 5 rounds", Image: "gol.png"},
		{Index: 26, Key: "vss", Name: "VSS Snaiperskaya Special", Description: "A silenced sniper rifle that fires subsonic rounds allowing for extreme stealth at the cost of accuracy and range.", Image: "vss.png"},
		{Index: 27, Key: "m95", Name: "M95 Sniper", Description: "A 5 round, bolt action, bullpup sniper rifle designed to fire the .50 caliber up to 2,000 meters.", Image: "m95.png"},
		{Index: 28, Key: "m9", Name: "M9 Pistol", Description: "A 12 round semi-automatic pistol suited for close quarters combat due to its substantial mobility and short range.", Image: "m9beretta.png", IsPistol: true},
		{Index: 29, Key: "mcs", Name: "870 Combat", Description: "A pump action shotgun offering great mobility, high-power but short range. It comes with a 4 round tubular magazine.", Image: "870mcs.png"},
		{Index: 30, Key: "s12k", Name: "Saiga 20k Semi", Description: "A semi automatic, 6 round magazine fed shotgun ideal in tight spaces. It has exceptional stopping power but limited range.", Image: "saiga12.png"},
		{Index: 31, Key: "mp443", Name: "MP-443 Grach", Description: "17 round semi-automatic pistol is known for high mobility making it an ideal selection for close quarters combat.", Image: "mp443grach.png", IsPistol: true},
		{Index: 32, Key: "m1911", Name: "WWII M1911 .45", Description: "A favorite sidearm among special forces this 7 round semi-automatic pistol boasts a .45 caliber stopping power enabling the operator to dispose of an opponent with lethal close range fire.", Image: "m1911colt45.png", IsPistol: true},
		{Index: 33, Key: "m1a1", Name: "WWII M1A1 Thompson", Description: "30 round fully automatic sub machine gun known for high mobility and capability of firing 600 rounds per minute.", Image: "m1a1thompson.png"},
		{Index: 34, Key: "mp412", Name: "MP-412 Rex", Description: "A 6 round revolver revered for its close range edge and notable Magnum firepower. Its top break design is a throw back to old 6-shooters of the \"Wild West\".", Image: "mp412rex.png", IsPistol: true},
		{Index: 35, Key: "m93r", Name: "M93R Burst", Description: "Mobile 3 round burst pistol capable of firing 20 rounds allowing it to inflict more damage during close quarters combat.", Image: "m93r.png", IsPistol: true},
		{Index: 36, Key: "spas12", Name: "SPAS-12 Combat", Description: "A unique looking 4 round, pump action shotgun suited for close quarters combat. It has excellent fire power but limited range.", Image: "spas12.png"},
		{Index: 37, Key: "mk14ebr", Name: "M14", Description: "20 round semi automatic battle rifle compatible with telescoping stock making it ideal for long range combat.", Image: "m14.png"},
		{Index: 38, Key: "g3", Name: "G3", Description: "20 round fully automatic battle rifle designed to fire 600 rounds per minute. It's capable of delivering high power at the expense of enormous recoil.", Image: "g3.png"},
		{Index: 39, Key: "u12", Name: "USAS-12 Auto", Description: "A fully automatic shotgun with limited range and power. It comes with a high capacity 7 round magazine.", Image: "usas-12.png"},
		{Index: 40, Key: "m1", Name: "M1", Description: "Semi-automatic 8 round battle rifle maintaining sound accuracy coupled with considerable power negate a lack of mobility.", Image: "m1garand.png"},
		{Index: 41, Key: "n2k", Name: "NEOSTEAD 2000 Combat", Description: "A bullpup shotgun excellent for confined spaces.", Image: "neostead.png"},
		{Index: 42, Key: "m16k", Name: "M16A2 - SPECACT", Description: "Special Activities Unit. Unique camouflaged M16A2 gives you a visual edge on the Battlefield.", Image: "m16k.png"},
		{Index: 43, Key: "mg3k", Name: "MG3 - SPECACT", Description: "Special Activities Unit. Unique camouflaged MG3 gives you a visual edge on the Battlefield.", Image: "mg3k.png"},
		{Index: 44, Key: "m95k", Name: "M95 - SPECACT", Description: "Special Activities Unit. Unique camouflaged M95 gives you a visual edge on the Battlefield.", Image: "m95k.png"},
		{Index: 45, Key: "umpk", Name: "UMP - SPECACT", Description: "Special Activities Unit. Unique camouflaged UMP-45 gives you a visual edge on the Battlefield.", Image: "umpk.png"},
		// Currently unclear if a weapon is missing here (no original snapshot contains a weapon 46)
		{Index: 47, Key: "m2v", Name: "Flamethrower Vietnam", Description: "", Image: "vm2_flamethrower.png"},
		{Index: 48, Key: "m16a1v", Name: "M16A1 Vietnam", Description: "", Image: "vm16.png"},
		{Index: 49, Key: "ak47v", Name: "AK47 Vietnam", Description: "", Image: "vak47.png"},
		{Index: 50, Key: "m14v", Name: "M14 Vietnam", Description: "", Image: "vm14.png"},
		{Index: 51, Key: "mac10v", Name: "MAC10 Vietnam", Description: "", Image: "vmac10.png"},
		{Index: 52, Key: "ppshv", Name: "PPSh Vietnam", Description: "", Image: "vppsh.png"},
		{Index: 53, Key: "uziv", Name: "UZI Vietnam", Description: "", Image: "vuzi.png"},
		{Index: 54, Key: "m60v", Name: "M60 Vietnam", Description: "", Image: "vm60.png"},
		{Index: 55, Key: "rpkv", Name: "RPK Vietnam", Description: "", Image: "vrpk.png"},
		{Index: 56, Key: "xm22v", Name: "XM22 Vietnam", Description: "", Image: "xm22.png"},
		{Index: 57, Key: "m40v", Name: "M40 Vietnam", Description: "", Image: "vm40.png"},
		{Index: 58, Key: "svdv", Name: "SVD Vietnam", Description: "", Image: "vsvd.png"},
		{Index: 59, Key: "m21v", Name: "M21 Vietnam", Description: "", Image: "vm21.png"},
		{Index: 60, Key: "tt33v", Name: "TT33 Vietnam", Description: "", Image: "vtt-33.png"}, // Not tracked as a pistol on the original site
	}
}

type vehicle struct {
	Index       int
	Key         string
	Class       string
	Name        string
	Description template.HTML // Using template.HTML to render special characters correctly
	Image       string
}

func getVehicles() []vehicle {
	return []vehicle{
		{Index: 0, Key: "hmv", Class: "tveh", Name: "HMMWV 4WD", Description: "High Mobility Multipurpose Wheeled Vehicle with 4WD it seats 4 and is armed with a .50 caliber HMG in a 360 degree weapons ring mount.", Image: "humv.png"},
		{Index: 1, Key: "vodn", Class: "tveh", Name: "VODNIK 4WD", Description: "High Mobility Multipurpose Combat Vehicle with 4WD it seats 4 and is armed with a 12.7mm caliber HMG in a 360 degree turret.", Image: "vodnik.png"},
		{Index: 2, Key: "cobr", Class: "tveh", Name: "COBRA 4WD", Description: "4WD, Light Armored Vehicle with a crew of 4, armed with a .50 caliber HMG in a 360 degree weapon ring mount.", Image: "cobra.png"},
		{Index: 3, Key: "quad", Class: "tveh", Name: "Quad Bike", Description: "A 2 person Ultra-light Tactical Vehicle designed to operate over any terrain, in all weather and at any altitude.", Image: "quadbike.png"},
		{Index: 4, Key: "cav", Class: "tveh", Name: "Cav", Description: "", Image: "cav.png"},
		{Index: 5, Key: "tru", Class: "tveh", Name: "Truck", Description: "", Image: "truck.png"},
		{Index: 6, Key: "gaz69v", Class: "tveh", Name: "GAZ69 Vietnam", Description: "", Image: "gaz69.png"},
		{Index: 7, Key: "minitruckv", Class: "tveh", Name: "TUK-TUK Vietnam", Description: "", Image: "minitruck.png"},
		{Index: 8, Key: "m15v", Class: "tveh", Name: "M151 Vietnam", Description: "", Image: "m151.png"},
		{Index: 9, Key: "m1a2", Class: "arm", Name: "M1A2 Abrams", Description: "U.S. Main Battle Tank armed with a 120mm smoothbore cannon and remotely operated .50 caliber HMG. Crew: 2", Image: "m1a2.png"},
		{Index: 10, Key: "t90", Class: "arm", Name: "T-90 MBT", Description: "Russian Main Battle Tank armed with a 125mm smoothbore cannon and 12.7mm HMG. Crew: 2", Image: "t90.png"},
		{Index: 11, Key: "m3a3", Class: "arm", Name: "M3A3 Bradley", Description: "Infantry Fighting Vehicle armed with a 30mm chain gun, .50 caliber HMG and 2 passenger firing ports. Seats: 4", Image: "m3a3.png"},
		{Index: 12, Key: "bmd3", Class: "arm", Name: "BMD-3 Bakhcha", Description: "Infantry Fighting Vehicle armed with a 30mm auto cannon, 12.7mm HMG and 2 passenger firing ports.", Image: "bmda.png"},
		{Index: 13, Key: "bmda", Class: "arm", Name: "BMD-3 Bakhcha AA", Description: "Anti-aircraft variant of the BMD-3, armed with an AA-Gun, Grenade Launcher and 2 side mounted passenger firing ports. Seats: 4", Image: "bmd3aa.png"},
		// Currently unclear if a vehicle is missing here (no original snapshot contains a vehicle 14)
		{Index: 15, Key: "m48v", Class: "arm", Name: "M48 Tank Vietnam", Description: "", Image: "m48.png"},
		{Index: 16, Key: "t54v", Class: "arm", Name: "T54 Tank Vietnam", Description: "", Image: "t54.png"},
		{Index: 17, Key: "jets", Class: "sea", Name: "Personal Water Craft", Description: "Unarmed 2 person high speed Personal Water Craft (PWC).", Image: "jetski.png"},
		{Index: 18, Key: "PBLB", Class: "sea", Name: "Patrol Boat", Description: "Modern .50 caliber machine gun. Lightweight, highly mobile and offers little to no recoil. Exceptional Force Multiplier.", Image: "patrolboat.png"},
		{Index: 19, Key: "pbrv", Class: "sea", Name: "PBR Vietnam", Description: "", Image: "pbr.png"},
		{Index: 20, Key: "ah60", Class: "air", Name: "UH-60 Transport", Description: "Tactical Transport Helicopter manned by a crew of 2, capable of transporting 3 fully armed troops and armed with 2 x 7.62mm Miniguns.", Image: "ah60.png"},
		{Index: 21, Key: "ah64", Class: "air", Name: "AH-64 Apache", Description: "U.S. Attack Helicopter armed with 30mm chain gun and dual 70mm rocket pods. Crew: 2. Armor protection up to and including .50 caliber muntion.", Image: "ah64.png"},
		{Index: 22, Key: "MI28", Class: "air", Name: "MI-28 Havoc", Description: "Russian Attack Helicopter armed with 30mm chain gun and dual 80mm rocket pods. Crew: 2. Armor protection up to .50 caliber.", Image: "mi28havoc.png"},
		{Index: 23, Key: "havoc", Class: "air", Name: "Mi-24 Hind", Description: "Large Helicopter Gunship operated by a crew of 2 and capable of transporting 2 fully armed troops.", Image: "mi24hind.png"},
		{Index: 24, Key: "uav", Class: "air", Name: "UAV", Description: "UAV capable of lasing targets for air strikes or monitoring battlefield operations. It can be armed with either a .50 HMG or Smoke Dispenser.", Image: "uav.png"},
		{Index: 25, Key: "hueyv", Class: "air", Name: "HUEY Vietnam", Description: "", Image: "huey.png"},
		{Index: 26, Key: "XM312", Class: "stav", Name: "Heavy MG X312", Description: "Modern U.S Military .50 caliber Machine Gun procured to replace the 80-year-old Browning .50 Caliber.", Image: "x312.png"},
		{Index: 27, Key: "KORD", Class: "stav", Name: "Heavy MG KORD", Description: "Russian 12.7mm HMG providing less recoil and improved accuracy at longer ranges.", Image: "kord.png"},
		{Index: 28, Key: "KORN", Class: "stav", Name: "Stationary AT KORN", Description: "Russian made supersonic, Anti-Tank wire Guided Missile (ATGM) designed to destroy main battle tanks and engage low flying helicopters.", Image: "korn.png"},
		{Index: 29, Key: "TOW2", Class: "stav", Name: "Stationary AT TOW2", Description: "Tube Launched Optically tracked Wire guided Missile system with a range of 4,000 meters designed to engage armor and low-level slow moving aircraft.", Image: "tow2.png"},
		{Index: 30, Key: "aav", Class: "stav", Name: "Anti-Air Gun", Description: "Russian 23mm auto cannon capable of anti-air defense and light ground support.", Image: "zu23.png"},
		{Index: 31, Key: "VADS", Class: "stav", Name: "VADS", Description: "", Image: "vads.png"},
		{Index: 32, Key: "XM307", Class: "stav", Name: "XM307", Description: "", Image: "x307.png"},
		{Index: 33, Key: "QLZ8", Class: "stav", Name: "QLZ8", Description: "", Image: "qlz8.png"},
	}
}

func getServiceStarClass(values map[string]any, key string) string {
	if getFloat(values, fmt.Sprintf("pl%s_00", key)) > 0 {
		return "plat"
	}

	if value := getFloat(values, fmt.Sprintf("go%s_00", key)); value > 0 {
		return fmt.Sprintf("gold%.0f", value)
	}

	if getFloat(values, fmt.Sprintf("si%s_00", key)) > 0 {
		return "silv"
	}

	if getFloat(values, fmt.Sprintf("br%s_00", key)) > 0 {
		return "bron"
	}

	return "none"
}

func calculateProgress(current, threshold float64) float64 {
	if threshold == 0 {
		return 0
	}
	percentage := (current / threshold) * 100
	// Cap at 100 and round to 2 decimal places
	if percentage > 100 {
		percentage = 100
	}
	return math.Round(percentage*100) / 100
}
