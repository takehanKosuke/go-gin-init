resource "google_cloud_run_service" "default" {
  name     = var.service
  location = var.region
  project  = var.project

  template {
    spec {
      containers {
        image = "asia.gcr.io/${var.project}/${var.server}/${var.buket}:latest"
        env {
          name  = "DB_USER"
          value = google_sql_user.app_name-user.name
        }
        env {
          name  = "DB_PASS"
          value = google_sql_user.app_name-user.password
        }
        env {
          name  = "INSTANCE_CONNECTION_NAME"
          value = "${var.buket}:asia-northeast1:${var.db}"
        }
        env {
          name  = "DB_PORT"
          value = 3306
        }
        env {
          name  = "DB_NAME"
          value = google_sql_database.app_name-db.name
        }
        env {
          name  = "ENV"
          value = "prd"
        }
      }
    }
  }
  metadata {
    annotations = {
      "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.app_name-db.connection_name
    }
  }
  autogenerate_revision_name = true
}
