package judgement

import "time"

type Purpose int8
type Resolution int8
type GraphicQuality int8
type TargetFPS int8
type ScreenRatio int8

const (
	Gaming      Purpose = 1
	Engineering Purpose = 2
)

const (
	TenEighty     Resolution = 1
	FourteenForty Resolution = 2
	FourK         Resolution = 4
)

const (
	Low    GraphicQuality = 1
	Medium GraphicQuality = 2
	High   GraphicQuality = 4
	Ultra  GraphicQuality = 8
)

const (
	Sixty            TargetFPS = 1
	HundredTwenty    TargetFPS = 2
	HundredFortyFour TargetFPS = 4
	TwoHundredForty  TargetFPS = 8
)

const (
	OneAndOne      ScreenRatio = 1
	FourAndThree   ScreenRatio = 2
	FiveAndFour    ScreenRatio = 4
	SixteenAndNine ScreenRatio = 8
)

type Judgement struct {
	ID             int64          `bson:"_id,omitempty"`
	UserID         int64          `bson:"user_id"`
	Purpose        Purpose        `bson:"purpose"`
	GameID         int64          `bson:"game_id"`
	Resolution     Resolution     `bson:"resolution"`
	GraphicQuality GraphicQuality `bson:"graphic_quality"`
	TargetFPS      TargetFPS      `bson:"target_fps"`
	Description    string         `bson:"description"`
	Care           Care           `bson:"care"`
	Budget         Budget         `bson:"budget"`
	CreatedAt      time.Time      `bson:"created_at"`
}

type Care struct {
	HighestFPS        bool `bson:"highest_fps"`
	BestVisualQuality bool `bson:"best_visual_quality"`
	BestValueForMoney bool `bson:"best_value_for_money"`
	QuietOperation    bool `bson:"quiet_operation"`
	GoodCooling       bool `bson:"good_cooling"`
	BatteryLife       bool `bson:"battery_life"`
	Portability       bool `bson:"portability"`
	Upgradability     bool `bson:"upgradability"`
}

type Budget struct {
	Adaptable    bool `bson:"adaptable"`
	DesireAmount bool `bson:"desire_amount"` // percentage based
}

type Additional struct {
	FutureUpgradable bool        `bson:"future_upgradable"`
	WifiNeeded       bool        `bson:"wifi_needed"`
	MonitorNeeded    bool        `bson:"monitor_needed"`
	AccessoryNeeded  bool        `bson:"accessory_needed"`
	ScreenRatio      ScreenRatio `bson:"screen_ratio"`
	TransferCare     bool        `bson:"transfer_care"`
	PreferUSBC       bool        `bson:"prefer_usb_c"`
	PreferRGB        bool        `bson:"prefer_rgb"`
}
