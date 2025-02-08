terraform {
  backend "s3" {
    bucket = "nuwan-lab"
    key    = "curricular-api.tfstate"
  }
}
