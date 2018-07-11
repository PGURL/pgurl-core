CREATE KEYSPACE short WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
USE short;
CREATE TABLE shorturl (
                        hashurl text,
                        url text,
                        createdtime timestamp,
                        primary key(hashurl)
                    );
