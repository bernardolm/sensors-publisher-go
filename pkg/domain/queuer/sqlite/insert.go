package sqlite

// const (
// 	queryInsert = "INSERT INTO messages (target, topic, payload, created_at) VALUES (?, ?, ?, ?)"
// )

// func (a *sqlite) Save(ctx context.Context, topic string, message any) error {
// 	if message == nil {
// 		return nil
// 	}

// 	payload, ok := message.([]byte)
// 	if !ok {
// 		return fmt.Errorf("publisher: influxdb payload type %T", message)
// 	}

// 	log.WithField("message", string(payload)).
// 		WithField("publisher", "influxdb").
// 		Debug("publisher: trying to publish")

// 	if err := sqliteclient.Save(ctx, topic, payload); err != nil {
// 		return err
// 	}

// 	return nil
// }
