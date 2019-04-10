package api

type Greet struct {
}

func (_ Greet) Hello(tokenString [1]string, reply *string) error {
	claims, err := getClaimsFromTokenString(tokenString[0])
	if err != nil {
		return err
	}
	*reply = "Welcome, " + claims.Username + "!"
	return nil
}
