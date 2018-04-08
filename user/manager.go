package user

import(
    "ct-authentication-server/client"
    "ct-authentication-server/exception"
    "ct-authentication-server/security"
    "ct-authentication-server/server"
    "golang.org/x/crypto/bcrypt"
    "github.com/go-sql-driver/mysql"
    "time"
)

func CreateUser(username string, password []byte) *User {
    encodedPassword, err := encodePassword(password)
    if err != nil {
        panic(exception.New(500, "Password could not be encoded"))
    }
    user := User{
        Username: username,
        Password: encodedPassword,
        CreatedAt: time.Now(),
    }
    stmt, err := server.App.DB.Prepare("INSERT INTO user__users(username, password, created_at) VALUES(?, ?, ?)")
    if err != nil {
        panic(exception.New(500, "User creation query preparation failed"))
    }
    res, err := stmt.Exec(user.Username, user.Password, user.CreatedAt)
    if err != nil {
        panic(exception.New(500, "User account could not be created"))
    }
    id, err := res.LastInsertId()
    if err != nil {
        panic(exception.New(500, "Created User ID could not be retrieved"))
    }
    user.Id = uint(id)
    return &user
}

func Connect(service *client.Client, username string, password []byte) *User {
    user, err := GetUserByUsername(username)
    if err != nil {
        panic(exception.New(404, "User not found"))
    }
    if err := bcrypt.CompareHashAndPassword(user.Password, password); err != nil {
        panic(exception.New(400, "Invalid credentials"))
    }
    user.AccessToken = security.GenerateRandomToken(32)
    user.RefreshToken = security.GenerateRandomToken(32)
    user.ExpiresAt = time.Now().Add(time.Hour * time.Duration(2))
    err = server.App.Redis.HMSet("user__" + string(user.AccessToken), map[string]interface{}{
        "username": user.Username,
        "access_token": user.AccessToken,
        "refresh_token": user.RefreshToken,
        "expires_at": user.ExpiresAt,
        "created_at": user.CreatedAt,
        "last_connected_at": user.LastConnectedAt,
    }).Err()
    if err != nil {
        panic(exception.New(500, "User Redis session could not be created"))
    }
    return user
}

func GetUserByUsername(username string) (*User, error) {
    user := User{
        Username: username,
    }
    var createdAt mysql.NullTime
    var lastConnectedAt mysql.NullTime
    err := server.App.DB.QueryRow("SELECT id, password, created_at, last_connected_at FROM user__users WHERE username = ?", username).Scan(
        &user.Id,
        &user.Password,
        &createdAt,
        &lastConnectedAt,
    )
    user.CreatedAt = createdAt.Time
    user.LastConnectedAt = lastConnectedAt.Time
    if err != nil {
        return nil, err
    }
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
