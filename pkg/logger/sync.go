package logger

// Sync calls the underlying Core's Sync method, flushing any buffered log entries.
// Applications should take care to call Sync before exiting.
func (l *Logger) Sync() error {
	return l.logger.Sync()
}
