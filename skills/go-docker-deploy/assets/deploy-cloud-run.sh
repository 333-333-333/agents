gcloud run deploy booking \
	--image gcr.io/PROJECT/booking:latest \
	--port 8080 \
	--set-env-vars "APP_ENV=production,DB_HOST=..." \
	--region us-central1
