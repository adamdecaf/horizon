FROM postgres:9.5
WORKDIR /var/log/postgresql
VOLUME ["/var/lib/postgresql/data", "/var/log/postgresql"]
ADD postgresql.conf /var/lib/postgresql/data/postgresql.conf
EXPOSE 5432
