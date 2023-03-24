package libv2ray

import (
	"fmt"
	"log"
	"sync"
)

/*V2RayPoint V2Ray Point Server
This is territory of Go, so no getter and setters!
*/
type V2RayPoint struct {
	SupportSet V2RayVPNServiceSupportsSet

	v2rayOP sync.Mutex

	IsRunning            bool
	ConfigureFileContent string

	// not used, just for compatibility
	DomainName   string
	AsyncResolve bool
}

/*V2RayVPNServiceSupportsSet To support Android VPN mode*/
type V2RayVPNServiceSupportsSet interface {
	Setup(Conf string) int
	Prepare() int
	Shutdown() int
	Protect(int) bool
	OnEmitStatus(int, string) int
}

/*RunLoop Run V2Ray main loop
 */
func (v *V2RayPoint) RunLoop(prefIPv6 bool) (err error) {
	v.v2rayOP.Lock()
	defer v.v2rayOP.Unlock()
	//Construct Context

	if !v.IsRunning {
		err = v.pointloop()
	}
	return
}

/*StopLoop Stop V2Ray main loop
 */
func (v *V2RayPoint) StopLoop() (err error) {
	v.v2rayOP.Lock()
	defer v.v2rayOP.Unlock()
	if v.IsRunning {
		v.shutdownInit()
		v.SupportSet.OnEmitStatus(0, "Closed")
	}
	return
}

//Delegate Funcation
func (v *V2RayPoint) QueryStats(tag string, direct string) int64 {
	return 0
}

func (v *V2RayPoint) shutdownInit() {
	v.IsRunning = false
	CloseVpoint()
}

func (v *V2RayPoint) pointloop() error {
	log.Println("loading core config")
	config, err := LoadJSONConfig(v.ConfigureFileContent)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("start core")
	v.IsRunning = true
	if err := StartVpoint(config); err != nil {
		v.IsRunning = false
		log.Println(err)
		return err
	}

	v.SupportSet.Prepare()
	v.SupportSet.Setup("")
	v.SupportSet.OnEmitStatus(0, "Running")
	return nil
}

func (v *V2RayPoint) MeasureDelay() (int64, error) {
	return 0, nil
}

// InitV2Env set v2 asset path
func InitV2Env(envPath string) {
}

//Delegate Funcation
func TestConfig(ConfigureFileContent string) error {
	_, err := LoadJSONConfig(ConfigureFileContent)
	return err
}

func MeasureOutboundDelay(ConfigureFileContent string) (int64, error) {
	return 0, nil
}

/*NewV2RayPoint new V2RayPoint*/
func NewV2RayPoint(s V2RayVPNServiceSupportsSet, adns bool) *V2RayPoint {
	return &V2RayPoint{
		SupportSet:   s,
		AsyncResolve: adns,
	}
}

/*CheckVersionX string
This func will return libv2ray binding version and V2Ray version used.
*/
func CheckVersionX() string {
	var version = 24
	return fmt.Sprintf("Lib v%d, Xray-core v1", version)
}
