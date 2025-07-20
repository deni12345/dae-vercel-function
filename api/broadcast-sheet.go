package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dae-vercel-function/cloud"
	"github.com/dae-vercel-function/model"
)

// @Summary      Broadcast sheet changes in real-time
// @Description  Stream sheet changes using Server-Sent Events (SSE)
// @Tags         sheets
// @Accept       json
// @Produce      text/event-stream
// @Param        sheetID  query     string                true  "Sheet ID to observe"
// @Success      200      {object}  model.DocumentChange  "Stream of sheet changes"
// @Failure      400      {string}  string                "Bad Request - Missing or invalid sheetID"
// @Failure      408      {string}  string                "Request Timeout - Client cancelled the request"
// @Failure      500      {string}  string                "Internal Server Error"
// @Router       /broadcast-sheet [get]
func BroadcastSheetHandler(w http.ResponseWriter, r *http.Request) {
	SetEventStreamHeaders(w)

	sheetID, err := GetURLParam(r, "sheetID")
	if err != nil {
		cloud.LogError("Failed to get sheetID from URL: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	firestore := cloud.NewFireStore(r.Context(), "drink-and-eat-b7e64")
	var (
		streamChan = make(chan *model.DocumentChange)
		errChan    = make(chan error)
	)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Failed to create streaming response", http.StatusInternalServerError)
		return
	}

	defer func() {
		close(streamChan)
		close(errChan)
		firestore.Close()
	}()

	go func() {
		if err := firestore.ObserveSheetCollection(r.Context(),
			cloud.ObserveSheetColletionReq{
				SheetID:    sheetID,
				StreamChan: streamChan,
			},
		); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-r.Context().Done():
		http.Error(w, "Request cancelled by client", http.StatusRequestTimeout)
		return
	case err := <-errChan:
		cloud.LogError("Failed to verify Firestore client initialization: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	case change := <-streamChan:
		fmt.Fprintf(w, "data: %s\n\n", toJSONString(change))
		flusher.Flush()
	}

}

func SetEventStreamHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

func GetURLParam(r *http.Request, key string) (string, error) {
	var res string
	if res = r.URL.Query().Get(key); res == "" {
		return "", fmt.Errorf("missing required parameter: %s", key)
	}
	return res, nil
}

func toJSONString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		cloud.LogError("Failed to convert to JSON string: %v", err)
		return ""
	}
	return string(bytes)
}
