package user

import(
    "ct-authentication-server/server"
    "golang.org/x/crypto/bcrypt"
    "time"
)

func CreateUser(username string, password []byte) (*User, error) {
    encodedPassword, err := encodePassword(password)
    if err != nil {
        return nil, err
    }
    user := User{
        Username: username,
        Password: encodedPassword,
        CreatedAt: time.Now(),
    }
    stmt, err := server.App.DB.Prepare("INSERT INTO user__users(username, password, created_at) VALUES(?, ?, ?)")
    if err != nil {
        return nil, err
    }
    res, err := stmt.Exec(user.Username, user.Password, user.CreatedAt)
    if err != nil {
        return nil, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return nil, err
    }
    user.Id = uint(id)
    return &user, nil
}

func encodePassword(password []byte) ([]byte, error) {
    // Use GenerateFromPassword to hash & salt pwd.
    // MinCost is just an integer constant provided by the bcrypt
    // package along with DefaultCost & MaxCost.
    // The cost can be any value you want provided it isn't lower
    // than the MinCost (4)
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
    if err != nil {
        return []byte(""), err
    }
    // GenerateFromPassword returns a byte slice so we need to
    // convert the bytes to a string and return it
    return hash, nil
}
