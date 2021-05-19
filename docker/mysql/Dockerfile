FROM percona:5.7

COPY --chown=mysql ./conf/ /etc/my.cnf.d/
COPY --chown=mysql ./data/ /docker-entrypoint-initdb.d/

RUN chmod 0444 /etc/my.cnf.d/mysql.cnf && chmod 0444 /etc/my.cnf.d/mysqld.cnf