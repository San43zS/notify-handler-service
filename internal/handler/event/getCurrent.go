package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"encoding/json"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("event")

func (h Handler) GetCurrentNotify(ctx context.Context, msg []byte) error {
	old, err := h.srv.NotificationPsql().GetCurrent(ctx, 43)
	if err != nil {
		return err
	}

	TMP := msg2.STRUCT{
		Type: "showCurrent",
		Data: old,
	}

	arr, err := json.Marshal(TMP)
	err = h.srv.NotificationRabbit().AddExpired(ctx, arr)
	if err != nil {
		log.Criticalf("failed to handle message: %v", err)
		return err
	}
	return nil
}
