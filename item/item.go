package item

import ()

type Item struct {
	Id                  int `db:"id"`                  //	int(11)	NO	PRI	0
	MinStatus           int `db:"minstatus"`           //	smallint(5)	NO	MUL	0
	Name                int `db:"Name"`                //	varchar(64)	NO	MUL
	Aagi                int `db:"aagi"`                //	int(11)	NO		0
	Ac                  int `db:"ac"`                  //	int(11)	NO		0
	Accuracy            int `db:"accuracy"`            //	int(11)	NO		0
	Acha                int `db:"acha"`                //	int(11)	NO		0
	Adex                int `db:"adex"`                //	int(11)	NO		0
	Aint                int `db:"aint"`                //	int(11)	NO		0
	ArtifactFlag        int `db:"artifactflag"`        //	tinyint(3) unsigned	NO		0
	Asta                int `db:"asta"`                //	int(11)	NO		0
	Astr                int `db:"astr"`                //	int(11)	NO		0
	Attack              int `db:"attack"`              //	int(11)	NO		0
	AugRestrict         int `db:"augrestrict"`         //	int(11)	NO		0
	AugSlot1Type        int `db:"augslot1type"`        //	tinyint(3)	NO		0
	AugSlot1Visible     int `db:"augslot1visible"`     //	tinyint(3)	YES		NULL
	AugSlot2Type        int `db:"augslot2type"`        //	tinyint(3)	NO		0
	AugSlot2Visible     int `db:"augslot2visible"`     //	tinyint(3)	YES		NULL
	AugSlot3Type        int `db:"augslot3type"`        //	tinyint(3)	NO		0
	AugSlot3Visible     int `db:"augslot3visible"`     //	tinyint(3)	YES		NULL
	AugSlot4Type        int `db:"augslot4type"`        //	tinyint(3)	NO		0
	AugSlot4Visible     int `db:"augslot4visible"`     //	tinyint(3)	YES		NULL
	AugSlot5Type        int `db:"augslot5type"`        //	tinyint(3)	NO		0
	AugSlot5Visible     int `db:"augslot5visible"`     //	tinyint(3)	YES		NULL
	AugType             int `db:"augtype"`             //	int(11)	NO		0
	Avoidance           int `db:"avoidance"`           //	int(11)	NO		0
	Awis                int `db:"awis"`                //	int(11)	NO		0
	BagSize             int `db:"bagsize"`             //	int(11)	NO		0
	BagSlots            int `db:"bagslots"`            //	int(11)	NO		0
	BagType             int `db:"bagtype"`             //	int(11)	NO		0
	BagWR               int `db:"bagwr"`               //	int(11)	NO		0
	BaneDmgAmt          int `db:"banedmgamt"`          //	int(11)	NO		0
	BaneDmgRaceAmt      int `db:"banedmgraceamt"`      //	int(11)	NO		0
	BaneDmgBody         int `db:"banedmgbody"`         //	int(11)	NO		0
	BaneDmgRace         int `db:"banedmgrace"`         //	int(11)	NO		0
	BardType            int `db:"bardtype"`            //	int(11)	NO		0
	BardValue           int `db:"bardvalue"`           //	int(11)	NO		0
	Book                int `db:"book"`                //	int(11)	NO		0
	CastTime            int `db:"casttime"`            //	int(11)	NO		0
	CastTime_           int `db:"casttime_"`           //	int(11)	NO		0
	CharmFile           int `db:"charmfile"`           //	varchar(32)	NO
	CharmFileID         int `db:"charmfileid"`         //	varchar(32)	NO
	Classes             int `db:"classes"`             //	int(11)	NO		0
	Color               int `db:"color"`               //	int(10) unsigned	NO		0
	CombatEffects       int `db:"combateffects"`       //	varchar(10)	NO
	ExtraDmgSkill       int `db:"extradmgskill"`       //	int(11)	NO		0
	ExtraDmgAmt         int `db:"extradmgamt"`         //	int(11)	NO		0
	Price               int `db:"price"`               //	int(11)	NO		0
	Cr                  int `db:"cr"`                  //	int(11)	NO		0
	Damage              int `db:"damage"`              //	int(11)	NO		0
	DamageShield        int `db:"damageshield"`        //	int(11)	NO		0
	Deity               int `db:"deity"`               //	int(11)	NO		0
	Delay               int `db:"delay"`               //	int(11)	NO		0
	AugDistiller        int `db:"augdistiller"`        //	int(11)	NO		0
	DotShielding        int `db:"dotshielding"`        //	int(11)	NO		0
	Dr                  int `db:"dr"`                  //	int(11)	NO		0
	ClickType           int `db:"clicktype"`           //	int(11)	NO		0
	ClickLevel2         int `db:"clicklevel2"`         //	int(11)	NO		0
	ElemDmgType         int `db:"elemdmgtype"`         //	int(11)	NO		0
	ElemDmgAmt          int `db:"elemdmgamt"`          //	int(11)	NO		0
	Endur               int `db:"endur"`               //	int(11)	NO		0
	FactionAmt1         int `db:"factionamt1"`         //	int(11)	NO		0
	FactionAmt2         int `db:"factionamt2"`         //	int(11)	NO		0
	FactionAmt3         int `db:"factionamt3"`         //	int(11)	NO		0
	FactionAmt4         int `db:"factionamt4"`         //	int(11)	NO		0
	FactionMod1         int `db:"factionmod1"`         //	int(11)	NO		0
	FactionMod2         int `db:"factionmod2"`         //	int(11)	NO		0
	FactionMod3         int `db:"factionmod3"`         //	int(11)	NO		0
	FactionMod4         int `db:"factionmod4"`         //	int(11)	NO		0
	Filename            int `db:"filename"`            //	varchar(32)	NO
	FocusEffect         int `db:"focuseffect"`         //	int(11)	NO		0
	Fr                  int `db:"fr"`                  //	int(11)	NO		0
	FvNoDrop            int `db:"fvnodrop"`            //	int(11)	NO		0
	Haste               int `db:"haste"`               //	int(11)	NO		0
	ClickLevel          int `db:"clicklevel"`          //	int(11)	NO		0
	Hp                  int `db:"hp"`                  //	int(11)	NO		0
	Regen               int `db:"regen"`               //	int(11)	NO		0
	Icon                int `db:"icon"`                //	int(11)	NO		0
	IdFile              int `db:"idfile"`              //	varchar(30)	NO
	ItemClass           int `db:"itemclass"`           //	int(11)	NO		0
	ItemType            int `db:"itemtype"`            //	int(11)	NO		0
	LdonPrice           int `db:"ldonprice"`           //	int(11)	NO		0
	LdonTheme           int `db:"ldontheme"`           //	int(11)	NO		0
	LdonSold            int `db:"ldonsold"`            //	int(11)	NO		0
	Light               int `db:"light"`               //	int(11)	NO		0
	Lore                int `db:"lore"`                //	varchar(80)	NO	MUL
	LoreGroup           int `db:"loregroup"`           //	int(11)	NO		0
	Magic               int `db:"magic"`               //	int(11)	NO		0
	Mana                int `db:"mana"`                //	int(11)	NO		0
	ManaRegen           int `db:"manaregen"`           //	int(11)	NO		0
	EnduranceRegen      int `db:"enduranceregen"`      //	int(11)	NO		0
	Material            int `db:"material"`            //	int(11)	NO		0
	MaxCharges          int `db:"maxcharges"`          //	int(11)	NO		0
	Mr                  int `db:"mr"`                  //	int(11)	NO		0
	NoDrop              int `db:"nodrop"`              //	int(11)	NO		0
	NoRent              int `db:"norent"`              //	int(11)	NO		0
	PendingLoreFlag     int `db:"pendingloreflag"`     //	tinyint(3) unsigned	NO		0
	Pr                  int `db:"pr"`                  //	int(11)	NO		0
	Procrate            int `db:"procrate"`            //	int(11)	NO		0
	Races               int `db:"races"`               //	int(11)	NO		0
	Range               int `db:"range"`               //	int(11)	NO		0
	RecLevel            int `db:"reclevel"`            //	int(11)	NO		0
	RecSkill            int `db:"recskill"`            //	int(11)	NO		0
	ReqLevel            int `db:"reqlevel"`            //	int(11)	NO		0
	SellRate            int `db:"sellrate"`            //	float	NO		0
	Shielding           int `db:"shielding"`           //	int(11)	NO		0
	Size                int `db:"size"`                //	int(11)	NO		0
	SkillModType        int `db:"skillmodtype"`        //	int(11)	NO		0
	SkillModValue       int `db:"skillmodvalue"`       //	int(11)	NO		0
	Slots               int `db:"slots"`               //	int(11)	NO		0
	ClickEffect         int `db:"clickeffect"`         //	int(11)	NO		0
	SpellShield         int `db:"spellshield"`         //	int(11)	NO		0
	StrikeThrough       int `db:"strikethrough"`       //	int(11)	NO		0
	StunResist          int `db:"stunresist"`          //	int(11)	NO		0
	SummonedFlag        int `db:"summonedflag"`        //	tinyint(3) unsigned	NO		0
	TradeSkills         int `db:"tradeskills"`         //	int(11)	NO		0
	Favor               int `db:"favor"`               //	int(11)	NO		0
	Weight              int `db:"weight"`              //	int(11)	NO		0
	Unk012              int `db:"UNK012"`              //	int(11)	NO		0
	Unk013              int `db:"UNK013"`              //	int(11)	NO		0
	BenefitFlag         int `db:"benefitflag"`         //	int(11)	NO		0
	Unk054              int `db:"UNK054"`              //	int(11)	NO		0
	Unk059              int `db:"UNK059"`              //	int(11)	NO		0
	BookType            int `db:"booktype"`            //	int(11)	NO		0
	RecastDelay         int `db:"recastdelay"`         //	int(11)	NO		0
	RecastType          int `db:"recasttype"`          //	int(11)	NO		0
	GuildFavor          int `db:"guildfavor"`          //	int(11)	NO		0
	Unk123              int `db:"UNK123"`              //	int(11)	NO		0
	Unk124              int `db:"UNK124"`              //	int(11)	NO		0
	Attuneable          int `db:"attuneable"`          //	int(11)	NO		0
	Nopet               int `db:"nopet"`               //	int(11)	NO		0
	UpdateD             int `db:"updated"`             //	datetime	NO		0000-00-00 00:00:00
	Comment             int `db:"comment"`             //	varchar(255)	NO
	Unk127              int `db:"UNK127"`              //	int(11)	NO		0
	PointType           int `db:"pointtype"`           //	int(11)	NO		0
	PotionBelt          int `db:"potionbelt"`          //	int(11)	NO		0
	PotionBeltSlots     int `db:"potionbeltslots"`     //	int(11)	NO		0
	StackSize           int `db:"stacksize"`           //	int(11)	NO		0
	NoTransfer          int `db:"notransfer"`          //	int(11)	NO		0
	Stackable           int `db:"stackable"`           //	int(11)	NO		0
	Unk134              int `db:"UNK134"`              //	varchar(255)	NO
	Unk137              int `db:"UNK137"`              //	int(11)	NO		0
	ProcEffect          int `db:"proceffect"`          //	int(11)	NO		0
	ProcType            int `db:"proctype"`            //	int(11)	NO		0
	ProcLevel2          int `db:"proclevel2"`          //	int(11)	NO		0
	ProcLevel           int `db:"proclevel"`           //	int(11)	NO		0
	Unk142              int `db:"UNK142"`              //	int(11)	NO		0
	WornEffect          int `db:"worneffect"`          //	int(11)	NO		0
	WornType            int `db:"worntype"`            //	int(11)	NO		0
	WornLevel2          int `db:"wornlevel2"`          //	int(11)	NO		0
	WornLevel           int `db:"wornlevel"`           //	int(11)	NO		0
	Unk147              int `db:"UNK147"`              //	int(11)	NO		0
	FocusType           int `db:"focustype"`           //	int(11)	NO		0
	FocusLevel2         int `db:"focuslevel2"`         //	int(11)	NO		0
	FocusLevel          int `db:"focuslevel"`          //	int(11)	NO		0
	Unk152              int `db:"UNK152"`              //	int(11)	NO		0
	ScrollEffect        int `db:"scrolleffect"`        //	int(11)	NO		0
	ScrollType          int `db:"scrolltype"`          //	int(11)	NO		0
	ScrollLevel2        int `db:"scrolllevel2"`        //	int(11)	NO		0
	ScrollLevel         int `db:"scrolllevel"`         //	int(11)	NO		0
	Unk157              int `db:"UNK157"`              //	int(11)	NO		0
	Serialized          int `db:"serialized"`          //	datetime	YES		NULL
	Verified            int `db:"verified"`            //	datetime	YES		NULL
	Serialization       int `db:"serialization"`       //	text	YES		NULL
	Source              int `db:"source"`              //	varchar(20)	NO
	Unk033              int `db:"UNK033"`              //	int(11)	NO		0
	LoreFile            int `db:"lorefile"`            //	varchar(32)	NO
	Unk014              int `db:"UNK014"`              //	int(11)	NO		0
	SvCorruption        int `db:"svcorruption"`        //	int(11)	NO		0
	Unk038              int `db:"UNK038"`              //	int(11)	NO		0
	Unk060              int `db:"UNK060"`              //	int(11)	NO		0
	AugSlot1Unk2        int `db:"augslot1unk2"`        //	int(11)	NO		0
	AugSlot2Unk2        int `db:"augslot2unk2"`        //	int(11)	NO		0
	AugSlot3Unk2        int `db:"augslot3unk2"`        //	int(11)	NO		0
	AugSlot4Unk2        int `db:"augslot4unk2"`        //	int(11)	NO		0
	AugSlot5Unk2        int `db:"augslot5unk2"`        //	int(11)	NO		0
	Unk120              int `db:"UNK120"`              //	int(11)	NO		0
	Unk121              int `db:"UNK121"`              //	int(11)	NO		0
	QuestItemFlag       int `db:"questitemflag"`       //	int(11)	NO		0
	Unk132              int `db:"UNK132"`              //	text	NO		NULL
	ClickUnk5           int `db:"clickunk5"`           //	int(11)	NO		0
	ClickUnk6           int `db:"clickunk6"`           //	varchar(32)	NO
	ClickUnk7           int `db:"clickunk7"`           //	int(11)	NO		0
	ProcUnk1            int `db:"procunk1"`            //	int(11)	NO		0
	ProcUnk2            int `db:"procunk2"`            //	int(11)	NO		0
	ProcUnk3            int `db:"procunk3"`            //	int(11)	NO		0
	ProcUnk4            int `db:"procunk4"`            //	int(11)	NO		0
	ProcUnk6            int `db:"procunk6"`            //	varchar(32)	NO
	ProcUnk7            int `db:"procunk7"`            //	int(11)	NO		0
	WornUnk1            int `db:"wornunk1"`            //	int(11)	NO		0
	WornUnk2            int `db:"wornunk2"`            //	int(11)	NO		0
	WornUnk3            int `db:"wornunk3"`            //	int(11)	NO		0
	WornUnk4            int `db:"wornunk4"`            //	int(11)	NO		0
	WornUnk5            int `db:"wornunk5"`            //	int(11)	NO		0
	WornUnk6            int `db:"wornunk6"`            //	varchar(32)	NO
	WornUnk7            int `db:"wornunk7"`            //	int(11)	NO		0
	FocusUnk1           int `db:"focusunk1"`           //	int(11)	NO		0
	FocusUnk2           int `db:"focusunk2"`           //	int(11)	NO		0
	FocusUnk3           int `db:"focusunk3"`           //	int(11)	NO		0
	FocusUnk4           int `db:"focusunk4"`           //	int(11)	NO		0
	FocusUnk5           int `db:"focusunk5"`           //	int(11)	NO		0
	FocusUnk6           int `db:"focusunk6"`           //	varchar(32)	NO
	FocusUnk7           int `db:"focusunk7"`           //	int(11)	NO		0
	ScrollUnk1          int `db:"scrollunk1"`          //	int(11)	NO		0
	ScrollUnk2          int `db:"scrollunk2"`          //	int(11)	NO		0
	ScrollUnk3          int `db:"scrollunk3"`          //	int(11)	NO		0
	ScrollUnk4          int `db:"scrollunk4"`          //	int(11)	NO		0
	ScrollUnk5          int `db:"scrollunk5"`          //	int(11)	NO		0
	ScrollUnk6          int `db:"scrollunk6"`          //	varchar(32)	NO
	ScrollUnk7          int `db:"scrollunk7"`          //	int(11)	NO		0
	Unk193              int `db:"UNK193"`              //	int(11)	NO		0
	Purity              int `db:"purity"`              //	int(11)	NO		0
	EvolvingLevel       int `db:"evolvinglevel"`       //	int(11)	NO		0
	ClickName           int `db:"clickname"`           //	varchar(64)	NO
	ProcName            int `db:"procname"`            //	varchar(64)	NO
	WornName            int `db:"wornname"`            //	varchar(64)	NO
	FocusName           int `db:"focusname"`           //	varchar(64)	NO
	ScrollName          int `db:"scrollname"`          //	varchar(64)	NO
	DsMitigation        int `db:"dsmitigation"`        //	smallint(6)	NO		0
	HeroicStr           int `db:"heroic_str"`          //	smallint(6)	NO		0
	HeroicInt           int `db:"heroic_int"`          //	smallint(6)	NO		0
	HeroicWis           int `db:"heroic_wis"`          //	smallint(6)	NO		0
	HeroicAgi           int `db:"heroic_agi"`          //	smallint(6)	NO		0
	HeroicDex           int `db:"heroic_dex"`          //	smallint(6)	NO		0
	HeroicSta           int `db:"heroic_sta"`          //	smallint(6)	NO		0
	HeroicCha           int `db:"heroic_cha"`          //	smallint(6)	NO		0
	HeroicPr            int `db:"heroic_pr"`           //	smallint(6)	NO		0
	HeroicDr            int `db:"heroic_dr"`           //	smallint(6)	NO		0
	HeroicFr            int `db:"heroic_fr"`           //	smallint(6)	NO		0
	HeroicCr            int `db:"heroic_cr"`           //	smallint(6)	NO		0
	HeroicMr            int `db:"heroic_mr"`           //	smallint(6)	NO		0
	HeroicSvCorrup      int `db:"heroic_svcorrup"`     //	smallint(6)	NO		0
	HealAmt             int `db:"healamt"`             //	smallint(6)	NO		0
	Spelldmg            int `db:"spelldmg"`            //	smallint(6)	NO		0
	Clairvoyance        int `db:"clairvoyance"`        //	smallint(6)	NO		0
	BackstabDmg         int `db:"backstabdmg"`         //	smallint(6)	NO		0
	Created             int `db:"created"`             //	varchar(64)	NO
	EliteMaterial       int `db:"elitematerial"`       //	smallint(6)	NO		0
	LdonSellBackRate    int `db:"ldonsellbackrate"`    //	smallint(6)	NO		0
	ScriptFileID        int `db:"scriptfileid"`        //	smallint(6)	NO		0
	ExpendableArrow     int `db:"expendablearrow"`     //	smallint(6)	NO		0
	PowerSourceCapacity int `db:"powersourcecapacity"` //	smallint(6)	NO		0
	BardEffect          int `db:"bardeffect"`          //	smallint(6)	NO		0
	BardEffectType      int `db:"bardeffecttype"`      //	smallint(6)	NO		0
	BardLevel2          int `db:"bardlevel2"`          //	smallint(6)	NO		0
	BardLevel           int `db:"bardlevel"`           //	smallint(6)	NO		0
	BardUnk1            int `db:"bardunk1"`            //	smallint(6)	NO		0
	BardUnk2            int `db:"bardunk2"`            //	smallint(6)	NO		0
	BardUnk3            int `db:"bardunk3"`            //	smallint(6)	NO		0
	BardUnk4            int `db:"bardunk4"`            //	smallint(6)	NO		0
	BardUnk5            int `db:"bardunk5"`            //	smallint(6)	NO		0
	BardName            int `db:"bardname"`            //	varchar(64)	NO
	BardUnk7            int `db:"bardunk7"`            //	smallint(6)	NO		0
	Unk214              int `db:"UNK214"`              //	smallint(6)	NO		0
	Unk219              int `db:"UNK219"`              //	int(11)	NO		0
	Unk220              int `db:"UNK220"`              //	int(11)	NO		0
	Unk221              int `db:"UNK221"`              //	int(11)	NO		0
	Unk222              int `db:"UNK222"`              //	int(11)	NO		0
	Unk223              int `db:"UNK223"`              //	int(11)	NO		0
	Unk224              int `db:"UNK224"`              //	int(11)	NO		0
	Unk225              int `db:"UNK225"`              //	int(11)	NO		0
	Unk226              int `db:"UNK226"`              //	int(11)	NO		0
	Unk227              int `db:"UNK227"`              //	int(11)	NO		0
	Unk228              int `db:"UNK228"`              //	int(11)	NO		0
	Unk229              int `db:"UNK229"`              //	int(11)	NO		0
	Unk230              int `db:"UNK230"`              //	int(11)	NO		0
	Unk231              int `db:"UNK231"`              //	int(11)	NO		0
	Unk232              int `db:"UNK232"`              //	int(11)	NO		0
	Unk233              int `db:"UNK233"`              //	int(11)	NO		0
	Unk234              int `db:"UNK234"`              //	int(11)	NO		0
	Unk235              int `db:"UNK235"`              //	int(11)	NO		0
	Unk236              int `db:"UNK236"`              //	int(11)	NO		0
	Unk237              int `db:"UNK237"`              //	int(11)	NO		0
	Unk238              int `db:"UNK238"`              //	int(11)	NO		0
	Unk239              int `db:"UNK239"`              //	int(11)	NO		0
	Unk240              int `db:"UNK240"`              //	int(11)	NO		0
	Unk241              int `db:"UNK241"`              //	int(11)	NO		0
	Unk242              int `db:"UNK242"`              //	int(11)	NO		0
}
