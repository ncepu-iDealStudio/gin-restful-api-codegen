// coding: utf-8
// @Author : lryself
// @Date : 2020/11/14 0:27
// @Software: GoLand

package views

//func IndexHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"UserName": temp.(ginModels.UserModel),
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "index.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
//func LoginHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "login.html", gin.H{
//			"UserName": temp.(modelManager.UserMin).UserName,
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "login.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
//func RegisterHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "register.html", gin.H{
//			"UserName": temp.(modelManager.UserMin).UserName,
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "register.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
//func ChangePasswordHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "changePassword.html", gin.H{
//			"UserName": temp.(modelManager.UserMin).UserName,
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "changePassword.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
//
//func MaterialHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "material.html", gin.H{
//			"UserName": temp.(modelManager.UserMin).UserName,
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "material.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
//func NoteHandler(c *gin.Context) {
//	session := sessions.Default(c)
//	temp := session.Get("user")
//	if temp != nil {
//		c.HTML(http.StatusOK, "note.html", gin.H{
//			"UserName": temp.(modelManager.UserMin).UserName,
//			"UserID":   temp.(modelManager.UserMin).UserID,
//			"IsActive": temp.(modelManager.UserMin).IsActive,
//		})
//	} else {
//		c.HTML(http.StatusOK, "note.html", gin.H{
//			"IsActive": false,
//		})
//	}
//}
