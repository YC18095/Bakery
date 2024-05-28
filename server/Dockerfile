FROM nginx:1.23.0
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY ./images /usr/share/nginx/images
# RUN mkdir -p /var/cache/data
CMD /bin/bash -c "envsubst '\$PORT \$FRONTEND_URL \$BACKEND_URL' < /etc/nginx/conf.d/default.conf > /etc/nginx/conf.d/default.conf" && nginx -g 'daemon off;'
# CMD ["nginx", "-g", "daemon off;"]
