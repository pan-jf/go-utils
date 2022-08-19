package predis

import (
	"fmt"
	"time"

	jsonIter "github.com/json-iterator/go"
	"github.com/pan-jf/go-utils/plog"
	"go.uber.org/zap"
)

func SetWithMarshal(redisKey string, data interface{}, expiration time.Duration) bool {
	if data == nil || len(redisKey) <= 0 {
		return false
	}
	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		plog.Error("SetWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false
	}

	err = GlobalRedis.Set(redisKey, dataMarshal, expiration)
	if err != nil {
		plog.Error("SetWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
	}

	return err == nil
}

func GetWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.Get(redisKey)
	if err != nil && err.Error() == RedisNil {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		plog.Error("GetWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("redisData", redisData),
		)
		return false, err
	}

	return true, nil
}

func HSetWithMarshal(redisKey string, fieldName string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 || len(fieldName) <= 0 {
		return false, nil
	}

	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		plog.Error("HSetWithMarshal err",
			zap.String("Param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.Error(err),
			zap.ByteString("dataMarshal", dataMarshal),
		)
		return false, err
	}

	err = GlobalRedis.HSet(redisKey, fieldName, dataMarshal)
	if err != nil {
		plog.Error("HSetWithMarshal set err",
			zap.String("Param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.ByteString("dataMarshal", dataMarshal))
		return false, err
	}

	return true, nil
}

func HGetWithUnMarshal(redisKey string, fieldName string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 || len(fieldName) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.HGet(redisKey, fieldName)
	if err != nil && err.Error() == RedisNil {
		return false, nil
	}
	if err != nil {
		plog.Error("LoadRedis", zap.Error(err))
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		plog.Error("HGetWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.Error(err),
			zap.String("redisData", redisData),
		)
		return false, err
	}

	return true, nil
}

func RPushWithMarshal(redisKey string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 {
		return false, nil
	}

	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		plog.Error("RPushWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	err = GlobalRedis.RPush(redisKey, dataMarshal)
	if err != nil {
		plog.Error("RPushWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	return true, nil
}

func LPushWithMarshal(redisKey string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 {
		return false, nil
	}

	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		plog.Error("LPushWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	err = GlobalRedis.LPush(redisKey, dataMarshal)
	if err != nil {
		plog.Error("LPushWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	return true, nil
}

func LPopWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.LPop(redisKey)
	if err == nil {
		err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
		if err != nil {
			plog.Error("LPopWithUnMarshal Json Unmarshal err",
				zap.String("param", fmt.Sprintf("(%v)", redisKey)),
				zap.Error(err),
				zap.String("redisData", redisData),
			)
			return false, err
		}
	} else {
		plog.Error("LPopWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData),
		)
		return false, err
	}

	return true, nil
}

func RPopWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.RPop(redisKey)
	if err == nil {
		err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
		if err != nil {
			plog.Error("RPopWithUnMarshal Json Unmarshal err",
				zap.String("param", fmt.Sprintf("(%v)", redisKey)),
				zap.Error(err),
				zap.String("redisData", redisData),
			)
			return false, err
		}
	} else {
		plog.Error("RPopWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData),
		)
		return false, err
	}

	return true, nil
}

func LTrimWithUnMarshal(redisKey string, start, stop int64) (bool, error) {
	if len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.LTrim(redisKey, start, stop)
	if err != nil {
		plog.Error("LTrimWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData))
		return false, err
	}

	return true, nil
}

func LIndexWithUnMarshal(redisKey string, index int64, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := GlobalRedis.LIndex(redisKey, index)
	if err != nil {
		plog.Error("LIndexWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData))
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		plog.Error("LIndexWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("redisData", redisData))
		return false, err
	}

	return true, nil
}
