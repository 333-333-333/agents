r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
})

r.GET("/ready", func(c *gin.Context) {
    // Check DB, messaging, etc.
    if err := db.PingContext(c.Request.Context()); err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not ready", "error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": "ready"})
})
