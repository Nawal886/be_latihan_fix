package config

var allowedOrigins = []string{
	"http://localhost:5173",
	// "https://belatihanfix-production.up.railway.app",
	"https://my-fe-teal.vercel.app",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}