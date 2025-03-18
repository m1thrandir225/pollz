package api

import "github.com/gin-gonic/gin"

func (server *Server) setupRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	/** Public Routes */
	v1.POST("/login", server.loginUser)
	v1.POST("/register", server.registerUser)
	v1.POST("/refresh-token", server.refreshToken)

	authRoutes := v1.Group("/").Use(authMiddleware(server.tokenMaker))

	/** User Routes */
	authRoutes.GET("/user/:id", server.getUser)

	/** Poll Routes */

	authRoutes.GET("/polls", server.getPolls) // Get your polls
	v1.GET("/polls/:id", server.getPoll)

	authRoutes.POST("/polls", server.createPoll)
	authRoutes.PUT("/polls/:id", server.updatePoll)
	authRoutes.PATCH("/polls/:id/status", server.updatePollStatus)
	authRoutes.DELETE("/polls/:id", server.deletePoll)

	/** Poll Option Routes */
	authRoutes.POST("/poll-options", server.createOption)
	authRoutes.PUT("/poll-options/:id", server.updateOption)
	authRoutes.DELETE("/poll-options/:id", server.deleteOption)

	/** Vote Routes */
	v1.POST("/vote", server.createVote)
	v1.PUT("/vote/:id", server.updateVote)
	v1.DELETE("/vote/:id", server.deleteVote)

	server.router = router
}
