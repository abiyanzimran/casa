package api

import "maps"

// NoticeHeldProcedures is a copy of noticeHeldProcedures, so the hold test
// iterates the set the boundary reads (plan 1.6 decision 7).
func NoticeHeldProcedures() map[string]bool { return maps.Clone(noticeHeldProcedures) }
