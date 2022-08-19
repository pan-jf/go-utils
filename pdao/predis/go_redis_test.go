package predis

import (
	"testing"
	"time"

	"go.uber.org/zap"

	jsonIter "github.com/json-iterator/go"
	"github.com/pan-jf/go-utils/plog"
)

type TeamUser struct {
	UID      string `json:"uid,omitempty"`
	Cow      int32  `json:"cow,omitempty"`      // AI
	CowLevel int32  `json:"cowLevel,omitempty"` // AI LEVEL
	SeatID   int32  `json:"seatId"`             // 座位ID
}

func initRedis() bool {
	redisCfg := &RedisCfg{
		Host:     "192.168.3.37:6379",
		Password: "mega@dev",
		DB:       7,
		Cluster:  false,
	}

	err := GlobalRedis.Setup(redisCfg)
	if err != nil {
		plog.Error("setup error", zap.Error(err))
		return false
	}
	return true
}

func TestRedis(t *testing.T) {
	var err error
	if !initRedis() {
		return
	}
	err = GlobalRedis.RPush("test", "key1", "key2")
	if err != nil {
		plog.Error("RPush test", zap.Error(err))
		return
	}

	count, err := GlobalRedis.LRem("test", 0, "key1")
	plog.Info("LRem test", zap.Any("count", count), zap.Error(err))

	count, err = GlobalRedis.LRem("test", 0, "key2")
	plog.Info("LRem test", zap.Any("count", count), zap.Error(err))

	count, err = GlobalRedis.LRem("test", 0, "key2")
	plog.Info("LRem test", zap.Any("count", count), zap.Error(err))

	v, err := GlobalRedis.LPop("test")
	plog.Info("LPop test", zap.Any("v", v), zap.Error(err))

	teamUser := &TeamUser{
		UID:      "234",
		Cow:      1,
		CowLevel: 2,
		SeatID:   0,
	}

	docData, err := jsonIter.Marshal(teamUser)

	//teamUserStr := fmt.Sprintf("%v|%v|%v|%v", teamUser.UID, teamUser.Cow, teamUser.CowLevel, teamUser.SeatID)

	err = GlobalRedis.RPush("test", docData)
	if err != nil {
		plog.Error("RPush test", zap.Error(err))
		return
	}

	v, err = GlobalRedis.LPop("test")
	plog.Info("LPop test", zap.Any("v", v), zap.Error(err))

	teamUser2 := &TeamUser{}
	_ = jsonIter.Unmarshal([]byte(v), teamUser2)

	err = GlobalRedis.RPush("test", docData)
	if err != nil {
		plog.Error("RPush test", zap.Error(err))
		return
	}

	teamUser3 := &TeamUser{
		UID:      "243234234",
		Cow:      11,
		CowLevel: 12,
		SeatID:   0,
	}

	docData3, err := jsonIter.Marshal(teamUser3)

	err = GlobalRedis.RPush("test", docData3)
	if err != nil {
		plog.Error("RPush test", zap.Error(err))
		return
	}

	teamUserList, _ := GlobalRedis.LRange("test", 0, 10)
	for _, teamUserItem := range teamUserList {
		teamUser4 := &TeamUser{}
		_ = jsonIter.Unmarshal([]byte(teamUserItem), teamUser4)

		plog.Info("Unmarshal", zap.Any("teamUser4", teamUser4), zap.Error(err))
	}

}

type TestJson struct {
	AAA string
}

func TestSetWithMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	SetWithMarshal("jsonData", &TestJson{AAA: "BBB"}, 1*time.Minute)
}

func TestGetWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	var data = &TestJson{}

	isOk, err := GetWithUnMarshal("jsonData", data)
	if err != nil {
		return
	}
	plog.Info("GetWithUnMarshal", zap.Bool("isOk", isOk), zap.Any("jsonData", data), zap.Error(err))
}

func TestHSetWithMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	isOk, err := HSetWithMarshal("jsonDataHSet", "filedTest", &TestJson{AAA: "BBB"})
	if err != nil {
		return
	}
	plog.Info("HSetWithMarshal", zap.Bool("isOk", isOk), zap.Error(err))
}

func TestHGetWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	var data = &TestJson{}

	isOk, err := HGetWithUnMarshal("jsonDataHSet", "filedTest", data)
	if err != nil {
		return
	}
	plog.Info("HGetWithUnMarshal", zap.Bool("isOk", isOk), zap.Any("data", data), zap.Error(err))
}

func TestRPushWithMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	isOk, err := RPushWithMarshal("jsonDataRPush", &TestJson{AAA: "BBB"})
	if err != nil {
		return
	}
	plog.Info("RPushWithMarshal", zap.Bool("isOk", isOk), zap.Error(err))
}

func TestLPopWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}
	var data = &TestJson{}

	isOk, err := LPopWithUnMarshal("jsonDataRPush", data)
	if err != nil {
		return
	}
	plog.Info("RPushWithMarshal", zap.Bool("isOk", isOk), zap.Any("data", data), zap.Error(err))
}

func TestLPushWithMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	isOk, err := LPushWithMarshal("jsonDataLPush", &TestJson{AAA: "444"})
	if err != nil {
		return
	}
	plog.Info("LPushWithMarshal", zap.Bool("isOk", isOk), zap.Error(err))
}

func TestRPopWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}
	var data = &TestJson{}

	isOk, err := RPopWithUnMarshal("jsonDataLPush", data)
	if err != nil {
		return
	}
	plog.Info("LPopWithUnMarshal", zap.Bool("isOk", isOk), zap.Any("data", data), zap.Error(err))
}

func TestLTrimWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}

	isOk, err := LTrimWithUnMarshal("jsonDataLPush", 1, -1)
	if err != nil {
		return
	}
	plog.Info("LTrimWithUnMarshal", zap.Bool("isOk", isOk), zap.Error(err))
}

func TestLIndexWithUnMarshal(t *testing.T) {
	if !initRedis() {
		return
	}
	var data = &TestJson{}

	isOk, err := LIndexWithUnMarshal("jsonDataLPush", 0, data)
	if err != nil {
		return
	}
	plog.Info("LIndexWithUnMarshal", zap.Bool("isOk", isOk), zap.Any("data", data), zap.Error(err))
}
