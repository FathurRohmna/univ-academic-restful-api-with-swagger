FROM nginx:alpine

COPY index.html /usr/share/nginx/html/
COPY Univ_Adacemic.swagger.yml /usr/share/nginx/html/

RUN sed -i 's/listen       80;/listen       8080;/' /etc/nginx/conf.d/default.conf

EXPOSE 8080

# gcloud run deploy univ-academic-swagger \
#   --image gcr.io/persuasive-axe-317613/univ-academic-swagger \
#   --platform managed \
#   --region us-central1 \
#   --allow-unauthenticated

# https://univ-academic-swagger-70017640279.us-central1.run.app/#/
# https://univ-academic-restful-api-70017640279.us-central1.run.app

# gcloud run deploy univ-academic-restful-api \
#   --image gcr.io/persuasive-axe-317613/univ-academic-restful-api \
#   --platform managed \
#   --region us-central1 \
#   --allow-unauthenticated