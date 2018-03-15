package metrics

import "github.com/prometheus/client_golang/prometheus"

const (
	arkNamespace = "heptio_ark"
)

var (
	backupCount = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: arkNamespace,
		Subsystem: "backup_sync_controller",
		Name:      "backup_count",
		Help:      "Number of backups stored in the object store",
	})

	backupsInProgress = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: arkNamespace,
		Subsystem: "backup_controller",
		Name:      "backups_in_progress",
		Help:      "Number of backups in progress",
	})

	backupsFailed = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: arkNamespace,
		Subsystem: "backup_controller",
		Name:      "backups_failed",
		Help:      "Number of backups in progress",
	})

	backupsUploaded = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: arkNamespace,
		Subsystem: "backup_service",
		Name:      "backup_upload_count",
		Help:      "Number of backups uploaded",
	})

	testCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: arkNamespace,
		Subsystem: "testSubsystem",
		Name:      "testCounter",
		Help:      "my test counter",
	})
)

func init() {
	prometheus.MustRegister(backupCount)
	prometheus.MustRegister(testCounter)
	prometheus.MustRegister(backupsInProgress)
	prometheus.MustRegister(backupsFailed)
}

// UpdateUnneededNodesCount records number of currently unneeded nodes
func UpdateBackupCount(count int) {
	backupCount.Set(float64(count))
}

// IncrementTestCounter increments testCounter by incrBy
func IncrementTestCounter(incrBy int) {
	testCounter.Add(float64(incrBy))
}

// IncrementBackupsInProgress increments backupsInProgress
func IncrementBackupsInProgress() {
	backupsInProgress.Inc()
}

// DecrementBackupsInProgress decrements backupsInProgress
func DecrementBackupsInProgress() {
	backupsInProgress.Dec()
}

// IncrementBackupsFailed increments backupsFailed
func IncrementBackupsFailed() {
	backupsFailed.Inc()
}

// DecrementBackupsFailed decrements backupsFailed
func DecrementBackupsFailed() {
	backupsFailed.Dec()
}

// IncrementBackupsUploaded increments backupsFailed
func IncrementBackupsUploaded() {
	backupsUploaded.Inc()
}
