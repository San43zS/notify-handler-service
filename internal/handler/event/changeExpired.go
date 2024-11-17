package event

import (
	message "Notify-handler-service/internal/handler/model/msg"
	"Notify-handler-service/internal/handler/model/msg/parser/notifyParser"
	"context"
)

func (h Handler) ChangeExpired(ctx context.Context, msg []byte) error {
	m, err := notifyParser.New().Parse(msg)
	if err != nil {
		return err
	}

	err = h.srv.NotificationPsql().ChangeStatus(ctx, m.Id, message.OldStatus)
	if err != nil {
		return err
	}

	return nil
}
