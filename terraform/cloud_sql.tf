resource "google_sql_database_instance" "app_name-db" {
  name             = var.db
  database_version = "MYSQL_8_0"
  region           = var.region

  settings {
    tier = "db-g1-small"
  }
}

resource "google_sql_user" "app_name-user" {
  instance = google_sql_database_instance.app_name-db.name
  name     = "root"
  password = "password"
}

resource "google_sql_database" "app_name-db" {
  instance = google_sql_database_instance.app_name-db.name
  name     = "app_name-db"
  charset  = "utf8mb4"
}
