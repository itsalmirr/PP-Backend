schema "public" {
    url = "ent://ent/schema"

    env "local" {
        url = getenv("DB_URL")
        schema {
            src = "data.composite_schema.schema.url"
            dev = "docker://pgvector/pg17/dev"
        }
    }
}