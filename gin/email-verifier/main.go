package main

import (
	"fmt"
	"log"
	"net/http"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
)

var (
	verifier = emailverifier.NewVerifier()
)

func main() {
	// EnableSMTPCheck enables check email by smtp
	// most ISPs block outgoing SMTP requests through port 25,
	// to prevent spam, we don't check smtp by default
	verifier = verifier.EnableSMTPCheck()

	verifier = verifier.EnableDomainSuggest()
	verifier = verifier.AddDisposableDomains([]string{"tractorjj.com"})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/verifyemail", verEmailGetHandler)
	router.POST("/verifyemail", verEmailPostHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func verEmailGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "ver-email.html", nil)
}

func verEmailPostHandler(c *gin.Context) {
	fmt.Println("verEmailPostHandler running")
	email := c.PostForm("email")
	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		c.HTML(http.StatusInternalServerError, "ver-email.html", gin.H{"message": "unable to register email addresss, please try again"})
		return
	}

	fmt.Println("email validation result", ret)
	fmt.Println("Email:", ret.Email, "\nReachable:", ret.Reachable, "\nSyntax:", ret.Syntax, "\nSMTP:", ret.SMTP, "\nGravatar:", ret.Gravatar, "\nSuggestion:", ret.Suggestion, "\nDisposable:", ret.Disposable, "\nRoleAccount:", ret.RoleAccount, "\nFree:", ret.Free, "\nHasMxRecords:", ret.HasMxRecords)

	// needs @ and . for starters
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		c.HTML(http.StatusBadRequest, "ver-email.html", gin.H{"message": "email address syntax is invalid"})
		return
	}
	if ret.Disposable {
		fmt.Println("sorry, we do not accept disposable email addresses")
		c.HTML(http.StatusBadRequest, "ver-email.html", gin.H{"message": "sorry, we do not accept disposable email addresses"})
		return
	}
	if ret.Suggestion != "" {
		fmt.Println("email address is not reachable, looking for ", ret.Suggestion, "instead?")
		c.HTML(http.StatusBadRequest, "ver-email.html", gin.H{"message": "email address is not reachable, looking for " + ret.Suggestion + " instead?"})
		return
	}
	// possible return string values: yes, no, unkown
	if ret.Reachable == "no" {
		fmt.Println("email address is not reachable")
		c.HTML(http.StatusBadRequest, "ver-email.html", gin.H{"message": "email address was unreachable"})
		return
	}
	// check MX records so we know DNS setup properly to recieve emails
	if !ret.HasMxRecords {
		fmt.Println("domain entered not properly setup to recieve emails, MX record not found")
		c.HTML(http.StatusBadRequest, "ver-email.html", gin.H{"message": "domain entered not properly setup to recieve emails, MX record not found"})
		return
	}
	// ... code to register user
	c.HTML(http.StatusOK, "ver-email-result.html", gin.H{"email": email})
}
