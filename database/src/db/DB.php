<?php

namespace OnlineShop\db;

class DB
{
    private $connect;

    private $ip;
    private $port;
    private $user;
    private $password;
    private $dbname;

    public function __construct()
    {
        $this->connect = $this->getDatabase();
        $this->ip = "localhost";
        $this->port = 5432;
        $this->user = "pguser";
        $this->password = "pgpass";
        $this->dbname = "postgres";
    }

    private function getDatabase(): false|\PgSql\Connection
    {
        return pg_connect("host=" . $this->ip . " port=" . $this->port . " dbname=" . $this->dbname . " user=" . $this->user . " password=" . $this->password);
    }

    private function getArrayBD($stmt): array
    {
        $result_query_api = array();
        while ($line = pg_fetch_array($stmt, null, PGSQL_ASSOC)) {
            $result_query_api[] = $line;
        }
        return $result_query_api;
    }

    public function query($sql, $params = []): array
    {
        $stmt = pg_query_params($this->connect, $sql, $params);
        return $this->getArrayBD($stmt);
    }

}