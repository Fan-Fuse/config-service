package service

// A list of allowed keys with their default values, value is a variant (string, int, bool, etc.)
var AllowedKeys = map[string]interface{}{
	"APP_ENV":                 "development",
	"APP_MAINTENANCE":         false,
	"APP_MAINTENANCE_MESSAGE": "We are currently down for maintenance, please check back later.",
	"APP_VERSION":             "alpha-0.0.1",
	"USER_REGISTRATION_OPEN":  true,
	"CRAWLER_ENABLED":         true,
	"CRAWLER_INTERVAL_ARTIST": 43200, // in seconds, per artist, default 12 hours
	"CRAWLER_INTERVAL_USER":   86400, // in seconds, per user, default 24 hours
	"CRAWLER_BATCH_SIZE":      10,    // number of artists that can be queued to crawl at once
}
