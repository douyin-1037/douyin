package cos

import "douyin/common/conf"

type CosVideo struct {
	VideoBucket string
	CoverBucket string
	SecretID    string
	SecretKey   string
}

var cosVideo CosVideo

func Init() {
	cosVideo = CosVideo{
		VideoBucket: conf.COS.VideoBucket,
		CoverBucket: conf.COS.CoverBucket,
		SecretID:    conf.COS.SecretID,
		SecretKey:   conf.COS.SecretKey,
	}
}
