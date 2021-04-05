db.createUser(
    {
        user: "ampgoadmin",
        pwd: "fuckyoumongo",
        roles: [
            {
                role: "readWrite",
                db: "ampgodb"
            }
        ]
    }
)
