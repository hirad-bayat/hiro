package Controllers

import (
	"Hiro/Database"
	"Hiro/Models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// LoginForm represents login request data
type LoginForm struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

// RegisterForm represents registration request data
type RegisterForm struct {
	Name     string `form:"name" binding:"required,min=2"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "Home",
	})
}

func LoginPage(c *gin.Context) {
	// Redirect if already logged in
	session := sessions.Default(c)
	if session.Get("user_id") != nil {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	c.HTML(http.StatusOK, "login", gin.H{
		"title": "Login",
	})
}

func RegisterPage(c *gin.Context) {
	// Redirect if already logged in
	session := sessions.Default(c)
	if session.Get("user_id") != nil {
		c.Redirect(http.StatusFound, "/dashboard")
		return
	}

	c.HTML(http.StatusOK, "register", gin.H{
		"title": "Register",
	})
}

func DashboardPage(c *gin.Context) {
	user, _ := c.Get("current_user")
	c.HTML(http.StatusOK, "dashboard", gin.H{
		"title": "Dashboard",
		"user":  user,
	})
}

func LoginHandler(c *gin.Context) {
	var form LoginForm

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "login", gin.H{
			"title": "Login",
			"error": "Invalid form data",
			"email": form.Email,
		})
		return
	}

	// Find user by email
	var user Models.User
	if err := Database.DB.Where("email = ?", form.Email).First(&user).Error; err != nil {
		c.HTML(http.StatusUnauthorized, "login", gin.H{
			"title": "Login",
			"error": "Invalid email or password",
			"email": form.Email,
		})
		return
	}

	// Check password
	if !CheckPasswordHash(form.Password, user.Password) {
		c.HTML(http.StatusUnauthorized, "login", gin.H{
			"title": "Login",
			"error": "Invalid email or password",
			"email": form.Email,
		})
		return
	}

	// Create session
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_email", user.Email)
	session.Set("user_name", user.Name)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "login", gin.H{
			"title": "Login",
			"error": "Failed to create session",
		})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func RegisterHandler(c *gin.Context) {
	var form RegisterForm

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "register", gin.H{
			"title": "Register",
			"error": "Invalid form data",
			"form":  form,
		})
		return
	}

	// Check if email already exists
	var existingUser Models.User
	if err := Database.DB.Where("email = ?", form.Email).First(&existingUser).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register", gin.H{
			"title": "Register",
			"error": "Email already registered",
			"form":  form,
		})
		return
	}

	// Hash password
	hashedPassword, err := HashPassword(form.Password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register", gin.H{
			"title": "Register",
			"error": "Failed to process password",
			"form":  form,
		})
		return
	}

	// Create user
	user := Models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: hashedPassword,
	}

	if err := Database.DB.Create(&user).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "register", gin.H{
			"title": "Register",
			"error": "Failed to create account",
			"form":  form,
		})
		return
	}

	// Auto-login after registration
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_email", user.Email)
	session.Set("user_name", user.Name)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusInternalServerError, "register", gin.H{
			"title": "Register",
			"error": "Account created but failed to login",
			"form":  form,
		})
		return
	}

	c.Redirect(http.StatusFound, "/dashboard")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/")
}
