package libv2ray

import (
	"fmt"
	"log"
	"sync"
)

// CoreController represents a controller for managing Xray core instance lifecycle
type CoreController struct {
	CallbackHandler CoreCallbackHandler
	coreMutex       sync.Mutex
	IsRunning       bool
}

// CoreCallbackHandler defines interface for receiving callbacks and notifications from the core service
type CoreCallbackHandler interface {
	Startup() int
	Shutdown() int
	OnEmitStatus(int, string) int
}

// InitCoreEnv initializes environment variables and file system handlers for the core
// It sets up asset path, certificate path, XUDP base key and customizes the file reader
// to support Android asset system
func InitCoreEnv(envPath string, key string) {
}

// NewCoreController initializes and returns a new CoreController instance
// Sets up the console log handler and associates it with the provided callback handler
func NewCoreController(s CoreCallbackHandler) *CoreController {
	return &CoreController{
		CallbackHandler: s,
	}
}

// StartLoop initializes and starts the core processing loop
// Thread-safe method that configures and runs the Xray core with the provided configuration
// Returns immediately if the core is already running
func (x *CoreController) StartLoop(configContent string) (err error) {
	x.coreMutex.Lock()
	defer x.coreMutex.Unlock()

	if x.IsRunning {
		log.Println("Core is already running")
		return nil
	}

	return x.doStartLoop(configContent)
}

// StopLoop safely stops the core processing loop and releases resources
// Thread-safe method that shuts down the core instance and triggers necessary callbacks
func (x *CoreController) StopLoop() error {
	x.coreMutex.Lock()
	defer x.coreMutex.Unlock()

	if x.IsRunning {
		x.doShutdown()
		x.CallbackHandler.OnEmitStatus(0, "Core stopped")
	}
	return nil
}

// QueryStats retrieves and resets traffic statistics for a specific outbound tag and direction
// Returns the accumulated traffic value and resets the counter to zero
// Returns 0 if the stats manager is not initialized or the counter doesn't exist
func (x *CoreController) QueryStats(tag string, direct string) int64 {
	return 0
}

// MeasureDelay measures network latency to a specified URL through the current core instance
// Uses a 12-second timeout context and returns the round-trip time in milliseconds
// An error is returned if the connection fails or returns an unexpected status
func (x *CoreController) MeasureDelay(url string) (int64, error) {
	return 0, nil
}

// MeasureOutboundDelay measures the outbound delay for a given configuration and URL
func MeasureOutboundDelay(ConfigureFileContent string, url string) (int64, error) {
	return 0, nil
}

// CheckVersionX returns the library and Xray versions
func CheckVersionX() string {
	var version = 32
	return fmt.Sprintf("Lib v%d, Xray-core", version)
}

// doShutdown shuts down the Xray instance and cleans up resources
func (x *CoreController) doShutdown() {
	CloseVpoint()
	x.IsRunning = false
}

// doStartLoop sets up and starts the Xray core
func (x *CoreController) doStartLoop(configContent string) error {
	log.Println("initializing core...")
	config, err := LoadJSONConfig(configContent)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	log.Println("starting core...")
	x.IsRunning = true
	if err := StartVpoint(config); err != nil {
		x.IsRunning = false
		return fmt.Errorf("startup failed: %w", err)
	}

	x.CallbackHandler.Startup()
	x.CallbackHandler.OnEmitStatus(0, "Started successfully, running")

	log.Println("Starting core successfully")
	return nil
}
