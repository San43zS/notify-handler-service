package event

import (
	msg2 "Notify-handler-service/internal/handler/model/msg"
	"context"
	"encoding/json"
)

func (h Handler) GetOld(ctx context.Context, msg []byte) error {
	old, err := h.srv.NotificationPsql().GetOld(ctx, 43)
	TMP := msg2.STRUCT{
		Type: "showOld",
		Data: old,
	}

	arr, err := json.Marshal(TMP)
	err = h.srv.NotificationRabbit().AddExpired(ctx, arr)
	if err != nil {
		return err
	}
	return nil
}
