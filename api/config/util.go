package config

import "os"

// GetEnvString 環境変数があればそれを使いなければ指定した値を返す
func GetEnvString(param string, defaultValue string) string {
       env := os.Getenv(param)
       if env != "" {
                       return env
       }
       return defaultValue
}
