<?php


namespace database;
class Database
{

    private static self|null $instance = null;
    private $connect;
    private $ip;
    private $port;
    private $user;
    private $password;
    private $dbname;

    private function __construct()
    {
        $this->ip = "localhost";
        $this->port = 5432;
        $this->user = "pguser";
        $this->password = "pgpass";
        $this->dbname = "postgres";
        $this->connect = $this->getConnection();
    }

    private function getConnection(): false|Connection
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

    public static function getDatabase(): Database
    {
        if (is_null(self::$instance)) {
            self::$instance = new Database();
        }

        return self::$instance;
    }

    public function query($sql, $params = []): array
    {
        $stmt = pg_query_params($this->connect, $sql, $params);
        return $this->getArrayBD($stmt);
    }

}

?>