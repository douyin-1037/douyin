package cos

type CosVideo struct {
	MachineId   uint16
	VideoBucket string
	CoverBucket string
	SecretID    string
	SecretKey   string
}

var cosVideo CosVideo

func Init() {

}
