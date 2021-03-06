package message

import (
	"fmt"
	"time"
)

// postgresql bot user_id

// getBotID : get bot user_id in postgresql
func getBotID() (botID int, err error) {

	rows, err := pgConn.Query("SELECT ID FROM users_userprofile WHERE username='bot'")
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		err = rows.Scan(&botID)
		if err != nil {
			return 0, err
		}
	}

	return botID, nil
}

// GetUserID : get id of all the users
func getUserMap() (userMap map[string]int, err error) {
	userMap = make(map[string]int)
	rows, err := pgConn.Query("SELECT id, username FROM users_userprofile")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id int
		var username string
		err = rows.Scan(&id, &username)
		if err != nil {
			return nil, err
		}
		userMap[username] = id
	}

	return userMap, nil
}

// Save2DB : messaeg写入posgres
func save2DB(message string) (err error) {
	mutex.Lock()
	defer mutex.Unlock()
	messageTitle := "来自bot的消息"
	messageContent := message

	sendDate := time.Now().Format("2006-01-02")

	var messageID int

	sql := fmt.Sprintf("INSERT INTO message_message (message_title, message_content, message_status, send_time, sender_id)"+
		"values('%s', '%s', '信息', '%s', %d) RETURNING message_id", messageTitle, messageContent, sendDate, botID)

	if err = pgConn.QueryRow(sql).Scan(&messageID); err != nil {
		return err
	}

	stmt, err := pgConn.Prepare("INSERT INTO message_message_receiver (message_id, userprofile_id) values ($1, $2)")
	defer stmt.Close()

	if err != nil {
		return err
	}

	for _, id := range UserMap {
		if _, err = stmt.Exec(messageID, id); err != nil {
			return err
		}
	}

	return nil
}

// AddMember : 向userMap中添加新用户
func AddMember(username string) (err error) {
	var userID int
	sql := fmt.Sprintf("SELECT id FROM users_userprofile WHERE username='%s'", username)

	if err = pgConn.QueryRow(sql).Scan(&userID); err != nil {
		return err
	}
	UserMap[username] = userID
	return nil
}
