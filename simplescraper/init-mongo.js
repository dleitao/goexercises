db.createUser(
    {
        user: "admin",
        pwd: "asd123",
        roles: [
            {
                role: "readWrite",
                db: "scraper-db"
            }
        ]

    }
)