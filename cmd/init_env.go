package main

import (
	"log"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	// Load .env jika ada (aman jika tidak ada)
	_ = godotenv.Load()

	// Biar viper juga baca dari environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Alias / fallback untuk nama variabel yang digunakan di workflow
	// Workflow menulis SERVICE_PORT / HTTP_HOST / DB_POSTGRES_*
	// Aplikasi sebelumnya menggunakan APP_PORT / APP_HOST, jadi kita pastikan fallbacknya.
	if viper.GetString("APP_PORT") == "" {
		if viper.IsSet("SERVICE_PORT") {
			viper.Set("APP_PORT", viper.GetString("SERVICE_PORT"))
		} else if viper.IsSet("HTTP_PORT") {
			viper.Set("APP_PORT", viper.GetString("HTTP_PORT"))
		}
	}

	if viper.GetString("APP_HOST") == "" {
		if viper.IsSet("HTTP_HOST") {
			viper.Set("APP_HOST", viper.GetString("HTTP_HOST"))
		} else if viper.IsSet("SERVICE_HOST") {
			viper.Set("APP_HOST", viper.GetString("SERVICE_HOST"))
		}
	}

	// Default yang aman jika tidak ada nilai
	viper.SetDefault("APP_PORT", "3005")
	viper.SetDefault("APP_HOST", "http://localhost:3005")

	// Logging ringkas agar mudah debugging CI (tidak menampilkan password)
	log.Printf("config: APP_HOST=%s APP_PORT=%s DB_USER=%s DB_HOST=%s DB_NAME=%s",
		viper.GetString("APP_HOST"),
		viper.GetString("APP_PORT"),
		viper.GetString("DB_POSTGRES_USER"),
		viper.GetString("DB_POSTGRES_HOST"),
		viper.GetString("DB_POSTGRES_NAME"),
	)
}