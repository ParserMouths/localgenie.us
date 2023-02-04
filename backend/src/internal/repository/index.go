package repository

import (
	domain_notification "htf/src/internal/domain/notification"
	domain_stall "htf/src/internal/domain/stall"
	domain_user "htf/src/internal/domain/user"
)

type Repositories struct {
	UserRepo         domain_user.Repository
	NotificationRepo domain_notification.Repository
	StallRepo        domain_stall.Repository
}
