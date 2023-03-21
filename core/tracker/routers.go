package tracker

func (t *tracker) Server() {

	t.g.Use(t.Download())
	t.g.POST("/upload", t.Upload)
	t.g.POST("/report-status", t.HanldeStorageServerReport)
	t.g.Run(t.config.ListenOn)
}
