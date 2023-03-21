//1. Redis
KEYS *User
//2. Dalam Neo4j
MATCH (u:User) RETURN u;
//3. Cassandra
SELECT * FROM users;