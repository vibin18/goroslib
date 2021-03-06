// Package sensor_msgs contains message definitions (autogenerated).
package sensor_msgs

import (
	"github.com/aler9/goroslib/msg"
	"github.com/aler9/goroslib/msgs/geometry_msgs"
	"github.com/aler9/goroslib/msgs/std_msgs"
	"time"
)

const (
	POWER_SUPPLY_STATUS_UNKNOWN               uint8  = 0
	POWER_SUPPLY_STATUS_CHARGING              uint8  = 1
	POWER_SUPPLY_STATUS_DISCHARGING           uint8  = 2
	POWER_SUPPLY_STATUS_NOT_CHARGING          uint8  = 3
	POWER_SUPPLY_STATUS_FULL                  uint8  = 4
	POWER_SUPPLY_HEALTH_UNKNOWN               uint8  = 0
	POWER_SUPPLY_HEALTH_GOOD                  uint8  = 1
	POWER_SUPPLY_HEALTH_OVERHEAT              uint8  = 2
	POWER_SUPPLY_HEALTH_DEAD                  uint8  = 3
	POWER_SUPPLY_HEALTH_OVERVOLTAGE           uint8  = 4
	POWER_SUPPLY_HEALTH_UNSPEC_FAILURE        uint8  = 5
	POWER_SUPPLY_HEALTH_COLD                  uint8  = 6
	POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE uint8  = 7
	POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE   uint8  = 8
	POWER_SUPPLY_TECHNOLOGY_UNKNOWN           uint8  = 0
	POWER_SUPPLY_TECHNOLOGY_NIMH              uint8  = 1
	POWER_SUPPLY_TECHNOLOGY_LION              uint8  = 2
	POWER_SUPPLY_TECHNOLOGY_LIPO              uint8  = 3
	POWER_SUPPLY_TECHNOLOGY_LIFE              uint8  = 4
	POWER_SUPPLY_TECHNOLOGY_NICD              uint8  = 5
	POWER_SUPPLY_TECHNOLOGY_LIMN              uint8  = 6
	TYPE_LED                                  uint8  = 0
	TYPE_RUMBLE                               uint8  = 1
	TYPE_BUZZER                               uint8  = 2
	COVARIANCE_TYPE_UNKNOWN                   uint8  = 0
	COVARIANCE_TYPE_APPROXIMATED              uint8  = 1
	COVARIANCE_TYPE_DIAGONAL_KNOWN            uint8  = 2
	COVARIANCE_TYPE_KNOWN                     uint8  = 3
	STATUS_NO_FIX                             int8   = -1
	STATUS_FIX                                int8   = 0
	STATUS_SBAS_FIX                           int8   = 1
	STATUS_GBAS_FIX                           int8   = 2
	SERVICE_GPS                               uint16 = 1
	SERVICE_GLONASS                           uint16 = 2
	SERVICE_COMPASS                           uint16 = 4
	SERVICE_GALILEO                           uint16 = 8
	INT8                                      uint8  = 1
	UINT8                                     uint8  = 2
	INT16                                     uint8  = 3
	UINT16                                    uint8  = 4
	INT32                                     uint8  = 5
	UINT32                                    uint8  = 6
	FLOAT32                                   uint8  = 7
	FLOAT64                                   uint8  = 8
	ULTRASOUND                                uint8  = 0
	INFRARED                                  uint8  = 1
)

type BatteryState struct {
	msg.Package           `ros:"sensor_msgs"`
	msg.Definitions       `ros:"uint8 POWER_SUPPLY_STATUS_UNKNOWN=0,uint8 POWER_SUPPLY_STATUS_CHARGING=1,uint8 POWER_SUPPLY_STATUS_DISCHARGING=2,uint8 POWER_SUPPLY_STATUS_NOT_CHARGING=3,uint8 POWER_SUPPLY_STATUS_FULL=4,uint8 POWER_SUPPLY_HEALTH_UNKNOWN=0,uint8 POWER_SUPPLY_HEALTH_GOOD=1,uint8 POWER_SUPPLY_HEALTH_OVERHEAT=2,uint8 POWER_SUPPLY_HEALTH_DEAD=3,uint8 POWER_SUPPLY_HEALTH_OVERVOLTAGE=4,uint8 POWER_SUPPLY_HEALTH_UNSPEC_FAILURE=5,uint8 POWER_SUPPLY_HEALTH_COLD=6,uint8 POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE=7,uint8 POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE=8,uint8 POWER_SUPPLY_TECHNOLOGY_UNKNOWN=0,uint8 POWER_SUPPLY_TECHNOLOGY_NIMH=1,uint8 POWER_SUPPLY_TECHNOLOGY_LION=2,uint8 POWER_SUPPLY_TECHNOLOGY_LIPO=3,uint8 POWER_SUPPLY_TECHNOLOGY_LIFE=4,uint8 POWER_SUPPLY_TECHNOLOGY_NICD=5,uint8 POWER_SUPPLY_TECHNOLOGY_LIMN=6"`
	Header                std_msgs.Header
	Voltage               float32
	Temperature           float32
	Current               float32
	Charge                float32
	Capacity              float32
	DesignCapacity        float32
	Percentage            float32
	PowerSupplyStatus     uint8
	PowerSupplyHealth     uint8
	PowerSupplyTechnology uint8
	Present               bool
	CellVoltage           []float32
	CellTemperature       []float32
	Location              string
	SerialNumber          string
}

type CameraInfo struct {
	msg.Package     `ros:"sensor_msgs"`
	Header          std_msgs.Header
	Height          uint32
	Width           uint32
	DistortionModel string
	D               []float64
	K               [9]float64
	R               [9]float64
	P               [12]float64
	BinningX        uint32
	BinningY        uint32
	Roi             RegionOfInterest
}

type ChannelFloat32 struct {
	msg.Package `ros:"sensor_msgs"`
	Name        string
	Values      []float32
}

type CompressedImage struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Format      string
	Data        []uint8
}

type FluidPressure struct {
	msg.Package   `ros:"sensor_msgs"`
	Header        std_msgs.Header
	FluidPressure float64
	Variance      float64
}

type Illuminance struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Illuminance float64
	Variance    float64
}

type Image struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Encoding    string
	IsBigendian uint8
	Step        uint32
	Data        []uint8
}

type Imu struct {
	msg.Package                  `ros:"sensor_msgs"`
	Header                       std_msgs.Header
	Orientation                  geometry_msgs.Quaternion
	OrientationCovariance        [9]float64
	AngularVelocity              geometry_msgs.Vector3
	AngularVelocityCovariance    [9]float64
	LinearAcceleration           geometry_msgs.Vector3
	LinearAccelerationCovariance [9]float64
}

type JointState struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Name        []string
	Position    []float64
	Velocity    []float64
	Effort      []float64
}

type Joy struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Axes        []float32
	Buttons     []int32
}

type JoyFeedback struct {
	msg.Package     `ros:"sensor_msgs"`
	msg.Definitions `ros:"uint8 TYPE_LED=0,uint8 TYPE_RUMBLE=1,uint8 TYPE_BUZZER=2"`
	Type            uint8
	Id              uint8
	Intensity       float32
}

type JoyFeedbackArray struct {
	msg.Package `ros:"sensor_msgs"`
	Array       []JoyFeedback
}

type LaserEcho struct {
	msg.Package `ros:"sensor_msgs"`
	Echoes      []float32
}

type LaserScan struct {
	msg.Package    `ros:"sensor_msgs"`
	Header         std_msgs.Header
	AngleMin       float32
	AngleMax       float32
	AngleIncrement float32
	TimeIncrement  float32
	ScanTime       float32
	RangeMin       float32
	RangeMax       float32
	Ranges         []float32
	Intensities    []float32
}

type MagneticField struct {
	msg.Package             `ros:"sensor_msgs"`
	Header                  std_msgs.Header
	MagneticField           geometry_msgs.Vector3
	MagneticFieldCovariance [9]float64
}

type MultiDOFJointState struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	JointNames  []string
	Transforms  []geometry_msgs.Transform
	Twist       []geometry_msgs.Twist
	Wrench      []geometry_msgs.Wrench
}

type MultiEchoLaserScan struct {
	msg.Package    `ros:"sensor_msgs"`
	Header         std_msgs.Header
	AngleMin       float32
	AngleMax       float32
	AngleIncrement float32
	TimeIncrement  float32
	ScanTime       float32
	RangeMin       float32
	RangeMax       float32
	Ranges         []LaserEcho
	Intensities    []LaserEcho
}

type NavSatFix struct {
	msg.Package            `ros:"sensor_msgs"`
	msg.Definitions        `ros:"uint8 COVARIANCE_TYPE_UNKNOWN=0,uint8 COVARIANCE_TYPE_APPROXIMATED=1,uint8 COVARIANCE_TYPE_DIAGONAL_KNOWN=2,uint8 COVARIANCE_TYPE_KNOWN=3"`
	Header                 std_msgs.Header
	Status                 NavSatStatus
	Latitude               float64
	Longitude              float64
	Altitude               float64
	PositionCovariance     [9]float64
	PositionCovarianceType uint8
}

type NavSatStatus struct {
	msg.Package     `ros:"sensor_msgs"`
	msg.Definitions `ros:"int8 STATUS_NO_FIX=-1,int8 STATUS_FIX=0,int8 STATUS_SBAS_FIX=1,int8 STATUS_GBAS_FIX=2,uint16 SERVICE_GPS=1,uint16 SERVICE_GLONASS=2,uint16 SERVICE_COMPASS=4,uint16 SERVICE_GALILEO=8"`
	Status          int8
	Service         uint16
}

type PointCloud struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Points      []geometry_msgs.Point32
	Channels    []ChannelFloat32
}

type PointCloud2 struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Height      uint32
	Width       uint32
	Fields      []PointField
	IsBigendian bool
	PointStep   uint32
	RowStep     uint32
	Data        []uint8
	IsDense     bool
}

type PointField struct {
	msg.Package     `ros:"sensor_msgs"`
	msg.Definitions `ros:"uint8 INT8=1,uint8 UINT8=2,uint8 INT16=3,uint8 UINT16=4,uint8 INT32=5,uint8 UINT32=6,uint8 FLOAT32=7,uint8 FLOAT64=8"`
	Name            string
	Offset          uint32
	Datatype        uint8
	Count           uint32
}

type Range struct {
	msg.Package     `ros:"sensor_msgs"`
	msg.Definitions `ros:"uint8 ULTRASOUND=0,uint8 INFRARED=1"`
	Header          std_msgs.Header
	RadiationType   uint8
	FieldOfView     float32
	MinRange        float32
	MaxRange        float32
	Range           float32
}

type RegionOfInterest struct {
	msg.Package `ros:"sensor_msgs"`
	XOffset     uint32
	YOffset     uint32
	Height      uint32
	Width       uint32
	DoRectify   bool
}

type RelativeHumidity struct {
	msg.Package      `ros:"sensor_msgs"`
	Header           std_msgs.Header
	RelativeHumidity float64
	Variance         float64
}

type Temperature struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	Temperature float64
	Variance    float64
}

type TimeReference struct {
	msg.Package `ros:"sensor_msgs"`
	Header      std_msgs.Header
	TimeRef     time.Time
	Source      string
}
