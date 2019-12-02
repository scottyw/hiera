Build an image with content baked in:

    docker build -f Dockerfile-content . -t gcr.io/mimosa-256008/hiera

Push to GCR and deploy in Cloud Run:

    docker push gcr.io/mimosa-256008/hiera

Run like this:

    go build -o terraform-provider-hiera && terraform init && terraform apply -var workspace=workspace1
