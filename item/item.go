package item

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Item struct {
	Id                  int            `db:"id"`              //	int(11)	NO	PRI	0
	Minstatus           int            `db:"minstatus"`       //	smallint(5)	NO	MUL	0
	Name                string         `db:"Name"`            //	varchar(64)	NO	MUL
	Aagi                int            `db:"aagi"`            //	int(11)	NO		0
	Ac                  int            `db:"ac"`              //	int(11)	NO		0
	Accuracy            int            `db:"accuracy"`        //	int(11)	NO		0
	Acha                int            `db:"acha"`            //	int(11)	NO		0
	Adex                int            `db:"adex"`            //	int(11)	NO		0
	Aint                int            `db:"aint"`            //	int(11)	NO		0
	Artifactflag        int            `db:"artifactflag"`    //	tinyint(3) unsigned	NO		0
	Asta                int            `db:"asta"`            //	int(11)	NO		0
	Astr                int            `db:"astr"`            //	int(11)	NO		0
	Attack              int            `db:"attack"`          //	int(11)	NO		0
	Augrestrict         int            `db:"augrestrict"`     //	int(11)	NO		0
	Augslot1type        int            `db:"augslot1type"`    //	tinyint(3)	NO		0
	Augslot1visible     sql.NullInt64  `db:"augslot1visible"` //	tinyint(3)	YES		NULL
	Augslot2type        int            `db:"augslot2type"`    //	tinyint(3)	NO		0
	Augslot2visible     sql.NullInt64  `db:"augslot2visible"` //	tinyint(3)	YES		NULL
	Augslot3type        int            `db:"augslot3type"`    //	tinyint(3)	NO		0
	Augslot3visible     sql.NullInt64  `db:"augslot3visible"` //	tinyint(3)	YES		NULL
	Augslot4type        int            `db:"augslot4type"`    //	tinyint(3)	NO		0
	Augslot4visible     sql.NullInt64  `db:"augslot4visible"` //	tinyint(3)	YES		NULL
	Augslot5type        int            `db:"augslot5type"`    //	tinyint(3)	NO		0
	Augslot5visible     sql.NullInt64  `db:"augslot5visible"` //	tinyint(3)	YES		NULL
	Augslot6type        int            `db:"augslot6type"`    //	tinyint(3)	NO		0
	Augslot6visible     sql.NullInt64  `db:"augslot6visible"` //	tinyint(3)	YES		NULL
	Augtype             int            `db:"augtype"`         //	int(11)	NO		0
	Avoidance           int            `db:"avoidance"`       //	int(11)	NO		0
	Awis                int            `db:"awis"`            //	int(11)	NO		0
	Bagsize             int            `db:"bagsize"`         //	int(11)	NO		0
	Bagslots            int            `db:"bagslots"`        //	int(11)	NO		0
	Bagtype             int            `db:"bagtype"`         //	int(11)	NO		0
	Bagwr               int            `db:"bagwr"`           //	int(11)	NO		0
	Banedmgamt          int            `db:"banedmgamt"`      //	int(11)	NO		0
	Banedmgraceamt      int            `db:"banedmgraceamt"`  //	int(11)	NO		0
	Banedmgbody         int            `db:"banedmgbody"`     //	int(11)	NO		0
	Banedmgrace         int            `db:"banedmgrace"`     //	int(11)	NO		0
	Bardtype            int            `db:"bardtype"`        //	int(11)	NO		0
	Bardvalue           int            `db:"bardvalue"`       //	int(11)	NO		0
	Book                int            `db:"book"`            //	int(11)	NO		0
	Casttime            int            `db:"casttime"`        //	int(11)	NO		0
	Casttime_           int            `db:"casttime_"`       //	int(11)	NO		0
	Charmfile           string         `db:"charmfile"`       //	varchar(32)	NO
	Charmfileid         string         `db:"charmfileid"`     //	varchar(32)	NO
	Classes             int            `db:"classes"`         //	int(11)	NO		0
	Color               int            `db:"color"`           //	int(10) unsigned	NO		0
	Combateffects       string         `db:"combateffects"`   //	varchar(10)	NO
	Extradmgskill       int            `db:"extradmgskill"`   //	int(11)	NO		0
	Extradmgamt         int            `db:"extradmgamt"`     //	int(11)	NO		0
	Price               int            `db:"price"`           //	int(11)	NO		0
	Cr                  int            `db:"cr"`              //	int(11)	NO		0
	Damage              int            `db:"damage"`          //	int(11)	NO		0
	Damageshield        int            `db:"damageshield"`    //	int(11)	NO		0
	Deity               int            `db:"deity"`           //	int(11)	NO		0
	Delay               int            `db:"delay"`           //	int(11)	NO		0
	Augdistiller        int            `db:"augdistiller"`    //	int(11)	NO		0
	Dotshielding        int            `db:"dotshielding"`    //	int(11)	NO		0
	Dr                  int            `db:"dr"`              //	int(11)	NO		0
	Clicktype           int            `db:"clicktype"`       //	int(11)	NO		0
	Clicklevel2         int            `db:"clicklevel2"`     //	int(11)	NO		0
	Elemdmgtype         int            `db:"elemdmgtype"`     //	int(11)	NO		0
	Elemdmgamt          int            `db:"elemdmgamt"`      //	int(11)	NO		0
	Endur               int            `db:"endur"`           //	int(11)	NO		0
	Factionamt1         int            `db:"factionamt1"`     //	int(11)	NO		0
	Factionamt2         int            `db:"factionamt2"`     //	int(11)	NO		0
	Factionamt3         int            `db:"factionamt3"`     //	int(11)	NO		0
	Factionamt4         int            `db:"factionamt4"`     //	int(11)	NO		0
	Factionmod1         int            `db:"factionmod1"`     //	int(11)	NO		0
	Factionmod2         int            `db:"factionmod2"`     //	int(11)	NO		0
	Factionmod3         int            `db:"factionmod3"`     //	int(11)	NO		0
	Factionmod4         int            `db:"factionmod4"`     //	int(11)	NO		0
	Filename            string         `db:"filename"`        //	varchar(32)	NO
	Focuseffect         int            `db:"focuseffect"`     //	int(11)	NO		0
	Fr                  int            `db:"fr"`              //	int(11)	NO		0
	Fvnodrop            int            `db:"fvnodrop"`        //	int(11)	NO		0
	Haste               int            `db:"haste"`           //	int(11)	NO		0
	Clicklevel          int            `db:"clicklevel"`      //	int(11)	NO		0
	Hp                  int            `db:"hp"`              //	int(11)	NO		0
	Regen               int            `db:"regen"`           //	int(11)	NO		0
	Icon                int            `db:"icon"`            //	int(11)	NO		0
	Idfile              string         `db:"idfile"`          //	varchar(30)	NO
	Itemclass           int            `db:"itemclass"`       //	int(11)	NO		0
	Itemtype            int            `db:"itemtype"`        //	int(11)	NO		0
	Ldonprice           int            `db:"ldonprice"`       //	int(11)	NO		0
	Ldontheme           int            `db:"ldontheme"`       //	int(11)	NO		0
	Ldonsold            int            `db:"ldonsold"`        //	int(11)	NO		0
	Light               int            `db:"light"`           //	int(11)	NO		0
	Lore                string         `db:"lore"`            //	varchar(80)	NO	MUL
	Loregroup           int            `db:"loregroup"`       //	int(11)	NO		0
	Magic               int            `db:"magic"`           //	int(11)	NO		0
	Mana                int            `db:"mana"`            //	int(11)	NO		0
	Manaregen           int            `db:"manaregen"`       //	int(11)	NO		0
	Enduranceregen      int            `db:"enduranceregen"`  //	int(11)	NO		0
	Material            int            `db:"material"`        //	int(11)	NO		0
	Herosforgemodel     int            `db:"herosforgemodel"`
	Maxcharges          int            `db:"maxcharges"`          //	int(11)	NO		0
	Mr                  int            `db:"mr"`                  //	int(11)	NO		0
	Nodrop              int            `db:"nodrop"`              //	int(11)	NO		0
	Norent              int            `db:"norent"`              //	int(11)	NO		0
	Pendingloreflag     int            `db:"pendingloreflag"`     //	tinyint(3) unsigned	NO		0
	Pr                  int            `db:"pr"`                  //	int(11)	NO		0
	Procrate            int            `db:"procrate"`            //	int(11)	NO		0
	Races               int            `db:"races"`               //	int(11)	NO		0
	Range               int            `db:"range"`               //	int(11)	NO		0
	Reclevel            int            `db:"reclevel"`            //	int(11)	NO		0
	Recskill            int            `db:"recskill"`            //	int(11)	NO		0
	Reqlevel            int            `db:"reqlevel"`            //	int(11)	NO		0
	Sellrate            float32        `db:"sellrate"`            //	float	NO		0
	Shielding           int            `db:"shielding"`           //	int(11)	NO		0
	Size                int            `db:"size"`                //	int(11)	NO		0
	Skillmodtype        int            `db:"skillmodtype"`        //	int(11)	NO		0
	Skillmodvalue       int            `db:"skillmodvalue"`       //	int(11)	NO		0
	Slots               int            `db:"slots"`               //	int(11)	NO		0
	Clickeffect         int            `db:"clickeffect"`         //	int(11)	NO		0
	Spellshield         int            `db:"spellshield"`         //	int(11)	NO		0
	Strikethrough       int            `db:"strikethrough"`       //	int(11)	NO		0
	Stunresist          int            `db:"stunresist"`          //	int(11)	NO		0
	Summonedflag        int            `db:"summonedflag"`        //	tinyint(3) unsigned	NO		0
	Tradeskills         int            `db:"tradeskills"`         //	int(11)	NO		0
	Favor               int            `db:"favor"`               //	int(11)	NO		0
	Weight              int            `db:"weight"`              //	int(11)	NO		0
	Unk012              int            `db:"UNK012"`              //	int(11)	NO		0
	Unk013              int            `db:"UNK013"`              //	int(11)	NO		0
	Benefitflag         int            `db:"benefitflag"`         //	int(11)	NO		0
	Unk054              int            `db:"UNK054"`              //	int(11)	NO		0
	Unk059              int            `db:"UNK059"`              //	int(11)	NO		0
	Booktype            int            `db:"booktype"`            //	int(11)	NO		0
	Recastdelay         int            `db:"recastdelay"`         //	int(11)	NO		0
	Recasttype          int            `db:"recasttype"`          //	int(11)	NO		0
	Guildfavor          int            `db:"guildfavor"`          //	int(11)	NO		0
	Unk123              int            `db:"UNK123"`              //	int(11)	NO		0
	Unk124              int            `db:"UNK124"`              //	int(11)	NO		0
	Attuneable          int            `db:"attuneable"`          //	int(11)	NO		0
	Nopet               int            `db:"nopet"`               //	int(11)	NO		0
	Updated             time.Time      `db:"updated"`             //	datetime	NO		0000-00-00 00:00:00
	Comment             string         `db:"comment"`             //	varchar(255)	NO
	Unk127              int            `db:"UNK127"`              //	int(11)	NO		0
	Pointtype           int            `db:"pointtype"`           //	int(11)	NO		0
	Potionbelt          int            `db:"potionbelt"`          //	int(11)	NO		0
	Potionbeltslots     int            `db:"potionbeltslots"`     //	int(11)	NO		0
	Stacksize           int            `db:"stacksize"`           //	int(11)	NO		0
	Notransfer          int            `db:"notransfer"`          //	int(11)	NO		0
	Stackable           int            `db:"stackable"`           //	int(11)	NO		0
	Unk134              string         `db:"UNK134"`              //	varchar(255)	NO
	Unk137              int            `db:"UNK137"`              //	int(11)	NO		0
	Proceffect          int            `db:"proceffect"`          //	int(11)	NO		0
	Proctype            int            `db:"proctype"`            //	int(11)	NO		0
	Proclevel2          int            `db:"proclevel2"`          //	int(11)	NO		0
	Proclevel           int            `db:"proclevel"`           //	int(11)	NO		0
	Unk142              int            `db:"UNK142"`              //	int(11)	NO		0
	Worneffect          int            `db:"worneffect"`          //	int(11)	NO		0
	Worntype            int            `db:"worntype"`            //	int(11)	NO		0
	Wornlevel2          int            `db:"wornlevel2"`          //	int(11)	NO		0
	Wornlevel           int            `db:"wornlevel"`           //	int(11)	NO		0
	Unk147              int            `db:"UNK147"`              //	int(11)	NO		0
	Focustype           int            `db:"focustype"`           //	int(11)	NO		0
	Focuslevel2         int            `db:"focuslevel2"`         //	int(11)	NO		0
	Focuslevel          int            `db:"focuslevel"`          //	int(11)	NO		0
	Unk152              int            `db:"UNK152"`              //	int(11)	NO		0
	Scrolleffect        int            `db:"scrolleffect"`        //	int(11)	NO		0
	Scrolltype          int            `db:"scrolltype"`          //	int(11)	NO		0
	Scrolllevel2        int            `db:"scrolllevel2"`        //	int(11)	NO		0
	Scrolllevel         int            `db:"scrolllevel"`         //	int(11)	NO		0
	Unk157              int            `db:"UNK157"`              //	int(11)	NO		0
	Serialized          mysql.NullTime `db:"serialized"`          //	datetime	YES		NULL
	Verified            mysql.NullTime `db:"verified"`            //	datetime	YES		NULL
	Serialization       sql.NullString `db:"serialization"`       //	text	YES		NULL
	Source              string         `db:"source"`              //	varchar(20)	NO
	Unk033              int            `db:"UNK033"`              //	int(11)	NO		0
	Lorefile            string         `db:"lorefile"`            //	varchar(32)	NO
	Unk014              int            `db:"UNK014"`              //	int(11)	NO		0
	Svcorruption        int            `db:"svcorruption"`        //	int(11)	NO		0
	Unk038              int            `db:"UNK038"`              //	int(11)	NO		0
	Unk060              int            `db:"UNK060"`              //	int(11)	NO		0
	Augslot1unk2        int            `db:"augslot1unk2"`        //	int(11)	NO		0
	Augslot2unk2        int            `db:"augslot2unk2"`        //	int(11)	NO		0
	Augslot3unk2        int            `db:"augslot3unk2"`        //	int(11)	NO		0
	Augslot4unk2        int            `db:"augslot4unk2"`        //	int(11)	NO		0
	Augslot5unk2        int            `db:"augslot5unk2"`        //	int(11)	NO		0
	Augslot6unk2        int            `db:"augslot6unk2"`        //	int(11)	NO		0
	Unk120              int            `db:"UNK120"`              //	int(11)	NO		0
	Unk121              int            `db:"UNK121"`              //	int(11)	NO		0
	Questitemflag       int            `db:"questitemflag"`       //	int(11)	NO		0
	Unk132              sql.NullString `db:"UNK132"`              //	text	NO		NULL
	Clickunk5           int            `db:"clickunk5"`           //	int(11)	NO		0
	Clickunk6           string         `db:"clickunk6"`           //	varchar(32)	NO
	Clickunk7           int            `db:"clickunk7"`           //	int(11)	NO		0
	Procunk1            int            `db:"procunk1"`            //	int(11)	NO		0
	Procunk2            int            `db:"procunk2"`            //	int(11)	NO		0
	Procunk3            int            `db:"procunk3"`            //	int(11)	NO		0
	Procunk4            int            `db:"procunk4"`            //	int(11)	NO		0
	Procunk6            string         `db:"procunk6"`            //	varchar(32)	NO
	Procunk7            int            `db:"procunk7"`            //	int(11)	NO		0
	Wornunk1            int            `db:"wornunk1"`            //	int(11)	NO		0
	Wornunk2            int            `db:"wornunk2"`            //	int(11)	NO		0
	Wornunk3            int            `db:"wornunk3"`            //	int(11)	NO		0
	Wornunk4            int            `db:"wornunk4"`            //	int(11)	NO		0
	Wornunk5            int            `db:"wornunk5"`            //	int(11)	NO		0
	Wornunk6            string         `db:"wornunk6"`            //	varchar(32)	NO
	Wornunk7            int            `db:"wornunk7"`            //	int(11)	NO		0
	Focusunk1           int            `db:"focusunk1"`           //	int(11)	NO		0
	Focusunk2           int            `db:"focusunk2"`           //	int(11)	NO		0
	Focusunk3           int            `db:"focusunk3"`           //	int(11)	NO		0
	Focusunk4           int            `db:"focusunk4"`           //	int(11)	NO		0
	Focusunk5           int            `db:"focusunk5"`           //	int(11)	NO		0
	Focusunk6           string         `db:"focusunk6"`           //	varchar(32)	NO
	Focusunk7           int            `db:"focusunk7"`           //	int(11)	NO		0
	Scrollunk1          int            `db:"scrollunk1"`          //	int(11)	NO		0
	Scrollunk2          int            `db:"scrollunk2"`          //	int(11)	NO		0
	Scrollunk3          int            `db:"scrollunk3"`          //	int(11)	NO		0
	Scrollunk4          int            `db:"scrollunk4"`          //	int(11)	NO		0
	Scrollunk5          int            `db:"scrollunk5"`          //	int(11)	NO		0
	Scrollunk6          string         `db:"scrollunk6"`          //	varchar(32)	NO
	Scrollunk7          int            `db:"scrollunk7"`          //	int(11)	NO		0
	Unk193              int            `db:"UNK193"`              //	int(11)	NO		0
	Purity              int            `db:"purity"`              //	int(11)	NO		0
	Evolvinglevel       int            `db:"evolvinglevel"`       //	int(11)	NO		0
	Clickname           string         `db:"clickname"`           //	varchar(64)	NO
	Procname            string         `db:"procname"`            //	varchar(64)	NO
	Wornname            string         `db:"wornname"`            //	varchar(64)	NO
	Focusname           string         `db:"focusname"`           //	varchar(64)	NO
	Scrollname          string         `db:"scrollname"`          //	varchar(64)	NO
	Dsmitigation        int            `db:"dsmitigation"`        //	smallint(6)	NO		0
	Heroicstr           int            `db:"heroic_str"`          //	smallint(6)	NO		0
	Heroicint           int            `db:"heroic_int"`          //	smallint(6)	NO		0
	Heroicwis           int            `db:"heroic_wis"`          //	smallint(6)	NO		0
	Heroicagi           int            `db:"heroic_agi"`          //	smallint(6)	NO		0
	Heroicdex           int            `db:"heroic_dex"`          //	smallint(6)	NO		0
	Heroicsta           int            `db:"heroic_sta"`          //	smallint(6)	NO		0
	Heroiccha           int            `db:"heroic_cha"`          //	smallint(6)	NO		0
	Heroicpr            int            `db:"heroic_pr"`           //	smallint(6)	NO		0
	Heroicdr            int            `db:"heroic_dr"`           //	smallint(6)	NO		0
	Heroicfr            int            `db:"heroic_fr"`           //	smallint(6)	NO		0
	Heroiccr            int            `db:"heroic_cr"`           //	smallint(6)	NO		0
	Heroicmr            int            `db:"heroic_mr"`           //	smallint(6)	NO		0
	Heroicsvcorrup      int            `db:"heroic_svcorrup"`     //	smallint(6)	NO		0
	Healamt             int            `db:"healamt"`             //	smallint(6)	NO		0
	Spelldmg            int            `db:"spelldmg"`            //	smallint(6)	NO		0
	Clairvoyance        int            `db:"clairvoyance"`        //	smallint(6)	NO		0
	Backstabdmg         int            `db:"backstabdmg"`         //	smallint(6)	NO		0
	Created             string         `db:"created"`             //	varchar(64)	NO
	Elitematerial       int            `db:"elitematerial"`       //	smallint(6)	NO		0
	Ldonsellbackrate    int            `db:"ldonsellbackrate"`    //	smallint(6)	NO		0
	Scriptfileid        int            `db:"scriptfileid"`        //	smallint(6)	NO		0
	Expendablearrow     int            `db:"expendablearrow"`     //	smallint(6)	NO		0
	Powersourcecapacity int            `db:"powersourcecapacity"` //	smallint(6)	NO		0
	Bardeffect          int            `db:"bardeffect"`          //	smallint(6)	NO		0
	Bardeffecttype      int            `db:"bardeffecttype"`      //	smallint(6)	NO		0
	Bardlevel2          int            `db:"bardlevel2"`          //	smallint(6)	NO		0
	Bardlevel           int            `db:"bardlevel"`           //	smallint(6)	NO		0
	Bardunk1            int            `db:"bardunk1"`            //	smallint(6)	NO		0
	Bardunk2            int            `db:"bardunk2"`            //	smallint(6)	NO		0
	Bardunk3            int            `db:"bardunk3"`            //	smallint(6)	NO		0
	Bardunk4            int            `db:"bardunk4"`            //	smallint(6)	NO		0
	Bardunk5            int            `db:"bardunk5"`            //	smallint(6)	NO		0
	Bardname            string         `db:"bardname"`            //	varchar(64)	NO
	Bardunk7            int            `db:"bardunk7"`            //	smallint(6)	NO		0
	Unk214              int            `db:"UNK214"`              //	smallint(6)	NO		0
	Unk219              int            `db:"UNK219"`              //	int(11)	NO		0
	Unk220              int            `db:"UNK220"`              //	int(11)	NO		0
	Unk221              int            `db:"UNK221"`              //	int(11)	NO		0
	Heirloom            int            `db:"heirloom"`            //	int(11)	NO		0
	Unk223              int            `db:"UNK223"`              //	int(11)	NO		0
	Unk224              int            `db:"UNK224"`              //	int(11)	NO		0
	Unk225              int            `db:"UNK225"`              //	int(11)	NO		0
	Unk226              int            `db:"UNK226"`              //	int(11)	NO		0
	Unk227              int            `db:"UNK227"`              //	int(11)	NO		0
	Unk228              int            `db:"UNK228"`              //	int(11)	NO		0
	Unk229              int            `db:"UNK229"`              //	int(11)	NO		0
	Unk230              int            `db:"UNK230"`              //	int(11)	NO		0
	Unk231              int            `db:"UNK231"`              //	int(11)	NO		0
	Unk232              int            `db:"UNK232"`              //	int(11)	NO		0
	Unk233              int            `db:"UNK233"`              //	int(11)	NO		0
	Unk234              int            `db:"UNK234"`              //	int(11)	NO		0
	Placeable           int            `db:"placeable"`           //	int(11)	NO		0
	Unk236              int            `db:"UNK236"`              //	int(11)	NO		0
	Unk237              int            `db:"UNK237"`              //	int(11)	NO		0
	Unk238              int            `db:"UNK238"`              //	int(11)	NO		0
	Unk239              int            `db:"UNK239"`              //	int(11)	NO		0
	Unk240              int            `db:"UNK240"`              //	int(11)	NO		0
	Unk241              int            `db:"UNK241"`              //	int(11)	NO		0
	epicitem            int            `db:"epicitem"`            //	int(11)	NO		0
}
