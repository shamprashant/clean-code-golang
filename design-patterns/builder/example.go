package main

type BackupJob struct {
	Source             string
	Destination        string
	CompressionEnabled bool
	EncryptionType     string
	RetryCount         int
	Schedule           string
	NotifyChannels     []string
}

// ✅ The Builder
type BackupJobBuilder struct {
	job BackupJob
}

// ✅ Create a new builder with defaults
func NewBackupJobBuilder() *BackupJobBuilder {
	return &BackupJobBuilder{
		job: BackupJob{
			CompressionEnabled: false,
			RetryCount:         0,
		},
	}
}

// 🔧 Step-by-step fluent methods

func (b *BackupJobBuilder) From(source string) *BackupJobBuilder {
	b.job.Source = source
	return b
}

func (b *BackupJobBuilder) To(destination string) *BackupJobBuilder {
	b.job.Destination = destination
	return b
}

func (b *BackupJobBuilder) EnableCompression() *BackupJobBuilder {
	b.job.CompressionEnabled = true
	return b
}

func (b *BackupJobBuilder) WithEncryption(enc string) *BackupJobBuilder {
	b.job.EncryptionType = enc
	return b
}

func (b *BackupJobBuilder) WithRetries(count int) *BackupJobBuilder {
	b.job.RetryCount = count
	return b
}

func (b *BackupJobBuilder) OnSchedule(schedule string) *BackupJobBuilder {
	b.job.Schedule = schedule
	return b
}

func (b *BackupJobBuilder) Notify(channels ...string) *BackupJobBuilder {
	b.job.NotifyChannels = append(b.job.NotifyChannels, channels...)
	return b
}

// ✅ Final call
func (b *BackupJobBuilder) Build() BackupJob {
	return b.job
}

// func main() {
// 	job := NewBackupJobBuilder().
// 		From("db-server").
// 		To("s3://daily-backups").
// 		EnableCompression().
// 		WithEncryption("AES256").
// 		WithRetries(3).
// 		OnSchedule("nightly").
// 		Notify("email", "slack").
// 		Build()

// 	fmt.Printf("📦 BackupJob: %+v\n", job)
// }

/*
✅ Benefits:
Clean, expressive, fluent API

Avoids messy constructors with 6+ arguments

Only Build() returns the complete object

Makes optional fields super readable
*/
