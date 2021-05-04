resource "google_storage_bucket" "app_name-bucket" {
  name     = var.bucket
  location = "ASIA-NORTHEAST1"
}

resource "google_compute_backend_bucket" "image-773-cdn" {
  bucket_name = google_storage_bucket.app_name-bucket.name
  name        = "backend-773-bucket"
  enable_cdn  = true
  description = "Contains 773images"
}

resource "google_storage_bucket_access_control" "app_name-access" {
  bucket = google_storage_bucket.app_name-bucket.name
  role   = "READER"
  entity = "allUsers"
}
